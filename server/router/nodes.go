package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	v1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/api"

	"github.com/gin-gonic/gin"
)

type K8sNodesRouter struct {
}

func (e *K8sNodesRouter) InitK8sNodesRouter(Router *gin.RouterGroup) {
	//K8sNodesRouter := Router.Group("k8sNodes").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	K8sNodesRouter := Router.Use(middleware.OperationRecord())
	var kubernetesApi = v1.ApiGroupApp.GvaK8sNodeApi
	{
		K8sNodesRouter.GET("getK8sNodesList", kubernetesApi.GetK8sNodesList) // 获取K8sNodes列表
	}
}
