package response

type K8sPods struct {
	ID           uint   `json:"id"`
	PodName      string `form:"podName" json:"podName"`
	PodIP        string `json:"podIP"`
	HostIP       string `json:"hostIP"`
	Status       string `json:"status"`
	StartTime    string `json:"startTime"`
	RestartCount int32  `json:"restartCount"`
	NameSpace    string `form:"namespace" json:"namespace"`
	Line         int64  `form:"line" json:"line"`
}
