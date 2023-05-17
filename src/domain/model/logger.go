package model

import "context"

type LoggerFields map[string]interface{}

//go:generate mockgen -destination=../mock/logger.go -package=mock -source=logger.go
type Logger interface {
	WithFields(map[string]interface{}) Logger
	WithRequestId(context.Context) Logger
	Print(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
	Printf(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Panicf(string, ...interface{})
}
