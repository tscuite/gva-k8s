package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GvaK8sNamespaceService struct {
}

func (GvaK8sNamespaceService *GvaK8sNamespaceService) Create(k8sNamespaces response.K8sNamespaces) (err error) {
	err = global.GVA_DB.Create(&k8sNamespaces).Error
	return err
}

func (GvaK8sNamespaceService *GvaK8sNamespaceService) GetList(k8sConf string) (list []*response.K8sNamespaces, total int64, err error) {
	// 初始化k8s客户端
	clientset, err := GvaK8sutils.GetK8sClient(k8sConf)
	if err != nil {
		fmt.Println("初始化k8s客户端失败: ", err.Error())
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("获取namespace失败: ", err.Error())
	}
	for key, ns := range namespaces.Items {
		//这样写本身就有问题，应该直接使用结构体传参，然后前端从中获取想要的数据，这样的话model结构体，以及表数据的生成要从新规划，同是表数据生成还有,api接口，vue菜单的预先定义也要写好
		creatTime := ns.CreationTimestamp.Time
		// 将time.Time类型转成指定时间格式
		formatTime := creatTime.Format("2006-01-02 15:04:05")
		res := &response.K8sNamespaces{
			ID:         uint(key),
			Namespace:  ns.Name,
			Status:     string(ns.Status.Phase),
			CreateTime: formatTime,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return list, total, nil
}
