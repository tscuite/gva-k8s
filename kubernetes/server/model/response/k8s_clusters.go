package response

type K8sCluster struct {
	ID             uint   `json:"ID" gorm:"primarykey;AUTO_INCREMENT" form:"ID"`
	ClusterName    string `json:"clusterName" gorm:"comment:集群名称"`
	KubeConfig     string `json:"kubeConfig" gorm:"comment:Kubeconfig内容;type:varchar(12800)"`
	ClusterVersion string `json:"clusterVersion" gorm:"comment:集群版本"`
}
