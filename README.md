# gva-plugin-kubernetes

#### 介绍
gin-vue-admin-plugin


复制粘贴了这个项目的代码https://github.com/openstack-test/gin-vue-devops    致谢这个项目的大佬


###server/initialize/plugin.go

	kubernetes "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server"

	PluginInit(PrivateGroup, kubernetes.Createk8sClusterPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sDeploymentsPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sNamespacesPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sNodesPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sPodsPlug())


###server/initialize/gorm.go

        kubernetes "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model"

        kubernetes.K8sCluster{},
  

###server/source/system/casbin.go

		{PType: "p", V0: "888", V1: "/k8sNamespaces/findK8sNamespaces", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sNamespaces/getK8sNamespacesList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sNamespaces/deleteK8sNamespaces", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/k8sDeployments/getK8sDeploymentList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sDeployments/createK8sDeployment", V2: "POST"},
		{PType: "p", V0: "888", V1: "/k8sDeployments/findK8sDeployment", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sDeployments/updateK8sDeployment", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/k8sDeployments/deleteK8sDeployment", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/k8sPods/getK8sPodsList", V2: "	GET"},
		{PType: "p", V0: "888", V1: "/k8sCluster/createK8sCluster", V2: "POST"},
		{PType: "p", V0: "888", V1: "/k8sCluster/getK8sClusterList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sCluster/findK8sCluster", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sCluster/deleteK8sCluster", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/k8sNodes/createK8sNodes", V2: "POST"},
		{PType: "p", V0: "888", V1: "/k8sNodes/deleteK8sNodes", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/k8sNodes/deleteK8sNodesByIds", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/k8sNodes/updateK8sNodes", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/k8sNodes/findK8sNodes", V2: "GET"},
		{PType: "p", V0: "888", V1: "/k8sNodes/getK8sNodesList", V2: "GET"},
  
  
###server/source/system/menu.go

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "kubernetes", Name: "kubernetes", Component: "view/kubernetes/index.vue", Sort: 2, Meta: Meta{Title: "Kubernetes管理", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "k8sCluster", Name: "k8sCluster", Component: "view/kubernetes/clusters/index.vue", Sort: 1, Meta: Meta{Title: "集群管理", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "k8sNamespaces", Name: "k8sNamespaces", Component: "view/kubernetes/namespaces/k8sNamespaces.vue", Sort: 3, Meta: Meta{Title: "命名空间管理", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "k8sDeployments", Name: "k8sDeployments", Component: "view/kubernetes/deployments/k8sDeployments.vue", Sort: 4, Meta: Meta{Title: "应用管理", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "k8sPods", Name: "k8sPods", Component: "view/kubernetes/pods/index.vue", Sort: 5, Meta: Meta{Title: "容器管理", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "k8sNodes", Name: "k8sNodes", Component: "view/kubernetes/nodes/k8sNodes.vue", Sort: 2, Meta: Meta{Title: "node管理", Icon: "menu"}},

###server/source/system/api.go

		{ApiGroup: "k8sNamespaces", Method: "GET", Path: "/k8sNamespaces/findK8sNamespaces", Description: "根据ID获取k8sNamespaces	"},
		{ApiGroup: "k8sNamespaces", Method: "GET", Path: "/k8sNamespaces/getK8sNamespacesList", Description: "获取所有k8sNamespaces"},
		{ApiGroup: "k8sNamespaces", Method: "DELETE", Path: "/k8sNamespaces/deleteK8sNamespaces", Description: "删除名称空间"},
		{ApiGroup: "k8sPods", Method: "GET", Path: "/k8sPods/getK8sPodsList", Description: "获取所有k8sPods"},
		{ApiGroup: "k8sCluster", Method: "POST", Path: "/k8sCluster/createK8sCluster", Description: "创建k8sCluster"},
		{ApiGroup: "k8sCluster", Method: "GET", Path: "/k8sCluster/getK8sClusterList", Description: "获取k8sCluster"},
		{ApiGroup: "k8sCluster", Method: "GET", Path: "/k8sCluster/findK8sCluster", Description: "用id查询k8sCluster"},
		{ApiGroup: "k8sCluster", Method: "DELETE", Path: "/k8sCluster/deleteK8sCluster", Description: "删除集群"},
		{ApiGroup: "k8sNodes", Method: "POST", Path: "/k8sNodes/createK8sNodes", Description: "	创建K8sNodes"},
		{ApiGroup: "k8sNodes", Method: "DELETE", Path: "/k8sNodes/deleteK8sNodes", Description: "删除K8sNodes"},
		{ApiGroup: "k8sNodes", Method: "DELETE", Path: "/k8sNodes/deleteK8sNodesByIds", Description: "批量删除K8sNodes"},
		{ApiGroup: "k8sNodes", Method: "PUT", Path: "/k8sNodes/updateK8sNodes", Description: "更新K8sNodes"},
		{ApiGroup: "k8sNodes", Method: "GET", Path: "/k8sNodes/findK8sNodes", Description: "用id查询K8sNodes"},
		{ApiGroup: "k8sNodes", Method: "GET", Path: "/k8sNodes/getK8sNodesList", Description: "分页获取K8sNodes列表"},
		{ApiGroup: "k8sDeployments", Method: "GET", Path: "/k8sDeployments/getK8sDeploymentList", Description: "获取所有k8sDeployments"},
		{ApiGroup: "k8sDeployments", Method: "POST", Path: "/k8sDeployments/createK8sDeployment", Description: "k8sDeployments"},
		{ApiGroup: "k8sDeployments", Method: "GET", Path: "/k8sDeployments/findK8sDeployment", Description: "findK8sDeployment"},
		{ApiGroup: "k8sDeployments", Method: "PUT", Path: "/k8sDeployments/updateK8sDeployment", Description: "updateK8sDeployment"},
		{ApiGroup: "k8sDeployments", Method: "DELETE", Path: "/k8sDeployments/deleteK8sDeployment", Description: "deleteK8sDeployment"},



  
