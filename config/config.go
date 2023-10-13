package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	byteValue []byte
)

func init() {
	agentFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("err")
	}
	defer agentFile.Close()
	byteValue, err = ioutil.ReadAll(agentFile)
	if err != nil {
		fmt.Println("err")
		fmt.Println(byteValue)
	}
}

func GetOtelHost() string {
	return gjson.Get(string(byteValue), "agent.host").Str
}

func GetPort() string {
	return gjson.Get(string(byteValue), "agent.port").Str
}

func GetPath() string {
	return gjson.Get(string(byteValue), "log.path").Str
}

func GetProtol() string {
	return gjson.Get(string(byteValue), "log.protocol").Str
}

func GenerateBody(body string, level string) string {
	updatedJson, err := sjson.Set(string(byteValue), "log.logbody.resourceLogs.0.resource.attributes.7.value.stringValue", level)
	if err != nil {
		fmt.Println(err)
	}
	ux_t := strconv.FormatInt(time.Now().UnixNano(), 10)
	updatedJson, err = sjson.Set(updatedJson, "log.logbody.resourceLogs.0.scopeLogs.0.logRecords.0.timeUnixNano", ux_t)
	if err != nil {
		fmt.Println(err)
	}
	updatedJson, err = sjson.Set(updatedJson, "log.logbody.resourceLogs.0.scopeLogs.0.logRecords.0.observedTimeUnixNano", ux_t)
	if err != nil {
		fmt.Println(err)
	}
	bodyConv := "body=\"" + body + "\""
	updatedJson, err = sjson.Set(updatedJson, "log.logbody.resourceLogs.0.scopeLogs.0.logRecords.0.body.stringValue", bodyConv)

	if err != nil {
		fmt.Println(err)
	}
	return gjson.Get(updatedJson, "log.logbody").Raw
}
