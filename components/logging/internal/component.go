package internal

import (
	"os"
	"strings"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	fieldAppName    = "app-name"
	fieldAppVersion = "app-version"
	fieldAppBuild   = "app-build"
	fieldHostname   = "hostname"
)

type Component struct {
	application shadow.Application
	config      config.Component
	level       zap.AtomicLevel
	wrapper     wrapper
}

func (c *Component) Name() string {
	return logging.ComponentName
}

func (c *Component) Version() string {
	return logging.ComponentVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(config.Component)
	c.level = zap.NewAtomicLevel()
	c.wrapper = logging.DefaultLogger().(wrapper)

	return nil
}

func (c *Component) Run() error {
	c.initLogger()

	return nil
}

func (c *Component) initLogger() {
	output := zapcore.Lock(os.Stderr)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	c.level.SetLevel(zapcore.Level(c.config.Int64(logging.ConfigLevel)))

	l := zap.New(
		zapcore.NewCore(encoder, output, c.level),
		zap.Fields(c.parseFields(c.config.String(logging.ConfigFields))...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.Level(c.config.Int64(logging.ConfigStacktraceLevel))))

	c.wrapper.SetLogger(l)
	zap.RedirectStdLog(l)
}

func (c *Component) parseFields(f string) []zap.Field {
	fields := make([]zap.Field, 0, 0)

	if len(f) == 0 {
		return fields
	}

	var parts []string

	for _, tag := range strings.Split(f, ",") {
		if parts = strings.Split(tag, "="); len(parts) > 1 {
			fields = append(fields, zap.String(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])))
		}
	}

	fields = append(fields, zap.String(fieldAppName, c.application.Name()))
	fields = append(fields, zap.String(fieldAppVersion, c.application.Version()))
	fields = append(fields, zap.String(fieldAppBuild, c.application.Build()))
	if hostname, err := os.Hostname(); err == nil {
		fields = append(fields, zap.String(fieldHostname, hostname))
	}

	return fields
}

func (c *Component) Logger() logging.Logger {
	return logging.DefaultLogger()
}

func (c *Component) Shutdown() error {
	if err := c.wrapper.Logger().Sync(); err != nil {
		return err
	}

	if err := c.wrapper.Sugar().Sync(); err != nil {
		return err
	}

	return nil
}
