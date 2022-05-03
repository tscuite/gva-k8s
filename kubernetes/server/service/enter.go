package service

type ServiceGroup struct {
	KubernetesService
	KubernetesclustersService
}

var ServiceGroupApp = new(ServiceGroup)
