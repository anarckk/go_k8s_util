package yaml

import (
	"encoding/json"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func ConvObjToUnstructed(obj interface{}) (*unstructured.Unstructured, error) {
	unstucturedMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	unstructured := &unstructured.Unstructured{Object: unstucturedMap}
	return unstructured, nil
}

func ToPod(_unstructured *unstructured.Unstructured) (*corev1.Pod, error) {
	b, err := json.Marshal(_unstructured)
	if err != nil {
		return nil, err
	}
	pod := corev1.Pod{}
	err = json.Unmarshal(b, &pod)
	if err != nil {
		return nil, err
	}
	return &pod, nil
}

func ConvObjToPod(obj runtime.Object) (*corev1.Pod, error) {
	unstructuredObj, err := ConvObjToUnstructed(obj)
	if err != nil {
		return nil, err
	}
	pod, err := ToPod(unstructuredObj)
	if err != nil {
		return nil, err
	}
	return pod, nil
}

func ToStatefulSet(_unstructured *unstructured.Unstructured) (*appv1.StatefulSet, error) {
	b, err := json.Marshal(_unstructured)
	if err != nil {
		return nil, err
	}
	withStatus := appv1.StatefulSet{}
	err = json.Unmarshal(b, &withStatus)
	if err != nil {
		return nil, err
	}
	return &withStatus, nil
}

func ConvObjToStatefulSet(obj runtime.Object) (*appv1.StatefulSet, error) {
	unstructuredObj, err := ConvObjToUnstructed(obj)
	if err != nil {
		return nil, err
	}
	StatefulSet, err := ToStatefulSet(unstructuredObj)
	if err != nil {
		return nil, err
	}
	return StatefulSet, nil
}

func ToDeployment(_unstructured *unstructured.Unstructured) (*appv1.Deployment, error) {
	b, err := json.Marshal(_unstructured)
	if err != nil {
		return nil, err
	}
	withStatus := appv1.Deployment{}
	err = json.Unmarshal(b, &withStatus)
	if err != nil {
		return nil, err
	}
	return &withStatus, nil
}

func ConvObjToDeployment(obj runtime.Object) (*appv1.Deployment, error) {
	unstructuredObj, err := ConvObjToUnstructed(obj)
	if err != nil {
		return nil, err
	}
	deployment, err := ToDeployment(unstructuredObj)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}
