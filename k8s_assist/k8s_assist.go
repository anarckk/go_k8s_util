package k8s_assist

import (
	"log"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

func IsNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found")
}

func IsAlreadyExists(err error) bool {
	return strings.Contains(err.Error(), "already exists")
}

func GetResourceCpu(resList corev1.ResourceList) float64 {
	res := resList[corev1.ResourceCPU]
	return res.AsApproximateFloat64()
}

func GetResourceMemory(resList corev1.ResourceList) int64 {
	res := resList[corev1.ResourceMemory]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func GetResourceEphemeralStorage(resList corev1.ResourceList) int64 {
	res := resList[corev1.ResourceEphemeralStorage]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}
