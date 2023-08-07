package ns

import (
	"context"
	"fmt"

	"gogs.bee.anarckk.me/anarckk/go_k8s_util/k8s_assist"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type NamespaceUtil struct {
	ClientSet *kubernetes.Clientset
}

func (nsUtil *NamespaceUtil) CreateNs(ctx context.Context, manifest *corev1.Namespace) (*corev1.Namespace, error) {
	return nsUtil.ClientSet.CoreV1().Namespaces().Create(ctx, manifest, metav1.CreateOptions{})
}

func (nsUtil *NamespaceUtil) CreateNsByName(ctx context.Context, ns string) (*corev1.Namespace, error) {
	manifest, err := nsUtil.ClientSet.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: ns,
		},
		Spec: corev1.NamespaceSpec{
			Finalizers: []corev1.FinalizerName{"kubernetes"},
		},
	}, metav1.CreateOptions{})
	return manifest, err
}

func (nsUtil *NamespaceUtil) NsExists(ctx context.Context, ns string) (bool, error) {
	_, err := nsUtil.GetNs(ctx, ns)
	if err != nil {
		if k8s_assist.IsNotFound(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (nsUtil *NamespaceUtil) GetNs(ctx context.Context, ns string) (*corev1.Namespace, error) {
	return nsUtil.ClientSet.CoreV1().Namespaces().Get(ctx, ns, metav1.GetOptions{})
}

func (nsUtil *NamespaceUtil) ListNs(ctx context.Context) ([]corev1.Namespace, error) {
	nsList, err := nsUtil.ClientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return nsList.Items, nil
}

func (nsUtil *NamespaceUtil) DeleteNs(ctx context.Context, ns string) error {
	return nsUtil.ClientSet.CoreV1().Namespaces().Delete(ctx, ns, metav1.DeleteOptions{})
}

func (nsUtil *NamespaceUtil) WatchNsByName(ctx context.Context, ns string) (watch.Interface, error) {
	return nsUtil.ClientSet.CoreV1().Namespaces().Watch(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name=%s", ns)})
}
