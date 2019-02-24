package internal

import (
	"os"
	"strings"
	"sync"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logging"
	"github.com/kihamo/shadow/components/logging/internal/wrapper"
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
	global      *wrapper.Wrapper

	lock          sync.Mutex
	restoreStdLog func()
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
	c.global = logging.DefaultLogger().(*wrapper.Wrapper)

	c.config = a.GetComponent(config.ComponentName).(config.Component)

	return nil
}

func (c *Component) Run(a shadow.Application, _ chan<- struct{}) error {
	<-a.ReadyComponent(config.ComponentName)
	c.initLogger()

	return nil
}

func (c *Component) initLogger() {
	var (
		encoderConfig zapcore.EncoderConfig
		encoder       zapcore.Encoder
	)

	if c.config.String(logging.ConfigMode) == logging.ModeProduction {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.MessageKey = "message"

	encoderConfig.TimeKey = "time"
	switch c.config.String(logging.ConfigEncoderTime) {
	case logging.EncoderTimeISO8601:
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	case logging.EncoderTimeMillis:
		encoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	case logging.EncoderTimeNanos:
		encoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
	case logging.EncoderTimeSeconds:
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	}

	switch c.config.String(logging.ConfigEncoderDuration) {
	case logging.EncoderDurationString:
		encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	case logging.EncoderDurationSeconds:
		encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	case logging.EncoderDurationNanos:
		encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	}

	encoderConfig.CallerKey = "file"
	switch c.config.String(logging.ConfigEncoderCaller) {
	case logging.EncoderCallerShort:
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	case logging.EncoderCallerFull:
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	}

	if c.config.String(logging.ConfigEncoderType) == logging.EncoderTypeJSON {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	output := zapcore.AddSync(os.Stdout)

	c.global.InitFull(
		encoder, output, zapcore.Level(c.config.Int64(logging.ConfigLevel)),
		zap.Fields(c.parseFields(c.config.String(logging.ConfigFields))...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.Level(c.config.Int64(logging.ConfigStacktraceLevel))),
	)

	zap.ReplaceGlobals(c.global.Logger())
	std := zap.RedirectStdLog(c.global.LoadOrStore("std").Logger())

	c.lock.Lock()
	if c.restoreStdLog == nil {
		c.restoreStdLog = std
	}
	c.lock.Unlock()
}

func (c *Component) parseFields(f string) []zap.Field {
	fields := make([]zap.Field, 0)

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
	return c.global
}

func (c *Component) Shutdown() error {
	c.lock.Lock()
	if c.restoreStdLog != nil {
		c.restoreStdLog()
	}
	c.lock.Unlock()

	if err := c.global.Logger().Sync(); err != nil {
		return err
	}

	return nil
}
