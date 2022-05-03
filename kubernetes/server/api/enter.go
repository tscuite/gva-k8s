package api

import "github.com/tscuite/gva-plugin/kubernetes/server/service"

type ApiGroup struct {
	KubernetesApi
	KubernetesclustersApi
}

var ApiGroupApp = new(ApiGroup)
var kubernetesService = service.ServiceGroupApp.KubernetesService
var kubernetesclustersService = service.ServiceGroupApp.KubernetesclustersService
