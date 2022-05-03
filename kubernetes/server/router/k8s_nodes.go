package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	v1 "github.com/tscuite/gva-plugin/kubernetes/server/api"

	"github.com/gin-gonic/gin"
)

type K8sNodesRouter struct {
}

func (e *K8sNodesRouter) InitK8sNodesRouter(Router *gin.RouterGroup) {
	//K8sNodesRouter := Router.Group("k8sNodes").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	K8sNodesRouter := Router.Use(middleware.OperationRecord())
	var kubernetesApi = v1.ApiGroupApp.KubernetesApi
	{
		K8sNodesRouter.POST("createK8sNodes", kubernetesApi.CreateK8sNodes)             // 新建K8sNodes
		K8sNodesRouter.DELETE("deleteK8sNodes", kubernetesApi.DeleteK8sNodes)           // 删除K8sNodes
		K8sNodesRouter.DELETE("deleteK8sNodesByIds", kubernetesApi.DeleteK8sNodesByIds) // 批量删除K8sNodes
		K8sNodesRouter.PUT("updateK8sNodes", kubernetesApi.UpdateK8sNodes)              // 更新K8sNodes
		K8sNodesRouter.GET("findK8sNodes", kubernetesApi.FindK8sNodes)                  // 根据ID获取K8sNodes
		K8sNodesRouter.GET("getK8sNodesList", kubernetesApi.GetK8sNodesList)            // 获取K8sNodes列表
	}
}
