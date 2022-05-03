package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/tscuite/gva-plugin/kubernetes/server/model/response"
	"github.com/tscuite/gva-plugin/kubernetes/server/utils"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/types"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubernetesService struct {
}

//@function: CreateK8sDeployment
//@description: 创建K8sDeployment记录
//@param: k8sDeployments model.K8sDeployment
//@return: err error
func (kubernetesService *KubernetesService) CreateK8sDeployment(k8sDeployments response.K8sDeployment) (err error) {
	err = global.GVA_DB.Create(&k8sDeployments).Error
	return err
}

//@function: DeleteK8sDeployment
//@description: 删除K8sDeployment记录
//@param: k8sDeployments model.K8sDeployment
//@return: err error
func (kubernetesService *KubernetesService) DeleteK8sDeployment(k8sDeployments response.K8sDeployment) (err error) {
	err = global.GVA_DB.Delete(k8sDeployments).Error
	return err
}

//@function: DeleteK8sDeploymentByIds
//@description: 批量删除K8sDeployment记录
//@param: ids request.IdsReq
//@return: err error
func (kubernetesService *KubernetesService) DeleteK8sDeploymentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]response.K8sDeployment{}, "id in ?", ids.Ids).Error
	return err
}

//@function: UpdateK8sDeployment
//@description: 更新K8sDeployment记录
//@param: k8sDeployments *model.K8sDeployment
//@return: err error
func (kubernetesService *KubernetesService) UpdateK8sDeployment(k8sDeployments response.K8sDeployment) (err error) {
	err = global.GVA_DB.Save(&k8sDeployments).Error
	return err
}

//@function: GetK8sDeployment
//@description: 根据id获取K8sDeployment记录
//@param: id uint
//@return: err error, k8sDeployments model.K8sDeployment
func (kubernetesService *KubernetesService) GetK8sDeployment(id uint) (err error, k8sDeployments response.K8sDeployment) {
	err = global.GVA_DB.Where("id = ?", id).First(&k8sDeployments).Error
	return
}

//@function: RestartK8sDeployment
//@description: 重启Deployment
//@param: namespace string,deployment string
//@return: err error, k8sDeployments model.K8sDeployment
func (kubernetesService *KubernetesService) RestartK8sDeployment(k8sConf, namespace, deployment string) (err error) {
	//初始化k8s客户端
	clientset, _ := utils.GetK8sClient(k8sConf)
	//重启指定namespace的deployment
	data := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"kubectl.kubernetes.io/restartedAt":"%s"}}}}}`, time.Now().String())
	if _, err := clientset.AppsV1().Deployments(namespace).Patch(context.Background(), deployment, types.StrategicMergePatchType, []byte(data), metav1.PatchOptions{FieldManager: "kubectl-rollout"}); err != nil {
		global.GVA_LOG.Error("deployment重启失败", zap.Any("err", err))
	} else {
		global.GVA_LOG.Info("deployment重启成功", zap.Any("success", deployment))
	}
	return nil
}

//@function: GetK8sDeploymentInfoList
//@description: 分页获取K8sDeployment列表
//@param: info request.K8sDeploymentSearch
//@return: err error, list []*model.K8sDeployment, total int64
func (kubernetesService *KubernetesService) GetK8sDeploymentInfoList(k8sConf, namespace string) (err error, list []*response.K8sDeployment, total int64) {
	//初始化k8s客户端
	clientset, _ := utils.GetK8sClient(k8sConf)
	//获取指定namespace的deployment
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for key, deployment := range deployments.Items {
		createTime := deployment.ObjectMeta.CreationTimestamp.Time
		// 将time.Time类型转成指定时间格式
		formatTime := createTime.Format("2006-01-02 15:04:05")
		res := &response.K8sDeployment{
			ID:         uint(key),
			Namespace:  namespace,
			Deployment: deployment.ObjectMeta.Name,
			Replicas:   deployment.Status.ReadyReplicas,
			CreateTime: formatTime,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return err, list, total
}
