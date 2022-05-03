package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/tscuite/gva-plugin/kubernetes/server/router"
)

type k8sClusterPlug struct {
}
type k8sDeploymentsPlug struct {
}
type k8sNamespacesPlug struct {
}
type k8sNodesPlug struct {
}
type k8sPodsPlug struct {
}
type containerPlug struct {
}

func Createk8sClusterPlug() *k8sClusterPlug {
	return &k8sClusterPlug{}
}
func Createk8sDeploymentsPlug() *k8sDeploymentsPlug {
	return &k8sDeploymentsPlug{}
}
func Createk8sNamespacesPlug() *k8sNamespacesPlug {
	return &k8sNamespacesPlug{}
}
func Createk8sNodesPlug() *k8sNodesPlug {
	return &k8sNodesPlug{}
}
func Createk8sPodsPlug() *k8sPodsPlug {
	return &k8sPodsPlug{}
}

//func CreatecontainerPlug() *containerPlug {
//	return &containerPlug{}
//}

func (*k8sClusterPlug) RouterPath() string {
	return "k8sCluster"
}
func (*k8sDeploymentsPlug) RouterPath() string {
	return "k8sDeployments"
}
func (*k8sNamespacesPlug) RouterPath() string {
	return "k8sNamespaces"
}
func (*k8sNodesPlug) RouterPath() string {
	return "k8sNodes"
}
func (*k8sPodsPlug) RouterPath() string {
	return "k8sPods"
}

//func (*containerPlug) RouterPath() string {
//	return "container"
//}
func (*k8sClusterPlug) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitK8sClusterRouter(group)
}
func (*k8sDeploymentsPlug) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitK8sDeploymentRouter(group)
}
func (*k8sNamespacesPlug) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitK8sNamespacesRouter(group)
}
func (*k8sNodesPlug) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitK8sNodesRouter(group)
}
func (*k8sPodsPlug) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitK8sPodsRouter(group)
}
func (*containerPlug) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitK8sContainer(group)
}
