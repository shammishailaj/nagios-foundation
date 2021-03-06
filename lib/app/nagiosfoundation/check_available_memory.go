// +build windows linux darwin

package nagiosfoundation

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jkerry/nagiosfoundation/lib/pkg/memory"
	"github.com/jkerry/nagiosfoundation/lib/pkg/nagiosformatters"
)

func CheckAvailableMemory() {
	var warning = flag.Float64("warning", 85, "the memory threshold to issue a warning alert")
	var critical = flag.Float64("critical", 95, "the memory threshold to issue a critical alert")
	var metricName = flag.String("metric_name", "available_memory_mbs", "the name of the metric generated by this check")
	flag.Parse()
	freememory, err := memory.GetFreeMemory()
	if err == nil {
		var msg string
		var retcode int
		msg, retcode = nagiosformatters.LesserFormatNagiosCheck("CheckAvailableMemoryMbytes", freememory, *warning, *critical, *metricName)
		fmt.Println(msg)
		os.Exit(retcode)
	} else {
		log.Println(err)
		os.Exit(3)
	}
}
