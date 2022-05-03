package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GvaK8sNodeApi struct {
}

func (GvaK8sNodeApi *GvaK8sNodeApi) GetK8sNodesList(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sNodesSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		fmt.Println("绑定查询失败: ", err.Error())
	}
	if list, total, err := GvaK8sNodeService.GetList(GvaK8sClusterService.ClusterID(c)); err != nil {
		global.GVA_LOG.Error("获取节点失败", zap.Any("获取节点失败", err))
		response.FailWithMessage("获取节点失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取节点成功", c)
	}
}
