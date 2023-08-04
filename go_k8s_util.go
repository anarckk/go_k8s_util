package go_k8s_util

import "gogs.bee.anarckk.me/anarckk/go_k8s_util/pod"

type K8sUtil struct {
	pod.PodUtil
}

func NewK8s() (*K8sUtil, error) {
	var k8sUtil K8sUtil
	clientset, _, err := GetK8sClientset()
	if err != nil {
		return nil, err
	}
	k8sUtil.PodUtil.ClientSet = clientset
	return &k8sUtil, nil
}
