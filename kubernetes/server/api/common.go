package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KubernetesApi struct {
}

// 公共方法, 获取指定k8s集群的kubeconfig
func (kubernetesApi *KubernetesApi) ClusterID(c *gin.Context) (string, error) {
	clusterID := c.DefaultQuery("clusterID", "1")
	clusterIDuint, err := strconv.ParseUint(clusterID, 10, 32)
	err, cluster := kubernetesclustersService.GetK8sCluster(uint(clusterIDuint))
	if err != nil {
		global.GVA_LOG.Error("获取集群失败", zap.Any("err", err))
		response.FailWithMessage("获取集群失败", c)
	}
	return cluster.KubeConfig, nil
}
