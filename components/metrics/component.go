package metrics

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/snitch"
	_ "github.com/kihamo/snitch/collector"
	"github.com/kihamo/snitch/storage"
)

const (
	ComponentName = "metrics"

	TagAppName    = "app_name"
	TagAppVersion = "app_version"
	TagAppBuild   = "app_build"
	TagHostname   = "hostname"
)

type Component struct {
	application shadow.Application

	config *config.Component
	logger logger.Logger

	mutex    sync.RWMutex
	prefix   string
	registry snitch.Registerer
	storage  *storage.Influx
}

type hasMetrics interface {
	Metrics() snitch.Collector
}

func (c *Component) GetName() string {
	return ComponentName
}

func (c *Component) GetVersion() string {
	return ComponentVersion
}

func (c *Component) GetDependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: dashboard.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(*config.Component)

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.GetName(), c.application)

	url := c.config.GetString(ConfigMetricsUrl)
	if url == "" {
		wg.Done()
		return fmt.Errorf("%s is empty", ConfigMetricsUrl)
	}

	storage, err := storage.NewInflux(
		c.config.GetString(ConfigMetricsUrl),
		c.config.GetString(ConfigMetricsDatabase),
		c.config.GetString(ConfigMetricsUsername),
		c.config.GetString(ConfigMetricsPassword),
		c.config.GetString(ConfigMetricsPrecision))
	if err != nil {
		wg.Done()
		return nil
	}

	c.prefix = c.config.GetString(ConfigMetricsPrefix)
	c.registry = snitch.DefaultRegisterer
	c.registry.AddStorages(storage)
	c.registry.SendInterval(c.config.GetDuration(ConfigMetricsInterval))

	c.storage = storage
	c.initLabels(c.config.GetString(ConfigMetricsLabels))

	// search metrics
	components, err := c.application.GetComponents()
	if err != nil {
		return err
	}

	for _, component := range components {
		if metrics, ok := component.(hasMetrics); ok {
			c.Register(metrics.Metrics())
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
