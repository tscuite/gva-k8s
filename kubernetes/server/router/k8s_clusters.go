package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/tscuite/gva-plugin/kubernetes/server/api"

	"github.com/gin-gonic/gin"
)

type K8sClusterRouter struct {
}

func (e *K8sClusterRouter) InitK8sClusterRouter(Router *gin.RouterGroup) {
	K8sClusterRouter := Router.Use(middleware.OperationRecord())
	//K8sClusterRouter := Router.Group("k8sCluster").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	var kubernetesclustersApi = api.ApiGroupApp.KubernetesclustersApi
	{
		K8sClusterRouter.POST("createK8sCluster", kubernetesclustersApi.CreateK8sCluster)             // 新建K8sCluster
		K8sClusterRouter.DELETE("deleteK8sCluster", kubernetesclustersApi.DeleteK8sCluster)           // 删除K8sCluster
		K8sClusterRouter.DELETE("deleteK8sClusterByIds", kubernetesclustersApi.DeleteK8sClusterByIds) // 批量删除K8sCluster
		K8sClusterRouter.PUT("updateK8sCluster", kubernetesclustersApi.UpdateK8sCluster)              // 更新K8sCluster
		K8sClusterRouter.GET("findK8sCluster", kubernetesclustersApi.FindK8sCluster)                  // 根据ID获取K8sCluster
		K8sClusterRouter.GET("getK8sClusterList", kubernetesclustersApi.GetK8sClusterList)            // 获取K8sCluster列表
	}
}
