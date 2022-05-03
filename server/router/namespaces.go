package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	v1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/api"
	"github.com/gin-gonic/gin"
)

type K8sNamespacesRouter struct {
}

func (e *K8sNamespacesRouter) InitK8sNamespacesRouter(Router *gin.RouterGroup) {
	//K8sNamespacesRouter := Router.Group("k8sNamespaces").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	K8sNamespacesRouter := Router.Use(middleware.OperationRecord())
	var kubernetesApi = v1.ApiGroupApp.GvaK8sNamespaceApi
	{
		K8sNamespacesRouter.POST("createK8sNamespaces", kubernetesApi.CreateK8sNamespaces) // 新建K8sNamespaces
		//K8sNamespacesRouter.DELETE("deleteK8sNamespaces", kubernetesApi.DeleteK8sNamespaces)           // 删除K8sNamespaces
		//K8sNamespacesRouter.DELETE("deleteK8sNamespacesByIds", kubernetesApi.DeleteK8sNamespacesByIds) // 批量删除K8sNamespaces
		//K8sNamespacesRouter.PUT("updateK8sNamespaces", kubernetesApi.UpdateK8sNamespaces)              // 更新K8sNamespaces
		//K8sNamespacesRouter.GET("findK8sNamespaces", kubernetesApi.FindK8sNamespaces)                  // 根据ID获取K8sNamespaces
		K8sNamespacesRouter.GET("getK8sNamespacesList", kubernetesApi.GetK8sNamespacesList) // 获取K8sNamespaces列表
	}
}
