package service

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesresponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GvaK8sPodService struct {
}

func (GvaK8sPodService *GvaK8sPodService) GetPodLog(k8sConfig string, k8sPods kubernetesresponse.K8sPods) (log string, err error) {
	// 初始化k8s客户端
	clientSet, _ := GvaK8sutils.GetK8sClient(k8sConfig)
	podLogOpts := v1.PodLogOptions{
		//Follow: true, // 类似kubectl logs -f命令
		TailLines: &k8sPods.Line,
	}
	req := clientSet.CoreV1().Pods(k8sPods.NameSpace).GetLogs(k8sPods.PodName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	log = buf.String()
	return log, nil
}

func (GvaK8sPodService *GvaK8sPodService) GetList(k8sConfig, namespace string) (list []*kubernetesresponse.K8sPods, total int64, err error) {
	// 初始化k8s客户端
	clientSet, _ := GvaK8sutils.GetK8sClient(k8sConfig)
	podList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	var c *gin.Context
	if err != nil {
		global.GVA_LOG.Error("获取pod失败", zap.Any("err", err))
		response.FailWithMessage("获取pod失败", c)
	}

	for key, pod := range podList.Items {
		var startTime string
		var restartCount int32
		if pod.Status.StartTime != nil {
			startTime = pod.Status.StartTime.Format("2006-01-02 15:04:05")
		} else {
			startTime = ""
		}
		if pod.Status.ContainerStatuses != nil {
			restartCount = pod.Status.ContainerStatuses[0].RestartCount
		}
		res := &kubernetesresponse.K8sPods{
			ID:           uint(key),
			PodName:      pod.Name,
			PodIP:        pod.Status.PodIP,
			HostIP:       pod.Status.HostIP,
			Status:       string(pod.Status.Phase),
			StartTime:    startTime,
			RestartCount: restartCount,
		}
		list = append(list, res)
	}
	total = int64(len(list))
	return list, total, nil
}
