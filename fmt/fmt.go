package fmt

import (
	go_fmt "fmt"

	logger "github.com/yuehaii/otel-agent/otellog"
)

//agent of https://pkg.go.dev/fmt

func Sprintf(format string, a ...any) string {
	logger.OtelLogger("SprintLevel", format, a...)
	return go_fmt.Sprintf(format, a...)
}

func Errorf(format string, a ...any) error {
	logger.OtelLogger("ErrorLevel", format, a...)
	return go_fmt.Errorf(format, a...)
}

func Println(a ...any) (n int, err error) {
	logger.OtelLogger("PrintLevel", "%s", a...)
	return go_fmt.Println(a...)
}
