package logrus

import (
	"bytes"
	"fmt"
	"net/http"

	sirupsen_logrus "github.com/sirupsen/logrus"
)

func Tracef(format string, args ...interface{}) {
	sirupsen_logrus.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	sirupsen_logrus.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	otelLogger("Info", format, args...)
	sirupsen_logrus.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	sirupsen_logrus.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	sirupsen_logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	sirupsen_logrus.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	sirupsen_logrus.Panicf(format, args...)
}

func Printf(format string, args ...interface{}) {
	sirupsen_logrus.Printf(format, args...)
}

func otelLogger(level string, format string, args ...interface{}) bool {
	url := "http://106.14.218.242:4318/v1/logs"

	// ref: https://github.com/open-telemetry/opentelemetry-proto/blob/main/examples/logs.json
	//ux_t := strconv.FormatInt(time.Now().UnixNano(), 10)
	body := []byte(`{
		"resourceLogs": [
			{
			  "resource": {
				"attributes": [
				  {
					"key": "service.name",
					"value": {
					  "stringValue": "my.service"
					}
				  }
				]
			  },
			  "scopeLogs": [
				{
				  "scope": {
					"name": "my.library",
					"version": "1.0.0",
					"attributes": [
					  {
						"key": "my.scope.attribute",
						"value": {
						  "stringValue": "some scope attribute"
						}
					  }
					]
				  },
				  "logRecords": [
					{
					  "timeUnixNano": 1697015555000000000,
					  "observedTimeUnixNano": 1697015555000000000,
					  "severityText": "Information",
					  "traceId": "5B8EFFF798038103D269B633813FC60C",
					  "spanId": "EEE19B7EC3C1B174",
					  "body": {
						"stringValue": "Example log record"
					  },
					  "attributes": [
						{
						  "key": "string.attribute",
						  "value": {
							"stringValue": "some string"
						  }
						},
						{
						  "key": "boolean.attribute",
						  "value": {
							"boolValue": true
						  }
						},
						{
						  "key": "int.attribute",
						  "value": {
							"intValue": 10
						  }
						},
						{
						  "key": "double.attribute",
						  "value": {
							"doubleValue": 637.704
						  }
						},
						{
						  "key": "array.attribute",
						  "value": {
							"arrayValue": {
							  "values": [
								{
								  "stringValue": "many"
								},
								{
								  "stringValue": "values"
								}
							  ]
							}
						  }
						},
						{
						  "key": "map.attribute",
						  "value": {
							"kvlistValue": {
							  "values": [
								{
								  "key": "some.map.key",
								  "value": {
									"stringValue": "some value"
								  }
								}
							  ]
							}
						  }
						}
					  ]
					}
				  ]
				}
			  ]
			}
		  ]
	}`)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		//panic(err)
		fmt.Println("Err: ", err.Error())
	}

	if response.StatusCode != http.StatusCreated {
		fmt.Println("StatusCode: ", response.StatusCode)
		fmt.Println("Status: ", response.Status)
		fmt.Println("Body: ", response.Body)
	}
	return true
}

func main() {
	Infof("Otel agent main func")
}
