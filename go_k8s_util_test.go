package go_k8s_util

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"
)

func TestSimpleNodeInfo(t *testing.T) {
	k8s, _ := NewK8s()
	k8s.SimpleNodeInfo("k8s1")
}

func TestSimp(t *testing.T) {
	cpuResource := resource.MustParse("100m") // 假设获得的 CPU 资源为 "100m"，即 100 毫核
	// 将 'resource.Quantity' 类型的资源分配给 'float64' 类型的变量
	cpuUsage := cpuResource.AsApproximateFloat64()
	fmt.Println(cpuUsage) // 输出：0.1，表示 0.1 核 CPU 资源的使用情况
}
