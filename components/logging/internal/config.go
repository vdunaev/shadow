package internal

import (
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(logging.ConfigMode, config.ValueTypeString).
			WithGroup("General").
			WithUsage("Mode").
			WithDefault(logging.ModeProduction).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{logging.ModeProduction, "Production"},
					{logging.ModeDevelopment, "Development"},
				},
			}),
		config.NewVariable(logging.ConfigLevel, config.ValueTypeInt).
			WithGroup("General").
			WithUsage("Global log level").
			WithEditable(true).
			WithDefault(int8(zapcore.InfoLevel)).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{int8(zapcore.DebugLevel), "Debug"},
					{int8(zapcore.InfoLevel), "Informational"},
					{int8(zapcore.WarnLevel), "Warning"},
					{int8(zapcore.ErrorLevel), "Error"},
					{int8(zapcore.PanicLevel), "Panic"},
					{int8(zapcore.DPanicLevel), "Development panic"},
					{int8(zapcore.FatalLevel), "Fatal"},
				},
			}),
		config.NewVariable(logging.ConfigStacktraceLevel, config.ValueTypeInt).
			WithGroup("General").
			WithUsage("Stacktrace log level").
			WithEditable(true).
			WithDefault(int8(zapcore.FatalLevel)).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{int8(zapcore.DebugLevel), "Debug"},
					{int8(zapcore.InfoLevel), "Informational"},
					{int8(zapcore.WarnLevel), "Warning"},
					{int8(zapcore.ErrorLevel), "Error"},
					{int8(zapcore.PanicLevel), "Panic"},
					{int8(zapcore.DPanicLevel), "Development panic"},
					{int8(zapcore.FatalLevel), "Fatal"},
				},
			}),
		config.NewVariable(logging.ConfigFields, config.ValueTypeString).
			WithGroup("General").
			WithUsage("Fields in format field_name=field1_value,field2_name=field2_value").
			WithEditable(true).
			WithView([]string{config.ViewTags}).
			WithViewOptions(map[string]interface{}{config.ViewOptionTagsDefaultText: "add a field"}),
		config.NewVariable(logging.ConfigEncoderType, config.ValueTypeString).
			WithGroup("Encoder").
			WithUsage("Type").
			WithDefault(logging.EncoderTypeJSON).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{logging.EncoderTypeJSON, "JSON"},
					{logging.EncoderTypeConsole, "Console"},
				},
			}),
		config.NewVariable(logging.ConfigEncoderTime, config.ValueTypeString).
			WithGroup("Encoder").
			WithUsage("Time format").
			WithDefault(logging.EncoderTimeISO8601).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{logging.EncoderTimeISO8601, "ISO8601"},
					{logging.EncoderTimeMillis, "Milliseconds"},
					{logging.EncoderTimeNanos, "Nanoseconds"},
					{logging.EncoderTimeSeconds, "Seconds"},
				},
			}),
		config.NewVariable(logging.ConfigEncoderDuration, config.ValueTypeString).
			WithGroup("Encoder").
			WithUsage("Duration format").
			WithDefault(logging.EncoderDurationString).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{logging.EncoderDurationSeconds, "Seconds"},
					{logging.EncoderDurationNanos, "Nanoseconds"},
					{logging.EncoderDurationString, "String"},
				},
			}),
		config.NewVariable(logging.ConfigEncoderCaller, config.ValueTypeString).
			WithGroup("Encoder").
			WithUsage("Caller").
			WithDefault(logging.EncoderCallerNone).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{logging.EncoderCallerNone, "None"},
					{logging.EncoderCallerShort, "Short"},
					{logging.EncoderCallerFull, "Full"},
				},
			}),
		config.NewVariable(logging.ConfigSentryEnabled, config.ValueTypeBool).
			WithGroup("Sentry").
			WithUsage("Enabled"),
		config.NewVariable(logging.ConfigSentryDSN, config.ValueTypeString).
			WithGroup("Sentry").
			WithUsage("DSN"),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher([]string{logging.ConfigLevel}, c.watchLoggerLevel),
		config.NewWatcher([]string{logging.ConfigStacktraceLevel}, c.watchLoggerStacktraceLevel),
		config.NewWatcher([]string{logging.ConfigFields}, c.watchLoggerFields),
	}
}

func (c *Component) watchLoggerLevel(_ string, newValue interface{}, _ interface{}) {
	c.global.SetLevelEnabler(true, zapcore.Level(c.config.Int64(logging.ConfigLevel)))
}

func (c *Component) watchLoggerStacktraceLevel(_ string, newValue interface{}, _ interface{}) {
	c.global.WithOptions(
		true,
		zap.AddStacktrace(zapcore.Level(c.config.Int64(logging.ConfigStacktraceLevel))),
	)
}

func (c *Component) watchLoggerFields(_ string, newValue interface{}, oldValue interface{}) {
	c.global.WithOptions(
		true,
		zap.Fields(c.parseFields(newValue.(string))...),
	)
}
