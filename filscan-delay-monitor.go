package main

import (
	"filscan-delay-monitor/monitor"
	"flag"
	"fmt"
	"sync"
)

var pushAddr = flag.String("push-address", "", "Address of the Pushgateway to send metrics")
var interval = flag.Int("interval", 1, "Interval in minutes to check the URLs")
var mainURL = flag.String("mainnet-url", "", "URL for main-net")
var caliURL = flag.String("calibration-url", "", "URL for calibration")

func main() {
	flag.Parse()

	if *pushAddr == "" || *mainURL == "" || *caliURL == "" {
		fmt.Println("push-address, mainnet-url and calibration-url must be provided")
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		monitor.RunMonitorLoop(*mainURL, *pushAddr, "main-net", *interval)
	}()

	go func() {
		defer wg.Done()
		monitor.RunMonitorLoop(*caliURL, *pushAddr, "calibration", *interval)
	}()

	wg.Wait()
}
