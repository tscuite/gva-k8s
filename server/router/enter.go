package router

type RouterGroup struct {
	K8sClusterRouter
	K8sDeploymentRouter
	K8sNamespacesRouter
	K8sNodesRouter
	K8sPodsRouter
}

var RouterGroupApp = new(RouterGroup)
