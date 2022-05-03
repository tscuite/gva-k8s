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

// @Tags K8sNamespaces
// @Summary 创建K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "创建K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/createK8sNamespaces [post]
func (kubernetesApi *KubernetesApi) CreateK8sNamespaces(c *gin.Context) {
	var k8sNamespaces kubernetes.K8sNamespaces
	_ = c.ShouldBindJSON(&k8sNamespaces)
	if err := kubernetesService.CreateK8sNamespaces(k8sNamespaces); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags K8sNamespaces
// @Summary 删除K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "删除K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNamespaces/deleteK8sNamespaces [delete]
func (kubernetesApi *KubernetesApi) DeleteK8sNamespaces(c *gin.Context) {
	var k8sNamespaces kubernetes.K8sNamespaces
	_ = c.ShouldBindJSON(&k8sNamespaces)
	if err := kubernetesService.DeleteK8sNamespaces(k8sNamespaces); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags K8sNamespaces
// @Summary 批量删除K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sNamespaces/deleteK8sNamespacesByIds [delete]
func (kubernetesApi *KubernetesApi) DeleteK8sNamespacesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := kubernetesService.DeleteK8sNamespacesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags K8sNamespaces
// @Summary 更新K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "更新K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sNamespaces/updateK8sNamespaces [put]
func (kubernetesApi *KubernetesApi) UpdateK8sNamespaces(c *gin.Context) {
	var k8sNamespaces kubernetes.K8sNamespaces
	_ = c.ShouldBindJSON(&k8sNamespaces)
	if err := kubernetesService.UpdateK8sNamespaces(k8sNamespaces); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sNamespaces
// @Summary 用id查询K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sNamespaces true "用id查询K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNamespaces/findK8sNamespaces [get]
func (kubernetesApi *KubernetesApi) FindK8sNamespaces(c *gin.Context) {
	var k8sNamespaces kubernetes.K8sNamespaces
	_ = c.ShouldBindQuery(&k8sNamespaces)
	if err, rek8sNamespaces := kubernetesService.GetK8sNamespaces(uint(k8sNamespaces.ID)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rek8sNamespaces": rek8sNamespaces}, c)
	}
}

// @Tags K8sNamespaces
// @Summary 获取指定集群的Namespace
// @Description 默认获取集群ID为1的namespace，获取指定集群的namespace接口：/k8sNamespaces/getK8sNamespacesList?clusterID={int}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterID query uint true "集群ID" default(1)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/getK8sNamespacesList [get]
func (kubernetesApi *KubernetesApi) GetK8sNamespacesList(c *gin.Context) {
	kubeConfig, err := kubernetesApi.ClusterID(c)
	if err != nil {
		fmt.Println("获取kubeConfig失败: ", err.Error())
	}
	var pageInfo kubernetesrequest.K8sNamespacesSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		fmt.Println("绑定查询失败: ", err.Error())
	}
	if err, list, total := kubernetesService.GetK8sNamespacesInfoList(kubeConfig); err != nil {
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
