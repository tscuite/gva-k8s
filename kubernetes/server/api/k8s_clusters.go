package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/tscuite/gva-plugin/kubernetes/server/model/request"
	kubernetes "github.com/tscuite/gva-plugin/kubernetes/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KubernetesclustersApi struct {
}

// @Tags K8sCluster
// @Summary 创建K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "创建K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sCluster/createK8sCluster [post]
func (kubernetesclustersApi *KubernetesclustersApi) CreateK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := kubernetesclustersService.CreateK8sCluster(K8sCluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags K8sCluster
// @Summary 删除K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "删除K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sCluster/deleteK8sCluster [delete]
func (kubernetesclustersApi *KubernetesclustersApi) DeleteK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := kubernetesclustersService.DeleteK8sCluster(&K8sCluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags K8sCluster
// @Summary 批量删除K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sCluster/deleteK8sClusterByIds [delete]
func (kubernetesclustersApi *KubernetesclustersApi) DeleteK8sClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := kubernetesclustersService.DeleteK8sClusterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags K8sCluster
// @Summary 更新K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "更新K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sCluster/updateK8sCluster [put]
func (kubernetesclustersApi *KubernetesclustersApi) UpdateK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := kubernetesclustersService.UpdateK8sCluster(&K8sCluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sCluster
// @Summary 用id查询K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "用id查询K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sCluster/findK8sCluster [get]
func (kubernetesclustersApi *KubernetesclustersApi) FindK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindQuery(&K8sCluster)
	token := c.Query("ID")
	global.GVA_LOG.Info("你的webhook正常运行!", zap.Any("查看token", K8sCluster.ID))
	global.GVA_LOG.Info("你的webhook正常运行!", zap.Any("查看token", token))
	if err, reK8sCluster := kubernetesclustersService.GetK8sCluster(K8sCluster.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reK8sCluster": reK8sCluster}, c)
	}
}

// @Tags K8sCluster
// @Summary 分页获取K8sCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sClusterSearch true "分页获取K8sCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sCluster/getK8sClusterList [get]
func (kubernetesclustersApi *KubernetesclustersApi) GetK8sClusterList(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sClusterSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		fmt.Println("绑定查询失败: ", err.Error())
	}
	if err, list, total := kubernetesclustersService.GetK8sClusterInfoList(pageInfo); err != nil {
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
