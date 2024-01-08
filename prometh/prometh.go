package prometh

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
)

func Push(pushAddr string, delay int64, net string) {
	jobName := fmt.Sprintf("filscan_delay_second")
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: jobName})
	gauge.Set(float64(delay))
	err := push.New(pushAddr, jobName).Grouping("module", "filscan").Grouping("net", net).Collector(gauge).Push()
	if err != nil {
		log.Printf("push prometheus %s failed:%s", pushAddr, err)
	}

}
