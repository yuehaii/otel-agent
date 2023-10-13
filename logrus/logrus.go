package logrus

import (
	sirupsen_logrus "github.com/sirupsen/logrus"
	logger "github.com/yuehaii/otel-agent/otellog"
)

//agent of https://github.com/sirupsen/logrus

func Tracef(format string, args ...interface{}) {
	logger.OtelLogger("TraceLevel", format, args...)
	sirupsen_logrus.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.OtelLogger("DebugLevel", format, args...)
	sirupsen_logrus.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.OtelLogger("InfoLevel", format, args...)
	sirupsen_logrus.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.OtelLogger("WarnLevel", format, args...)
	sirupsen_logrus.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.OtelLogger("ErrorLevel", format, args...)
	sirupsen_logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.OtelLogger("FatalLevel", format, args...)
	sirupsen_logrus.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.OtelLogger("PanicLevel", format, args...)
	sirupsen_logrus.Panicf(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.OtelLogger("InfoLevel", format, args...)
	sirupsen_logrus.Printf(format, args...)
}

type Entry = sirupsen_logrus.Entry
type Logger = sirupsen_logrus.Logger
type JSONFormatter = sirupsen_logrus.JSONFormatter
type TextFormatter = sirupsen_logrus.TextFormatter
type Level = sirupsen_logrus.Level

var (
	StandardLogger  = sirupsen_logrus.StandardLogger
	SetLevel        = sirupsen_logrus.SetLevel
	SetFormatter    = sirupsen_logrus.SetFormatter
	SetReportCaller = sirupsen_logrus.SetReportCaller
	SetOutput       = sirupsen_logrus.SetOutput
	WithContext     = sirupsen_logrus.WithContext
)
