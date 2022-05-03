package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"
)

type K8sPodsSearch struct {
	response.K8sDeployment
	request.PageInfo
}
