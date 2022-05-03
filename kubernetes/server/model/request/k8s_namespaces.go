package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	kubernetes "github.com/tscuite/gva-plugin/kubernetes/server/model/response"
)

type K8sNamespacesSearch struct {
	kubernetes.K8sDeployment
	request.PageInfo
}
