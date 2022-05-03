package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/tscuite/gva-plugin/kubernetes/server/model/request"
	kubernetes "github.com/tscuite/gva-plugin/kubernetes/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags K8sDeployment
// @Summary 创建K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "创建K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/createK8sDeployment [post]
func (kubernetesApi *KubernetesApi) CreateK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetes.K8sDeployment
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := kubernetesService.CreateK8sDeployment(k8sDeployments); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 删除K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "删除K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployment [delete]
func (kubernetesApi *KubernetesApi) DeleteK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetes.K8sDeployment
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := kubernetesService.DeleteK8sDeployment(k8sDeployments); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 批量删除K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sDeployments/deleteK8sDeploymentByIds [delete]
func (kubernetesApi *KubernetesApi) DeleteK8sDeploymentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := kubernetesService.DeleteK8sDeploymentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 更新K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "更新K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sDeployments/updateK8sDeployment [put]
func (kubernetesApi *KubernetesApi) UpdateK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetes.K8sDeployment
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := kubernetesService.UpdateK8sDeployment(k8sDeployments); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 用id查询K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "用id查询K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sDeployments/findK8sDeployment [get]
func (kubernetesApi *KubernetesApi) FindK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetes.K8sDeployment
	_ = c.ShouldBindQuery(&k8sDeployments)
	if err, rek8sDeployments := kubernetesService.GetK8sDeployment(uint(k8sDeployments.ID)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rek8sDeployments": rek8sDeployments}, c)
	}
}

// @Tags K8sDeployment
// @Summary 重启Deployment
// @Description 重启应用接口：/k8sDeployments/restartK8sDeployment?namespace=xx&deployment=xx
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param namespace query string true "命名空间"
// @Param deployment query string true "应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重启应用成功"}"
// @Router /k8sDeployments/restartK8sDeployment [put]
func (kubernetesApi *KubernetesApi) RestartK8sDeployment(c *gin.Context) {
	// 指定Cluster
	kubeConfig, _ := kubernetesApi.ClusterID(c)
	namespace := c.Query("namespace")
	deployment := c.Query("deployment")
	if err := kubernetesService.RestartK8sDeployment(kubeConfig, namespace, deployment); err != nil {
		global.GVA_LOG.Error("重启应用失败", zap.Any("err", err))
		response.FailWithMessage("重启应用失败", c)
	} else {
		response.OkWithMessage("重启应用成功", c)
	}
}

// @Tags K8sDeployment
// @Summary 分页获取K8sDeployment列表
// @Description 默认获取集群ID为1的所有namespace的应用,获取指定集群的namespace的应用接口：/k8sDeployments/getK8sDeploymentList?clusterID={int}&namespace={string}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterID query uint true "集群ID" default(1)
// @Param namespace query string true "命名空间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/getK8sDeploymentList [get]
func (kubernetesApi *KubernetesApi) GetK8sDeploymentList(c *gin.Context) {
	// 指定Cluster
	kubeConfig, _ := kubernetesApi.ClusterID(c)
	var pageInfo kubernetesrequest.K8sDeploymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	namespace := c.Query("namespace")
	if err, list, total := kubernetesService.GetK8sDeploymentInfoList(kubeConfig, namespace); err != nil {
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
