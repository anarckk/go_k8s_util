package go_k8s_util

import (
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/ns"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/pod"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/yaml"
)

type K8sUtil struct {
	yaml.YamlUtil
	ns.NamespaceUtil
	pod.PodUtil
}

func NewK8s() (*K8sUtil, error) {
	var k8sUtil K8sUtil
	clientset, config, err := GetK8sClientset()
	if err != nil {
		return nil, err
	}
	k8sUtil.K8sConfig = config
	k8sUtil.YamlUtil.ClientSet = clientset
	k8sUtil.NamespaceUtil.ClientSet = clientset
	k8sUtil.PodUtil.ClientSet = clientset
	return &k8sUtil, nil
}
