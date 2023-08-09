package metrics

import (
	"context"
	"log"
	"os"

	"gogs.bee.anarckk.me/anarckk/go_bit_util"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/connect"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/k8s_assist"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func SimpleMetrics() {
	kubeconfig, err := connect.GetKubeConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	metricClient := versioned.NewForConfigOrDie(config)
	k8s1Metrics, err := metricClient.MetricsV1beta1().NodeMetricses().Get(context.Background(), "k8s1", v1.GetOptions{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	c := k8s1Metrics.Usage.Cpu().AsApproximateFloat64()
	m := k8s_assist.GetResourceMemory(k8s1Metrics.Usage)
	log.Printf("k8s1 cpu: %f, memory: %d %s", c, m, go_bit_util.ByteCountBinary(m))
}
