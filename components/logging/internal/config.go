package internal

import (
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(logging.ConfigLevel, config.ValueTypeInt).
			WithUsage("Log level").
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
			WithUsage("Fields in format field_name=field1_value,field2_name=field2_value").
			WithEditable(true).
			WithView([]string{config.ViewTags}).
			WithViewOptions(map[string]interface{}{config.ViewOptionTagsDefaultText: "add a field"}),
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
	c.level.SetLevel(zapcore.Level(c.config.Int64(logging.ConfigLevel)))
}

func (c *Component) watchLoggerStacktraceLevel(_ string, newValue interface{}, _ interface{}) {
	c.wrapper.SetLogger(c.wrapper.Logger().WithOptions(
		zap.AddStacktrace(zapcore.Level(c.config.Int64(logging.ConfigStacktraceLevel))),
	))
}

func (c *Component) watchLoggerFields(_ string, newValue interface{}, oldValue interface{}) {
	c.wrapper.SetLogger(c.wrapper.Logger().WithOptions(
		zap.Fields(c.parseFields(newValue.(string))...),
	))
}
