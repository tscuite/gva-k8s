package service

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/utils"

type ServiceGroup struct {
	GvaK8sClusterService
	GvaK8sNamespaceService
	GvaK8sDeploymentService
	GvaK8sNodeService
	GvaK8sPodService
}

var ServiceGroupApp = new(ServiceGroup)
var GvaK8sutils = utils.GroupApp.GvaK8sutils
