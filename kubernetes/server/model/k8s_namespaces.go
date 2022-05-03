package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type K8sNamespaces struct {
	global.GVA_MODEL
	Namespace  string `json:"namespace" gorm:"comment:命名空间"`
	Status     string `json:"status" gorm:"comment:状态"`
	CreateTime string `json:"createTime" gorm:"comment:创建时间"`
}
