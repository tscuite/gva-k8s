package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type K8sCluster struct {
	global.GVA_MODEL
	ClusterName    string `json:"clusterName" gorm:"comment:集群名称"`
	KubeConfig     string `json:"kubeConfig" gorm:"comment:Kubeconfig内容;type:varchar(12800)"`
	ClusterVersion string `json:"clusterVersion" gorm:"comment:集群版本"`
	Start          bool   `json:"start" gorm:"comment:启用"`
}
