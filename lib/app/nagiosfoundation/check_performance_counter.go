// +build windows

package nagiosfoundation

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jkerry/nagiosfoundation/lib/pkg/nagiosformatters"
	"github.com/jkerry/nagiosfoundation/lib/pkg/perfcounters"
)

func CheckPerformanceCounter() {
	var warning = flag.Float64("warning", 0, "the threshold to issue a warning alert")
	var critical = flag.Float64("critical", 0, "the threshold to issue a critical alert")
	var greaterThan = flag.Bool("greater_than", false, "issue warnings if the metric is greater than the expected thresholds")
	var pollingAttempts = flag.Int("polling_attempts", 2, "the number of times to fetch and average the performance counter")
	var pollingDelay = flag.Int("polling_delay", 1, "the number of seconds to delay between polling attempts")
	var metricName = flag.String("metric_name", "", "the name of the metric generated by this check")
	var counterName = flag.String("counter_name", "", "the name of the performance counter to check")
	flag.Parse()
	counter, err := perfcounters.ReadPerformanceCounter(*counterName, *pollingAttempts, *pollingDelay)
	if err == nil {
		var msg string
		var retcode int
		if *greaterThan {
			msg, retcode = nagiosformatters.GreaterFormatNagiosCheck(*counterName, counter.Value, *warning, *critical, *metricName)
			fmt.Println(msg)
			os.Exit(retcode)
		}
		msg, retcode = nagiosformatters.LesserFormatNagiosCheck(*counterName, counter.Value, *warning, *critical, *metricName)
		fmt.Println(msg)
		os.Exit(retcode)
	} else {
		log.Println(err)
		os.Exit(3)
	}
}
