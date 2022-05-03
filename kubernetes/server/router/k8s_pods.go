package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	v1 "github.com/tscuite/gva-plugin/kubernetes/server/api"

	"github.com/gin-gonic/gin"
)

type K8sPodsRouter struct {
}

func (e *K8sPodsRouter) InitK8sPodsRouter(Router *gin.RouterGroup) {
	//K8sPodsRouter := Router.Group("k8sPods").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	K8sPodsRouter := Router.Use(middleware.OperationRecord())
	var kubernetesApi = v1.ApiGroupApp.KubernetesApi
	{
		K8sPodsRouter.POST("createK8sPods", kubernetesApi.CreateK8sPods)             // 新建K8sPods
		K8sPodsRouter.DELETE("deleteK8sPods", kubernetesApi.DeleteK8sPods)           // 删除K8sPods
		K8sPodsRouter.DELETE("deleteK8sPodsByIds", kubernetesApi.DeleteK8sPodsByIds) // 批量删除K8sPods
		K8sPodsRouter.PUT("updateK8sPods", kubernetesApi.UpdateK8sPods)              // 更新K8sPods
		K8sPodsRouter.GET("getPodLog", kubernetesApi.GetPodLog)                      // 获取Pod日志
		K8sPodsRouter.GET("getK8sPodsList", kubernetesApi.GetK8sPodsList)            // 获取K8sPods列表
	}
}
func (e *K8sPodsRouter) InitK8sContainer(Router *gin.RouterGroup) {
	//containerRouter := Router.Group("container")
	containerRouter := Router.Use(middleware.OperationRecord())
	var kubernetesApi = v1.ApiGroupApp.KubernetesApi
	{
		containerRouter.GET("execContainer", kubernetesApi.ExecContainer) // 进入container终端
	}
}
