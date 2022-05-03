package service

import (
	"context"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"
	"go.uber.org/zap"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GvaK8sDeploymentService struct {
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) Create(k8sConf string, k8sDeployments response.DeploymentSpec) (err error) {
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
	deployment, err := clientset.AppsV1().Deployments(k8sDeployments.Namespaces).Create(context.TODO(), GvaK8sutils.Deployment(k8sDeployments), metav1.CreateOptions{
		TypeMeta:        metav1.TypeMeta{},
		DryRun:          []string{},
		FieldManager:    "",
		FieldValidation: "",
	})
	global.GVA_LOG.Info("deployment创建成功!", zap.Any("deployment", deployment.Name))
	return err
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) Delete(k8sConf string, k8sDeployments response.DeploymentSpec) (err error) {
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
	err = clientset.AppsV1().Deployments(k8sDeployments.Namespaces).Delete(context.TODO(), k8sDeployments.Name, metav1.DeleteOptions{})
	return err
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) DeleteId(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]response.K8sDeployment{}, "id in ?", ids.Ids).Error
	return err
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) Update(k8sConf string, k8sDeployments response.DeploymentSpec) (err error) {
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
	deployment, err := clientset.AppsV1().Deployments(k8sDeployments.Namespaces).Update(context.TODO(), GvaK8sutils.Deployment(k8sDeployments), metav1.UpdateOptions{
		TypeMeta:        metav1.TypeMeta{},
		DryRun:          []string{},
		FieldManager:    "",
		FieldValidation: "",
	})
	global.GVA_LOG.Info("deployment更新成功!", zap.Any("deployment", deployment.Name))
	return err
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) Get(k8sConf string, k8sDeployment response.DeploymentSpec) (k8sDeployments response.DeploymentSpec, err error) {
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
	deployment, err := clientset.AppsV1().Deployments(k8sDeployment.Namespaces).Get(context.TODO(), k8sDeployment.Name, metav1.GetOptions{})
	list, port := GvaK8sDeploymentService.cdz(deployment)
	res := &response.DeploymentSpec{
		Name:       deployment.Name,
		Namespaces: deployment.ObjectMeta.Namespace,
		Replicas:   deployment.Spec.Replicas,
		Images:     list,
		Port:       port,
	}
	return *res, err
}
func (GvaK8sDeploymentService *GvaK8sDeploymentService) cdz(deployment *v1.Deployment) (list []*response.Images, port []*response.Port) {
	for _, Image := range deployment.Spec.Template.Spec.Containers {
		res := &response.Images{
			Name:           Image.Name,
			Images:         Image.Image,
			Check:          Image.LivenessProbe.TCPSocket.Port.IntValue(),
			LimitsCPU:      int(Image.Resources.Limits.Cpu().Value()),
			LimitsMemory:   int(Image.Resources.Limits.Memory().Value()) / 1048576,
			RequestsCPU:    int(Image.Resources.Requests.Cpu().Value()),
			RequestsMemory: int(Image.Resources.Requests.Memory().Value()) / 1048576,
		}
		list = append(list, res)
		for _, Ports := range Image.Ports {
			cdz := &response.Port{
				Name:   Ports.Name,
				Images: Image.Name,
				Port:   Ports.ContainerPort,
			}
			port = append(port, cdz)
		}
	}
	return list, port
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) GetList(k8sConf, namespace string) (list []*response.K8sDeployment, total int64, err error) {
	//初始化k8s客户端
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
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
			Namespace:  deployment.ObjectMeta.Namespace,
			Deployment: deployment.ObjectMeta.Name,
			Replicas:   deployment.Status.ReadyReplicas,
			CreateTime: formatTime,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return list, total, err
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) GetConfigMap(k8sConf, namespace string) (list []*response.K8sConfigMap, total int64, err error) {
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
	configMaps, err := clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get config map:", err)
	}
	for key, configMap := range configMaps.Items {
		createTime := configMap.ObjectMeta.CreationTimestamp.Time
		// 将time.Time类型转成指定时间格式
		formatTime := createTime.Format("2006-01-02 15:04:05")
		res := &response.K8sConfigMap{
			ID:         uint(key),
			Namespaces: configMap.ObjectMeta.Namespace,
			ConfigMap:  configMap.ObjectMeta.Name,
			Config:     configMap.Data,
			CreateTime: formatTime,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return list, total, err
}

func (GvaK8sDeploymentService *GvaK8sDeploymentService) FindGetConfigMap(k8sConf string, k8sDeployment response.K8sConfigMap) (k8sDeployments response.K8sConfigMap, err error) {
	clientset, _ := GvaK8sutils.GetK8sClient(k8sConf)
	ConfigMap, err := clientset.CoreV1().ConfigMaps(k8sDeployment.Namespaces).Get(context.TODO(), k8sDeployment.ConfigMap, metav1.GetOptions{})
	res := &response.K8sConfigMap{
		ConfigMap:  ConfigMap.Name,
		Namespaces: ConfigMap.ObjectMeta.Namespace,
		Config:     ConfigMap.Data,
	}
	return *res, err
}
