package metrics

import (
	"context"
	"log"
	"testing"
	"time"

	"gogs.bee.anarckk.me/anarckk/go_bit_util"
)

func TestSimpleMetrics(t *testing.T) {
	SimpleMetrics()
}

func TestGetCpuMemory(t *testing.T) {
	metricsUtil, err := NewMetricsUtil()
	if err != nil {
		t.Error(err)
	}
	c, m, e := metricsUtil.GetCpuMemory(context.Background(), "k8s1")
	if e != nil {
		t.Error(err)
	}
	log.Printf("k8s1 cpu: %f, memory: %d %s", c, m, go_bit_util.ByteCountBinary(m))
	time.Sleep(time.Second)
}
