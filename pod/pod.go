package pod

import (
	"context"
	"fmt"
	"strings"

	"gogs.bee.anarckk.me/anarckk/go_k8s_util/k8s_assist"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type PodUtil struct {
	ClientSet *kubernetes.Clientset
}

func (podUtil *PodUtil) PodIsExistsByLables(ctx context.Context, ns string, labels string) (bool, error) {
	_pods, err := podUtil.ListPodsByLabels(ctx, ns, labels)
	if err != nil {
		return false, err
	}
	return len(_pods) > 0, nil
}

func (podUtil *PodUtil) ListPods(ctx context.Context, ns string) ([]corev1.Pod, error) {
	podList, err := podUtil.ClientSet.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
}

func (podUtil *PodUtil) ListPodsByLabels(ctx context.Context, ns string, labels string) ([]corev1.Pod, error) {
	listOpts := metav1.ListOptions{}
	listOpts.LabelSelector = labels
	podList, err := podUtil.ClientSet.CoreV1().Pods(ns).List(ctx, listOpts)
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
}

func (podUtil *PodUtil) PodExistByName(ctx context.Context, ns string, podName string) (bool, error) {
	_, err := podUtil.GetPodByName(ctx, ns, podName)
	if err != nil {
		if k8s_assist.IsNotFound(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (podUtil *PodUtil) GetPodByName(ctx context.Context, ns string, podName string) (*corev1.Pod, error) {
	pod, err := podUtil.ClientSet.CoreV1().Pods(ns).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return pod, nil
}

func (podUtil *PodUtil) WatchPodsByOptions(ctx context.Context, ns string, opts metav1.ListOptions) (watch.Interface, error) {
	return podUtil.ClientSet.CoreV1().Pods(ns).Watch(ctx, opts)
}

func (podUtil *PodUtil) WatchPodsByLabels(ctx context.Context, ns string, label string) (watch.Interface, error) {
	return podUtil.ClientSet.CoreV1().Pods(ns).Watch(ctx, metav1.ListOptions{LabelSelector: label})
}

func (podUtil *PodUtil) WatchPodByName(ctx context.Context, ns string, name string) (watch.Interface, error) {
	return podUtil.ClientSet.CoreV1().Pods(ns).Watch(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name==%s", name)})
}

func (podUtil *PodUtil) CreatePod(ctx context.Context, ns string, manifest *corev1.Pod) (*corev1.Pod, error) {
	return podUtil.ClientSet.CoreV1().Pods(ns).Create(ctx, manifest, metav1.CreateOptions{})
}

func (podUtil *PodUtil) DeletePod(ctx context.Context, ns string, name string) error {
	return podUtil.ClientSet.CoreV1().Pods(ns).Delete(ctx, name, metav1.DeleteOptions{})
}

func (podUtil *PodUtil) GetPodsByNodeName(ctx context.Context, nodeName string) ([]*corev1.Pod, error) {
	fieldSelector, err := fields.ParseSelector(fmt.Sprintf("status.phase!=%s,status.phase!=%s,spec.nodeName=%s",
		string(corev1.PodSucceeded), string(corev1.PodFailed), nodeName))
	if err != nil {
		return nil, err
	}

	podList, err := podUtil.ClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{FieldSelector: fieldSelector.String()})
	if err != nil {
		return nil, err
	}

	var pods = make([]*corev1.Pod, 0)
	for i := range podList.Items {
		pods = append(pods, &podList.Items[i])
	}

	return pods, nil
}
