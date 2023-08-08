package go_k8s_util

import (
	"context"
	"log"

	"gogs.bee.anarckk.me/anarckk/go_bit_util"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/connect"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/k8s_assist"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/node"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/ns"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/pod"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/util"
	"gogs.bee.anarckk.me/anarckk/go_k8s_util/yaml"
	"gogs.bee.anarckk.me/anarckk/go_map_util"
	corev1 "k8s.io/api/core/v1"
)

type K8sUtil struct {
	yaml.YamlUtil
	ns.NamespaceUtil
	pod.PodUtil
	node.NodeUtil
}

func NewK8s() (*K8sUtil, error) {
	var k8sUtil K8sUtil
	clientset, config, err := connect.GetK8sClientset()
	if err != nil {
		return nil, err
	}
	k8sUtil.K8sConfig = config
	k8sUtil.YamlUtil.ClientSet = clientset
	k8sUtil.NamespaceUtil.ClientSet = clientset
	k8sUtil.PodUtil.ClientSet = clientset
	k8sUtil.NodeUtil.ClientSet = clientset
	return &k8sUtil, nil
}

func (k8sUtil *K8sUtil) SimpleNodeInfo(nodeName string) {
	_node, err := k8sUtil.GetNode(context.Background(), nodeName)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("name: %s\n", _node.Name)
	log.Printf("labels: %s\n", util.ComposeMap(_node.Labels))
	log.Printf("address: %s\n", go_map_util.ComposeStrArray2(_node.Status.Addresses, func(a corev1.NodeAddress) string {
		return a.Address
	}))
	log.Printf("allocatable cpu: %f\n", node.GetAllocatableCpu(_node))
	log.Printf("allocatable memory: %d %s\n", node.GetAllocatableMemory(_node), go_bit_util.ByteCountBinary(node.GetAllocatableMemory(_node)))
	log.Printf("allocatable disk: %d %s\n", node.GetAllocatableEphemeralStorage(_node), go_bit_util.ByteCountBinary(node.GetAllocatableEphemeralStorage(_node)))
	log.Printf("capacity cpu: %f\n", node.GetCapacityCpu(_node))
	log.Printf("capacity memory: %d %s\n", node.GetCapacityMemory(_node), go_bit_util.ByteCountBinary(node.GetCapacityMemory(_node)))
	log.Printf("capacity disk: %d %s\n", node.GetCapacityEphemeralStorage(_node), go_bit_util.ByteCountBinary(node.GetCapacityEphemeralStorage(_node)))

	pods, err := k8sUtil.GetPodsByNodeName(context.Background(), nodeName)
	if err != nil {
		log.Println(err)
		return
	}
	var cpu float64 = 0
	var memory int64 = 0
	var disk int64 = 0
	for _, p := range pods {
		for _, c := range p.Spec.Containers {
			c1 := k8s_assist.GetResourceCpu2(c.Resources.Requests)
			m1 := k8s_assist.GetResourceMemory(c.Resources.Requests)
			d1 := k8s_assist.GetResourceEphemeralStorage(c.Resources.Requests)
			cpu += c1
			memory += m1
			disk += d1
			log.Printf("--- %s allocated cpu: %f\n", c.Name, c1)
			log.Printf("--- %s allocated memory: %d %s\n", c.Name, m1, go_bit_util.ByteCountBinary(m1))
			log.Printf("--- %s allocated disk: %d %s\n", c.Name, d1, go_bit_util.ByteCountBinary(d1))

		}
	}
	// 这个值和dashboard里的node里的下限一样，也就是说下限宿主机中所有pod请求的和
	log.Printf("allocated cpu: %f\n", cpu)
	log.Printf("allocated memory: %d %s\n", memory, go_bit_util.ByteCountBinary(memory))
	log.Printf("allocated disk: %d %s\n", disk, go_bit_util.ByteCountBinary(disk))
}
