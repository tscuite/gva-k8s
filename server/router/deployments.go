package router

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/api"

	"github.com/flipped-aurora/gin-vue-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type K8sDeploymentRouter struct {
}

func (e *K8sDeploymentRouter) InitK8sDeploymentRouter(Router *gin.RouterGroup) {
	//K8sDeploymentRouter := Router.Group("k8sDeployments").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	K8sDeploymentRouter := Router.Use(middleware.OperationRecord())
	var kubernetesApi = v1.ApiGroupApp.GvaK8sDeploymentApi
	{
		K8sDeploymentRouter.POST("createK8sDeployment", kubernetesApi.CreateK8sDeployment)   // 新建K8sDeployment
		K8sDeploymentRouter.DELETE("deleteK8sDeployment", kubernetesApi.DeleteK8sDeployment) // 删除K8sDeployment
		//K8sDeploymentRouter.DELETE("deleteK8sDeploymentByIds", kubernetesApi.DeleteK8sDeploymentByIds) // 批量删除K8sDeployment
		K8sDeploymentRouter.PUT("updateK8sDeployment", kubernetesApi.UpdateK8sDeployment)   // 更新K8sDeployment
		K8sDeploymentRouter.GET("findK8sDeployment", kubernetesApi.FindK8sDeployment)       // 根据ID获取K8sDeployment
		K8sDeploymentRouter.GET("getK8sDeploymentList", kubernetesApi.GetK8sDeploymentList) // 获取K8sDeployment列表
		K8sDeploymentRouter.GET("getConfigMap", kubernetesApi.GetConfigMap)
		K8sDeploymentRouter.GET("findConfigMap", kubernetesApi.FindConfigMap)

		//K8sDeploymentRouter.PUT("restartK8sDeployment", kubernetesApi.RestartK8sDeployment)            // 重启K8sDeployment
	}
}
