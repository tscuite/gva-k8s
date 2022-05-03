package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/request"
	kubernetes "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GvaK8sNamespaceApi struct {
}

func (GvaK8sNamespaceApi *GvaK8sNamespaceApi) CreateK8sNamespaces(c *gin.Context) {
	var k8sNamespaces kubernetes.K8sNamespaces
	_ = c.ShouldBindJSON(&k8sNamespaces)
	if err := GvaK8sNamespaceService.Create(k8sNamespaces); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (GvaK8sNamespaceApi *GvaK8sNamespaceApi) GetK8sNamespacesList(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sNamespacesSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		fmt.Println("绑定查询失败: ", err.Error())
	}
	if list, total, err := GvaK8sNamespaceService.GetList(GvaK8sClusterService.ClusterID(c)); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
