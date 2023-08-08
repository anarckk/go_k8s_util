package connect

import (
	"errors"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") //windows
}

// GetKubeConfig 获得kube配置文件的位置
func GetKubeConfig() (*string, error) {
	var kubeConfig *string
	home := homeDir()
	if home != "" {
		kubeConfig = new(string)
		*kubeConfig = filepath.Join(home, ".kube", "config")
		return kubeConfig, nil
	}
	return nil, errors.New("没有得到配置文件路径")
}

func GetK8sClientset() (*kubernetes.Clientset, *rest.Config, error) {
	var config *rest.Config
	kubeconfig, err := GetKubeConfig()
	if err == nil {
		// 首先使用 inCluster 模式(需要区配置对应的RBAC 权限,默认的sa是default-->是没有获取deployment的List权限)
		if config, err = rest.InClusterConfig(); err != nil {
			// 使用Kubeconfig文件配置集群配置Config对象
			if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
				return nil, nil, clientcmd.ErrEmptyCluster
			}
		}
		// 已经获得了rest.Config对象
		// 创建Clientset对象
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, nil, err
		}
		return clientset, config, nil
	}

	return nil, nil, err
}
