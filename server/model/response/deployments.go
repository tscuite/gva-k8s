package response

type K8sDeployment struct {
	ID         uint   `json:"id"`
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Replicas   int32  `json:"replicas"`
	CreateTime string `json:"createTime"`
}

type K8sConfigMap struct {
	ID         uint              `json:"id"`
	Namespaces string            `json:"namespaces"`
	ConfigMap  string            `json:"configmap"`
	Config     map[string]string `json:"config"`
	CreateTime string            `json:"createTime"`
}

type DeploymentSpec struct {
	Name       string    `json:"name,omitempty"`
	Port       []*Port   `json:"port"`
	Images     []*Images `json:"images"`
	Replicas   *int32    `json:"replicas" protobuf:"varint,3,opt,name=replicas"`
	Namespaces string    `json:"namespaces,omitempty"`
}

type DeploymentSpecres struct {
	Name       string      `json:"name,omitempty"`
	Port       []*Port     `json:"port"`
	Images     []Imagesres `json:"images"`
	Replicas   *int32      `json:"replicas" protobuf:"varint,3,opt,name=replicas"`
	Namespaces string      `json:"namespaces,omitempty"`
}

type Images struct {
	Name           string `json:"name"`
	Images         string `json:"images"`
	Check          int    `json:"check"`
	LimitsCPU      int    `json:"LimitsCPU"`
	LimitsMemory   int    `json:"LimitsMemory"`
	RequestsCPU    int    `json:"RequestsCPU"`
	RequestsMemory int    `json:"RequestsMemory"`
}

type Port struct {
	Name   string `json:"name"`
	Images string `json:"images"`
	Port   int32  `json:"port"`
}

type Imagesres struct {
	Name           string  `json:"name"`
	Images         string  `json:"images"`
	Check          int     `json:"check"`
	LimitsCPU      int     `json:"LimitsCPU"`
	LimitsMemory   int     `json:"LimitsMemory"`
	RequestsCPU    int     `json:"RequestsCPU"`
	RequestsMemory int     `json:"RequestsMemory"`
	Port           []*Port `json:"port"`
}
