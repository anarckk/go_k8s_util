package yaml

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

type YamlUtil struct {
	K8sConfig *rest.Config
	ClientSet *kubernetes.Clientset
}

func (yamlUtil *YamlUtil) CreateResourceByUnstructured(ctx context.Context, ns string, unstructuredObj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	clientSet := yamlUtil.ClientSet
	k8sConfig := yamlUtil.K8sConfig

	gr, err := restmapper.GetAPIGroupResources(clientSet.Discovery())
	if err != nil {
		return nil, err
	}

	gvk := unstructuredObj.GetObjectKind().GroupVersionKind()
	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}

	dd, err := dynamic.NewForConfig(k8sConfig)
	if err != nil {
		return nil, err
	}

	var dri dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		if unstructuredObj.GetNamespace() == "" {
			unstructuredObj.SetNamespace(ns)
		}
		dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
	} else {
		dri = dd.Resource(mapping.Resource)
	}

	return_unstructured, err := dri.Create(ctx, unstructuredObj, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return return_unstructured, nil
}

func (yamlUtil *YamlUtil) DeleteResourceByUnstructured(ctx context.Context, ns string, unstructuredObj *unstructured.Unstructured) error {
	clientSet := yamlUtil.ClientSet
	k8sConfig := yamlUtil.K8sConfig

	dd, err := dynamic.NewForConfig(k8sConfig)
	if err != nil {
		return err
	}

	gvk := unstructuredObj.GetObjectKind().GroupVersionKind()

	gr, err := restmapper.GetAPIGroupResources(clientSet.Discovery())
	if err != nil {
		return err
	}
	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return err
	}

	var dri dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		if unstructuredObj.GetNamespace() == "" {
			unstructuredObj.SetNamespace(ns)
		}
		dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
	} else {
		dri = dd.Resource(mapping.Resource)
	}

	delOpts := metav1.DeleteOptions{}
	delOpts.Kind = unstructuredObj.GetKind()
	delOpts.APIVersion = unstructuredObj.GetAPIVersion()

	err = dri.Delete(ctx, unstructuredObj.GetName(), delOpts)
	if err != nil {
		return err
	}
	return nil
}

func (yamlUtil *YamlUtil) CreateResourceByYaml(ctx context.Context, ns string, yamlContent []byte) ([]*unstructured.Unstructured, error) {
	yamlArr, err := ParseYaml(yamlContent)
	if err != nil {
		return nil, err
	}
	var results []*unstructured.Unstructured
	for _, yamlU := range yamlArr {
		_unstructed, err := yamlUtil.CreateResourceByUnstructured(ctx, ns, yamlU.Unstructured)
		if err != nil {
			return nil, err
		}
		results = append(results, _unstructed)
	}
	return results, nil
}
