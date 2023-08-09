package metrics

import (
	"testing"
	"time"
)

func TestSimpleMetrics(t *testing.T) {
	SimpleMetrics()
}

func TestSimpleMetricsLoop(t *testing.T) {
	for {
		SimpleMetrics()
		time.Sleep(time.Second)
	}
}
