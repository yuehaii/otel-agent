package otellog

import (
	"bytes"
	"fmt"
	"net/http"

	agentconfig "github.com/yuehaii/otel-agent/config"
)

func OtelLogger(level string, format string, args ...interface{}) bool {
	response, err := http.Post(agentconfig.GetProtol()+"://"+agentconfig.GetOtelHost()+":"+agentconfig.GetPort()+agentconfig.GetPath(),
		"application/json",
		bytes.NewBuffer([]byte(agentconfig.GenerateBody(fmt.Sprintf(format, args...), level))))
	if err != nil {
		fmt.Println("Err: ", err.Error())
		return false
	}

	if response != nil && response.StatusCode > 300 {
		fmt.Println("StatusCode: ", response.StatusCode)
		fmt.Println("Status: ", response.Status)
		fmt.Println("Body: ", response.Body)
	}
	return true
}
