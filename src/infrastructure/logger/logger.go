package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"strings"
	"time"
)

const JsonFormat = "JSON"

type Fields map[string]interface{}

type logger struct {
	logger  *logrus.Logger
	dFields map[string]interface{}
}

type entry struct {
	entry   *logrus.Entry
	dFields map[string]interface{}
}

// Config is used to configure the Logger
type Config struct {
	ServiceName     string
	EnvironmentName string
	LogLevel        string
	LogFormat       string
	DefaultFields   map[string]interface{}
}

// New creates a new Logger from some configuration
func New(config *Config) model.Logger {
	fields := addIfNotEmpty(config.DefaultFields, "serviceName", config.ServiceName)
	fields = addIfNotEmpty(fields, "environment", config.EnvironmentName)
	newLogger := &logger{
		logger:  logrus.StandardLogger(),
		dFields: fields,
	}
	configure(config)
	return newLogger
}

// WithFields returns a logger with given fields, new fields overrides old ones
func (l *logger) WithFields(fields map[string]interface{}) model.Logger {
	return &entry{
		entry:   l.logger.WithFields(collectFields(l.dFields, fields)),
		dFields: l.dFields,
	}
}

// WithContext returns a logger that contains the values from a given context
func (l *logger) WithRequestId(ctx context.Context) model.Logger {
	return withRequestId(ctx, &entry{l.logger.WithContext(ctx), l.dFields})
}

func (l *logger) Print(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Print(message...)
}

func (l *logger) Debug(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Debug(message...)
}

func (l *logger) Info(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Info(message...)
}

func (l *logger) Warn(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Warn(message...)
}

func (l *logger) Error(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Error(message...)
}

func (l *logger) Fatal(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Fatal(message...)
}

func (l *logger) Panic(message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Panic(message...)
}

func (l *logger) Printf(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Printf(format, message...)
}

func (l *logger) Debugf(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Debugf(format, message...)
}

func (l *logger) Infof(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Infof(format, message...)
}

func (l *logger) Warnf(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Warnf(format, message...)
}

func (l *logger) Errorf(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Errorf(format, message...)
}

func (l *logger) Fatalf(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Fatalf(format, message...)
}

func (l *logger) Panicf(format string, message ...interface{}) {
	l.logger.WithFields(collectFields(l.dFields, map[string]interface{}{})).Panicf(format, message...)
}

func (e *entry) WithFields(fields map[string]interface{}) model.Logger {
	return &entry{
		entry:   e.entry.WithFields(collectFields(e.dFields, fields)),
		dFields: e.dFields,
	}
}

func (e *entry) WithRequestId(ctx context.Context) model.Logger {
	return withRequestId(ctx, &entry{e.entry.WithContext(ctx), e.dFields})
}

func withRequestId(ctx context.Context, baseLogger model.Logger) model.Logger {
	if reqId := GetRequestId(ctx); reqId != "" {
		return baseLogger.WithFields(Fields{requestIdKeyStr: reqId})
	}
	return baseLogger
}

func (e *entry) Print(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Print(message...)
}

func (e *entry) Debug(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Debug(message...)
}

func (e *entry) Info(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Info(message...)
}

func (e *entry) Warn(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Warn(message...)
}

func (e *entry) Error(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Error(message...)
}

func (e *entry) Fatal(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Fatal(message...)
}

func (e *entry) Panic(message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Panic(message...)
}

func (e *entry) Printf(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Printf(format, message...)
}

func (e *entry) Debugf(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Debugf(format, message...)
}

func (e *entry) Infof(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Infof(format, message...)
}

func (e *entry) Warnf(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Warnf(format, message...)
}

func (e *entry) Errorf(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Errorf(format, message...)
}

func (e *entry) Fatalf(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Fatalf(format, message...)
}

func (e *entry) Panicf(format string, message ...interface{}) {
	e.entry.WithFields(collectFields(e.dFields, map[string]interface{}{})).Panicf(format, message...)
}

func collectFields(a map[string]interface{}, b map[string]interface{}) map[string]interface{} {
	var allFields = make(map[string]interface{}, len(a)+len(b))
	for k, v := range a {
		allFields[k] = v
	}
	for k, v := range b {
		allFields[k] = v
	}
	return allFields
}

func configure(configuration *Config) {
	logrus.SetLevel(getLevel(configuration.LogLevel))
	logrus.SetFormatter(getFormatter(configuration.LogFormat))
}

func getLevel(logLevel string) logrus.Level {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return logrus.InfoLevel
	}
	return level
}

func getFormatter(format string) logrus.Formatter {
	envType := valueOrDefault(format, "plain")
	if strings.ToUpper(envType) == JsonFormat {
		return &logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano}
	}
	return &logrus.TextFormatter{TimestampFormat: time.RFC3339Nano}
}

func valueOrDefault(name string, defaultValue string) string {
	v := strings.TrimSpace(name)
	if v == "" {
		return defaultValue
	}
	return v
}

func addIfNotEmpty(fields map[string]interface{}, key string, value string) map[string]interface{} {
	if strings.TrimSpace(value) != "" {
		newFields := fields
		if len(newFields) == 0 {
			newFields = make(map[string]interface{})
		}
		newFields[key] = value
		return newFields
	}
	return fields
}
