package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/request"
	kubernetesresponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "k8s.io/api/core/v1"
)

type GvaK8sPodApi struct {
}

func (GvaK8sPodApi *GvaK8sPodApi) GetPodLog(c *gin.Context) {
	// 获取指定ClusterID的kubeconfig
	var k8sPods kubernetesresponse.K8sPods
	err := c.ShouldBindQuery(&k8sPods)
	if err != nil {
		fmt.Println("绑定查询参数失败: ", err.Error())
	}
	if rek8sPods, err := GvaK8sPodService.GetPodLog(GvaK8sClusterService.ClusterID(c), k8sPods); err != nil {
		global.GVA_LOG.Error("获取pod日志失败", zap.Any("err", err))
		response.FailWithMessage("获取pod日志失败", c)
	} else {
		response.OkWithData(gin.H{"rek8sPods": rek8sPods}, c)
	}
}

func (GvaK8sPodApi *GvaK8sPodApi) GetK8sPodsList(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sPodsSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		fmt.Println("绑定查询失败: ", err.Error())
	}
	namespace := c.Query("namespace")
	if list, total, err := GvaK8sPodService.GetList(GvaK8sClusterService.ClusterID(c), namespace); err != nil {
		global.GVA_LOG.Error("获取pod失败", zap.Any("err", err))
		response.FailWithMessage("获取pod失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取pod成功", c)
	}
}
