package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/service"
)

type ApiGroup struct {
	GvaK8sClusterApi
	GvaK8sNamespaceApi
	GvaK8sDeploymentApi
	GvaK8sNodeApi
	GvaK8sPodApi
}

var ApiGroupApp = new(ApiGroup)
var GvaK8sClusterService = service.ServiceGroupApp.GvaK8sClusterService
var GvaK8sNamespaceService = service.ServiceGroupApp.GvaK8sNamespaceService
var GvaK8sDeploymentService = service.ServiceGroupApp.GvaK8sDeploymentService
var GvaK8sNodeService = service.ServiceGroupApp.GvaK8sNodeService
var GvaK8sPodService = service.ServiceGroupApp.GvaK8sPodService
