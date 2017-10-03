package internal

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
	"github.com/kihamo/shadow/components/profiling"
	"github.com/kihamo/snitch"
	_ "github.com/kihamo/snitch/collector"
	"github.com/kihamo/snitch/storage"
)

const (
	TagAppName    = "app_name"
	TagAppVersion = "app_version"
	TagAppBuild   = "app_build"
	TagHostname   = "hostname"
)

type Component struct {
	application shadow.Application

	config config.Component
	logger logger.Logger
	routes []dashboard.Route

	mutex    sync.RWMutex
	registry snitch.Registerer
	storage  *storage.Influx
}

func (c *Component) GetName() string {
	return metrics.ComponentName
}

func (c *Component) GetVersion() string {
	return metrics.ComponentVersion
}

func (c *Component) GetDependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: dashboard.ComponentName,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: profiling.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(config.Component)
	c.registry = snitch.DefaultRegisterer

	if c.application.HasComponent(profiling.ComponentName) {
		c.registry.AddStorages(storage.NewExpvarWithId(metrics.ComponentName))
	}

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.GetName(), c.application)

	url := c.config.GetString(metrics.ConfigUrl)
	if url == "" {
		wg.Done()
		return fmt.Errorf("%s is empty", metrics.ConfigUrl)
	}

	storage, err := storage.NewInflux(
		c.config.GetString(metrics.ConfigUrl),
		c.config.GetString(metrics.ConfigDatabase),
		c.config.GetString(metrics.ConfigUsername),
		c.config.GetString(metrics.ConfigPassword),
		c.config.GetString(metrics.ConfigPrecision))
	if err != nil {
		wg.Done()
		return nil
	}

	c.registry.AddStorages(storage)
	c.registry.SendInterval(c.config.GetDuration(metrics.ConfigInterval))

	c.storage = storage
	c.initLabels(c.config.GetString(metrics.ConfigLabels))

	// search metrics
	components, err := c.application.GetComponents()
	if err != nil {
		return err
	}

	for _, component := range components {
		if m, ok := component.(metrics.HasMetrics); ok {
			c.Register(m.Metrics())
		}
	}

	return nil
}

func (c *Component) Registry() snitch.Registerer {
	return c.registry
}

func (c *Component) Register(cs ...snitch.Collector) {
	c.registry.Register(cs...)
}

func (c *Component) initLabels(labels string) {
	l := snitch.Labels{
		&snitch.Label{Key: TagAppName, Value: c.application.GetName()},
		&snitch.Label{Key: TagAppVersion, Value: c.application.GetVersion()},
		&snitch.Label{Key: TagAppBuild, Value: c.application.GetBuild()},
	}

	if hostname, err := os.Hostname(); err == nil {
		l = append(l, &snitch.Label{
			Key:   TagHostname,
			Value: hostname,
		})
	}

	if len(labels) > 0 {
		var parts []string

		for _, tag := range strings.Split(labels, ",") {
			parts = strings.Split(tag, "=")

			if len(parts) > 1 {
				l = append(l, &snitch.Label{
					Key:   strings.TrimSpace(parts[0]),
					Value: strings.TrimSpace(parts[1]),
				})
			}
		}
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.registry.SetLabels(l)
}
