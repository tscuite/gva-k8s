package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type K8sDeployment struct {
	global.GVA_MODEL
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Replicas   int32  `json:"replicas"`
	CreateTime string `json:"createTime"`
}
