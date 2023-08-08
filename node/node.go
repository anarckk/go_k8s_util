package node

import (
	"context"
	"fmt"
	"log"

	"gogs.bee.anarckk.me/anarckk/go_bit_util"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/util"
	"gogs.bee.anarckk.me/anarckk/go_map_util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type NodeUtil struct {
	ClientSet *kubernetes.Clientset
}

func (nodeUtil *NodeUtil) ListNode(ctx context.Context) ([]corev1.Node, error) {
	nodeList, err := nodeUtil.ClientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return nodeList.Items, nil
}

func (nodeUtil *NodeUtil) ListNodeByLabels(ctx context.Context, label string) ([]corev1.Node, error) {
	nodeList, err := nodeUtil.ClientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return nil, err
	}
	return nodeList.Items, nil
}

func (nodeUtil *NodeUtil) GetNode(ctx context.Context, name string) (*corev1.Node, error) {
	return nodeUtil.ClientSet.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
}

func (nodeUtil *NodeUtil) WatchNode(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return nodeUtil.ClientSet.CoreV1().Nodes().Watch(ctx, opts)
}

func (nodeUtil *NodeUtil) WatchNodeByLabel(ctx context.Context, label string) (watch.Interface, error) {
	return nodeUtil.ClientSet.CoreV1().Nodes().Watch(ctx, metav1.ListOptions{LabelSelector: label})
}

func (nodeUtil *NodeUtil) WatchNodeByName(ctx context.Context, name string) (watch.Interface, error) {
	return nodeUtil.ClientSet.CoreV1().Nodes().Watch(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name=%s", name)})
}

func GetAllocatableCpu(node *corev1.Node) int64 {
	res := node.Status.Allocatable[corev1.ResourceCPU]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func GetAllocatableMemory(node *corev1.Node) int64 {
	res := node.Status.Allocatable[corev1.ResourceMemory]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func GetAllocatableEphemeralStorage(node *corev1.Node) int64 {
	res := node.Status.Allocatable[corev1.ResourceEphemeralStorage]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func GetCapacityCpu(node *corev1.Node) int64 {
	res := node.Status.Capacity["cpu"]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func GetCapacityMemory(node *corev1.Node) int64 {
	res := node.Status.Capacity["memory"]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func GetCapacityEphemeralStorage(node *corev1.Node) int64 {
	res := node.Status.Capacity["ephemeral-storage"]
	i64, b := res.AsInt64()
	if !b {
		log.Println("occur error", res)
		return 0
	}
	return i64
}

func SimpleNode(node *corev1.Node) {
	log.Printf("name: %s\n", node.Name)
	log.Printf("labels: %s\n", util.ComposeMap(node.Labels))
	log.Printf("address: %s\n", go_map_util.ComposeStrArray2(node.Status.Addresses, func(a corev1.NodeAddress) string {
		return a.Address
	}))
	log.Printf("allocatable cpu: %d\n", GetAllocatableCpu(node))
	log.Printf("allocatable memory: %d %s\n", GetAllocatableMemory(node), go_bit_util.ByteCountBinary(GetAllocatableMemory(node)))
	log.Printf("allocatable disk: %d %s\n", GetAllocatableEphemeralStorage(node), go_bit_util.ByteCountBinary(GetAllocatableEphemeralStorage(node)))
	log.Printf("capacity cpu: %d\n", GetCapacityCpu(node))
	log.Printf("capacity memory: %d %s\n", GetCapacityMemory(node), go_bit_util.ByteCountBinary(GetCapacityMemory(node)))
	log.Printf("capacity disk: %d %s\n", GetCapacityEphemeralStorage(node), go_bit_util.ByteCountBinary(GetCapacityEphemeralStorage(node)))
}
