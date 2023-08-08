package node

import (
	"context"
	"log"
	"testing"

	"gogs.bee.anarckk.me/anarckk/go_k8s_util/connect"
)

func TestListNode(t *testing.T) {
	clientset, _, err := connect.GetK8sClientset()
	if err != nil {
		t.Error(err)
	}
	nodeUtil := &NodeUtil{clientset}
	nodes, err := nodeUtil.ListNode(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, n := range nodes {
		log.Println(n.Name)
	}
}

func getNode(t *testing.T) *NodeUtil {
	clientset, _, err := connect.GetK8sClientset()
	if err != nil {
		t.Error(err)
	}
	return &NodeUtil{clientset}
}
