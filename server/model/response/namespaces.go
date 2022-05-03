package response

type K8sNamespaces struct {
	ID         uint   `json:"id"`
	Namespace  string `json:"namespace" gorm:"comment:命名空间"`
	Status     string `json:"status" gorm:"comment:状态"`
	CreateTime string `json:"createTime" gorm:"comment:创建时间"`
}
