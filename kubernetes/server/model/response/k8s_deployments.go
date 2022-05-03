package response

type K8sDeployment struct {
	ID         uint   `json:"id"`
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Replicas   int32  `json:"replicas"`
	CreateTime string `json:"createTime"`
}
