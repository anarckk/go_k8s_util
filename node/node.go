package node

import (
	"context"
	"fmt"

	"gogs.bee.anarckk.me/anarckk/go_k8s_util/k8s_assist"
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

func GetAllocatableCpu(node *corev1.Node) float64 {
	return k8s_assist.GetResourceCpu(node.Status.Allocatable)
}

func GetAllocatableMemory(node *corev1.Node) int64 {
	return k8s_assist.GetResourceMemory(node.Status.Allocatable)
}

func GetAllocatableEphemeralStorage(node *corev1.Node) int64 {
	return k8s_assist.GetResourceEphemeralStorage(node.Status.Allocatable)
}

func GetCapacityCpu(node *corev1.Node) float64 {
	return k8s_assist.GetResourceCpu(node.Status.Capacity)
}

func GetCapacityMemory(node *corev1.Node) int64 {
	return k8s_assist.GetResourceMemory(node.Status.Capacity)
}

func GetCapacityEphemeralStorage(node *corev1.Node) int64 {
	return k8s_assist.GetResourceEphemeralStorage(node.Status.Capacity)
}
