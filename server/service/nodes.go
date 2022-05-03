package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GvaK8sNodeService struct {
}

func (GvaK8sNodeService *GvaK8sNodeService) GetList(k8sConf string) (list []*response.K8sNodes, total int64, err error) {
	// 初始化k8s客户端
	clientset, err := GvaK8sutils.GetK8sClient(k8sConf)
	if err != nil {
		fmt.Println("初始化k8s客户端失败: ", err)
	}
	// 获取所有节点
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("获取节点失败: ", err)
	}
	for key, node := range nodes.Items {
		creatTime := node.CreationTimestamp.Time
		// 将time.Time类型转成指定时间格式
		formatTime := creatTime.Format("2006-01-02 15:04:05")
		// 判断节点是否存在该key,存在则为master,反之为node
		var role string
		if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
			role = "master"
		} else {
			role = "node"
		}
		// 获取节点内存和cpu资源
		nodeMemory := node.Status.Capacity.Memory().Value()/1024/1024/1024 + 1
		resource := fmt.Sprintf("%sCore %dGB", node.Status.Capacity.Cpu().String(), nodeMemory)
		res := &response.K8sNodes{
			ID:         uint(key),
			NodeName:   node.Name,
			Version:    node.Status.NodeInfo.KubeletVersion,
			Status:     string(node.Status.Conditions[len(node.Status.Conditions)-1].Type),
			IP:         node.Status.Addresses[0].Address,
			Resource:   resource,
			Roles:      role,
			CreateTime: formatTime,
		}
		/*		for key, value := range node.Labels {
				fmt.Println("Key:",key, "Value:",value)
			}*/
		list = append(list, res)
	}
	total = int64(len(list))
	return list, total, nil
}
