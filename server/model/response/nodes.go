package response

type K8sNodes struct {
	ID         uint   `json:"id"`
	NodeName   string `json:"nodeName" gorm:"comment:节点名称"`
	IP         string `json:"ip" gorm:"comment:ip地址"`
	Status     string `json:"status" gorm:"comment:节点状态"`
	Roles      string `json:"roles" gorm:"comment:节点角色"`
	CreateTime string `json:"createTime" gorm:"comment:创建时间"`
	Version    string `json:"version" gorm:"comment:节点版本"`
	Resource   string `json:"resource" gorm:"comment:节点资源"`
}
