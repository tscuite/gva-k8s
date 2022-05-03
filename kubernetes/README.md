# gva-plugin-kubernetes

#### 介绍
gin-vue-admin-plugin


复制粘贴了这个项目的代码https://github.com/openstack-test/gin-vue-devops


###server/initialize/plugin.go

	kubernetes "github.com/tscuite/gva-plugin/plugin/kubernetes/server"

	PluginInit(PrivateGroup, kubernetes.Createk8sClusterPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sDeploymentsPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sNamespacesPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sNodesPlug())
	PluginInit(PrivateGroup, kubernetes.Createk8sPodsPlug())


###create table或加入初始化代码

    CREATE TABLE `k8s_clusters`  (
      `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
      `created_at` datetime(3) NULL DEFAULT NULL,
      `updated_at` datetime(3) NULL DEFAULT NULL,
      `deleted_at` datetime(3) NULL DEFAULT NULL,
      `cluster_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '集群名称',
      `kube_config` varchar(12800) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'Kubeconfig内容',
      `cluster_version` float NULL DEFAULT NULL COMMENT '集群版本',
      PRIMARY KEY (`id`) USING BTREE,
      INDEX `idx_k8s_clusters_deleted_at`(`deleted_at`) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;
  

###server/source/system/casbin.go

	{PType: "p", V0: "888", V1: "/k8sCluster/createK8sCluster", V2: "POST"},
	{PType: "p", V0: "888", V1: "/k8sCluster/deleteK8sCluster", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/k8sCluster/findK8sCluster", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sCluster/getK8sClusterList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sDeployments/getK8sDeploymentList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNamespaces/deleteK8sNamespaces", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/k8sNamespaces/findK8sNamespaces", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNamespaces/getK8sNamespacesList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNodes/createK8sNodes", V2: "POST"},
	{PType: "p", V0: "888", V1: "/k8sNodes/deleteK8sNodes", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/k8sNodes/deleteK8sNodesByIds", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/k8sNodes/findK8sNodes", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNodes/getK8sNodesList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNodes/updateK8sNodes", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/k8sPods/getK8sPodsList", V2: "GET"},
  
  
###server/source/system/menu.go

	{MenuLevel: 0, ParentId: "0", Path: "kubernetes", Name: "kubernetes", Hidden: false, Component: "view/kubernetes/index.vue", Sort: 2, Meta: system.Meta{Title: "Kubernetes管理", Icon: "cloudy"}},
	{MenuLevel: 0, ParentId: "26", Path: "k8sCluster", Name: "k8sCluster", Hidden: false, Component: "view/kubernetes/clusters/index.vue", Sort: 1, Meta: system.Meta{Title: "集群管理", Icon: "menu"}},
	{MenuLevel: 0, ParentId: "26", Path: "k8sNamespaces", Name: "k8sNamespaces", Hidden: false, Component: "view/kubernetes/namespaces/k8sNamespaces.vue", Sort: 3, Meta: system.Meta{Title: "命名空间管理", Icon: "menu"}},
	{MenuLevel: 0, ParentId: "26", Path: "k8sDeployments", Name: "k8sDeployments", Hidden: false, Component: "view/kubernetes/deployments/k8sDeployments.vue", Sort: 4, Meta: system.Meta{Title: "应用管理", Icon: "s-grid"}},
	{MenuLevel: 0, ParentId: "26", Path: "k8sPods", Name: "k8sPods", Hidden: false, Component: "view/kubernetes/pods/index.vue", Sort: 5, Meta: system.Meta{Title: "容器管理", Icon: "s-grid"}},
	{MenuLevel: 0, ParentId: "26", Path: "k8sNodes", Name: "k8sNodes", Hidden: false, Component: "view/kubernetes/nodes/k8sNodes.vue", Sort: 2, Meta: system.Meta{Title: "node管理", Icon: "menu"}},


###server/source/system/api.go

	{ApiGroup: "k8sNamespaces", Method: "GET", Path: "/k8sNamespaces/findK8sNamespaces", Description: "根据ID获取k8sNamespaces"},
	{ApiGroup: "k8sNamespaces", Method: "GET", Path: "/k8sNamespaces/getK8sNamespacesList", Description: "获取所有k8sNamespaces"},
	{ApiGroup: "k8sDeployments", Method: "GET", Path: "/k8sDeployments/getK8sDeploymentList", Description: "获取所有k8sDeployments"},
	{ApiGroup: "k8sPods", Method: "GET", Path: "/k8sPods/getK8sPodsList", Description: "获取所有k8sPods"},
	{ApiGroup: "k8sCluster", Method: "POST", Path: "/k8sCluster/createK8sCluster", Description: "创建k8sCluster"},
	{ApiGroup: "k8sCluster", Method: "GET", Path: "/k8sCluster/getK8sClusterList", Description: "获取k8sCluster"},
	{ApiGroup: "k8sCluster", Method: "GET", Path: "/k8sCluster/findK8sCluster", Description: "用id查询k8sCluster"},
	{ApiGroup: "k8sNodes", Method: "POST", Path: "/k8sNodes/createK8sNodes", Description: "创建K8sNodes"},
	{ApiGroup: "k8sNodes", Method: "DELETE", Path: "/k8sNodes/deleteK8sNodes", Description: "删除K8sNodes"},
	{ApiGroup: "k8sNodes", Method: "DELETE", Path: "/k8sNodes/deleteK8sNodesByIds", Description: "批量删除K8sNodes"},
	{ApiGroup: "k8sNodes", Method: "PUT", Path: "/k8sNodes/updateK8sNodes", Description: "更新K8sNodes"},
	{ApiGroup: "k8sNodes", Method: "GET", Path: "/k8sNodes/findK8sNodes", Description: "用id查询K8sNodes"},
	{ApiGroup: "k8sNodes", Method: "GET", Path: "/k8sNodes/getK8sNodesList", Description: "分页获取K8sNodes列表"},
	{ApiGroup: "k8sNamespaces", Method: "DELETE", Path: "/k8sNamespaces/deleteK8sNamespaces", Description: "删除名称空间"},
	{ApiGroup: "k8sCluster", Method: "DELETE", Path: "/k8sCluster/deleteK8sCluster", Description: "删除集群"},



  
