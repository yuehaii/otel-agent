package main

import (
	"strconv"

	otelfmt "github.com/yuehaii/otel-agent/fmt"
	otellog "github.com/yuehaii/otel-agent/logrus"
)

func main() {
	// multiple threads stress test
	for i := 0; i < 2; i++ {
		//go otellog.Fatalf("lablador log record: %s", strconv.Itoa(i)+"th image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		//go otellog.Panicf("lablador log record: %s", strconv.Itoa(i)+"th image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")

		otellog.Infof("lablador log record: %s", strconv.Itoa(i)+"Infof image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		go otellog.Tracef("lablador log record: %s", strconv.Itoa(i)+"Tracef image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		go otellog.Debugf("lablador log record: %s", strconv.Itoa(i)+"Debugf image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		go otellog.Warnf("lablador log record: %s", strconv.Itoa(i)+"Warnf image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		go otellog.Errorf("lablador log record: %s", strconv.Itoa(i)+"Errorf image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		go otellog.Printf("lablador log record: %s", strconv.Itoa(i)+"Printf image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")

		otelfmt.Sprintf("lablador log record: %s", strconv.Itoa(i)+"Sprintf image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")

		go otelfmt.Errorf("lablador log record: %s", strconv.Itoa(i)+"Errorf2 image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")
		go otelfmt.Println(strconv.Itoa(i) + "Println image - after last shop front series I imagined IKEA shop front in beautiful lane of Italy or Spain :) it will be one of a kind #unique #bedifferent")

	}
	channel := make(chan int)
	channel <- 1
}
