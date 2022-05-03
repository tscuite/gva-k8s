package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/request"
	kubernetesresponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GvaK8sDeploymentApi struct {
}

// @Tags K8sDeployment
// @Summary 创建K8sDeployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sDeployment true "创建K8sDeployment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/createK8sDeployment [post]
func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) CreateK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetesresponse.DeploymentSpec
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := GvaK8sDeploymentService.Create(GvaK8sClusterService.ClusterID(c), k8sDeployments); err != nil {
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
func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) DeleteK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetesresponse.DeploymentSpec
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := GvaK8sDeploymentService.Delete(GvaK8sClusterService.ClusterID(c), k8sDeployments); err != nil {
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
// func (kubernetesApi *KubernetesApi) DeleteK8sDeploymentByIds(c *gin.Context) {
// 	var IDS request.IdsReq
// 	_ = c.ShouldBindJSON(&IDS)
// 	if err := kubernetesService.DeleteK8sDeploymentByIds(IDS); err != nil {
// 		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
// 		response.FailWithMessage("批量删除失败", c)
// 	} else {
// 		response.OkWithMessage("批量删除成功", c)
// 	}
// }

func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) UpdateK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetesresponse.DeploymentSpec
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := GvaK8sDeploymentService.Update(GvaK8sClusterService.ClusterID(c), k8sDeployments); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) FindK8sDeployment(c *gin.Context) {
	var k8sDeployments kubernetesresponse.DeploymentSpec
	_ = c.ShouldBindQuery(&k8sDeployments)
	if rek8sDeployments, err := GvaK8sDeploymentService.Get(GvaK8sClusterService.ClusterID(c), k8sDeployments); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rek8sDeployments, "获取成功", c)
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
// func (kubernetesApi *KubernetesApi) RestartK8sDeployment(c *gin.Context) {
// 	// 指定Cluster
// 	kubeConfig, _ := kubernetesApi.ClusterID(c)
// 	namespace := c.Query("namespace")
// 	deployment := c.Query("deployment")
// 	if err := kubernetesService.RestartK8sDeployment(kubeConfig, namespace, deployment); err != nil {
// 		global.GVA_LOG.Error("重启应用失败", zap.Any("err", err))
// 		response.FailWithMessage("重启应用失败", c)
// 	} else {
// 		response.OkWithMessage("重启应用成功", c)
// 	}
// }

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
func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) GetK8sDeploymentList(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sDeploymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	namespace := c.Query("namespace")
	if list, total, err := GvaK8sDeploymentService.GetList(GvaK8sClusterService.ClusterID(c), namespace); err != nil {
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

func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) GetConfigMap(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sDeploymentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	namespace := c.Query("namespace")
	if list, total, err := GvaK8sDeploymentService.GetConfigMap(GvaK8sClusterService.ClusterID(c), namespace); err != nil {
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

func (GvaK8sDeploymentApi *GvaK8sDeploymentApi) FindConfigMap(c *gin.Context) {
	var K8sConfigMap kubernetesresponse.K8sConfigMap
	_ = c.ShouldBindQuery(&K8sConfigMap)
	if rek8sDeployments, err := GvaK8sDeploymentService.FindGetConfigMap(GvaK8sClusterService.ClusterID(c), K8sConfigMap); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rek8sDeployments, "获取成功", c)
	}
}
