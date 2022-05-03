package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/tscuite/gva-plugin/kubernetes/server/model/response"
	"github.com/tscuite/gva-plugin/kubernetes/server/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//@function: CreateK8sNamespaces
//@description: 创建K8sNamespaces记录
//@param: k8sNamespaces model.K8sNamespaces
//@return: err error

func (kubernetesService *KubernetesService) CreateK8sNamespaces(k8sNamespaces response.K8sNamespaces) (err error) {
	err = global.GVA_DB.Create(&k8sNamespaces).Error
	return err
}

//@function: DeleteK8sNamespaces
//@description: 删除K8sNamespaces记录
//@param: k8sNamespaces model.K8sNamespaces
//@return: err error

func (kubernetesService *KubernetesService) DeleteK8sNamespaces(k8sNamespaces response.K8sNamespaces) (err error) {
	err = global.GVA_DB.Delete(k8sNamespaces).Error
	return err
}

//@function: DeleteK8sNamespacesByIds
//@description: 批量删除K8sNamespaces记录
//@param: ids request.IdsReq
//@return: err error

func (kubernetesService *KubernetesService) DeleteK8sNamespacesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]response.K8sNamespaces{}, "id in ?", ids.Ids).Error
	return err
}

//@function: UpdateK8sNamespaces
//@description: 更新K8sNamespaces记录
//@param: k8sNamespaces *model.K8sNamespaces
//@return: err error

func (kubernetesService *KubernetesService) UpdateK8sNamespaces(k8sNamespaces response.K8sNamespaces) (err error) {
	err = global.GVA_DB.Save(&k8sNamespaces).Error
	return err
}

//@function: GetK8sNamespaces
//@description: 根据id获取K8sNamespaces记录
//@param: id uint
//@return: err error, k8sNamespaces model.K8sNamespaces

func (kubernetesService *KubernetesService) GetK8sNamespaces(id uint) (err error, k8sNamespaces response.K8sNamespaces) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sNamespaces).Error
	return
}

//@function: GetK8sNamespacesInfoList
//@description: 分页获取K8sNamespaces
//@param: info request.K8sNamespacesSearch
//@return: err error, list []*model.K8sNamespaces, total int64

func (kubernetesService *KubernetesService) GetK8sNamespacesInfoList(k8sConf string) (err error, list []*response.K8sNamespaces, total int64) {
	// 初始化k8s客户端
	clientset, err := utils.GetK8sClient(k8sConf)
	if err != nil {
		fmt.Println("初始化k8s客户端失败: ", err.Error())
	}

	// 获取所有Namespaces
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
	return nil, list, total
}
