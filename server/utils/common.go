package utils

import (
	"fmt"

	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	GvaResponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type GvaK8sutils struct {
}

// 获取k8s Clientset
func (GvaK8sutils *GvaK8sutils) GetK8sClient(k8sConf string) (*kubernetes.Clientset, error) {
	var c *gin.Context
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf))
	if err != nil {
		global.GVA_LOG.Error("KubeConfig内容错误", zap.Any("err", err))
		response.FailWithMessage("KubeConfig内容错误", c)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.GVA_LOG.Error("创建Clientset失败", zap.Any("err", err))
		response.FailWithMessage("创建Clientset失败", c)
	}
	return clientSet, nil
}

// 获取k8s RESTConfig
func (GvaK8sutils *GvaK8sutils) GetRestConf(k8sConf string) (restConf *rest.Config, err error) {
	if restConf, err = clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf)); err != nil {
		fmt.Println("err: ", err)
	}
	return
}

func (GvaK8sutils *GvaK8sutils) Deployment(deployment GvaResponse.DeploymentSpec) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.Name,
			Namespace: deployment.Namespaces,
			Labels: map[string]string{
				"app": deployment.Name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: deployment.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deployment.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deployment.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: GvaK8sutils.Pod(deployment),
				},
			},
		},
	}
}
func (GvaK8sutils *GvaK8sutils) Pod(k8sDeployments GvaResponse.DeploymentSpec) (list []corev1.Container) {
	for _, Image := range k8sDeployments.Images {
		port := GvaK8sutils.Port(k8sDeployments, Image.Name)
		pod := corev1.Container{
			Name:  Image.Name,
			Image: Image.Images,
			Ports: port,
			LivenessProbe: &corev1.Probe{
				FailureThreshold:    5,
				InitialDelaySeconds: 60,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				TimeoutSeconds:      5,
				ProbeHandler: corev1.ProbeHandler{
					TCPSocket: &corev1.TCPSocketAction{
						Port: intstr.FromInt(int(Image.Check)),
					},
				},
			},
			ReadinessProbe: &corev1.Probe{
				FailureThreshold:    5,
				InitialDelaySeconds: 60,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				TimeoutSeconds:      5,
				ProbeHandler: corev1.ProbeHandler{
					TCPSocket: &corev1.TCPSocketAction{
						Port: intstr.FromInt(int(Image.Check)),
					},
				},
			},
			Resources: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(Image.LimitsCPU)),
					corev1.ResourceMemory: resource.MustParse(strconv.Itoa(Image.LimitsMemory) + "Mi"),
				},
				Requests: corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(Image.LimitsCPU)),
					corev1.ResourceMemory: resource.MustParse(strconv.Itoa(Image.LimitsMemory) + "Mi"),
				},
			},
			ImagePullPolicy: corev1.PullIfNotPresent,
		}
		list = append(list, pod)
	}
	return list
}

func (GvaK8sutils *GvaK8sutils) Port(k8sDeployments GvaResponse.DeploymentSpec, name string) (list []corev1.ContainerPort) {
	for _, Ports := range k8sDeployments.Port {
		cdz := corev1.ContainerPort{
			Name:          Ports.Name,
			ContainerPort: Ports.Port,
			Protocol:      "TCP",
		}
		if Ports.Images == name {
			list = append(list, cdz)
		}
	}
	return list
}
