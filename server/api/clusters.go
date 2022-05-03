package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	kubernetesrequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/request"
	kubernetes "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GvaK8sClusterApi struct {
}

func (GvaK8sClusterApi *GvaK8sClusterApi) CreateK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := GvaK8sClusterService.Create(K8sCluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (GvaK8sClusterApi *GvaK8sClusterApi) DeleteK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := GvaK8sClusterService.Delete(&K8sCluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (GvaK8sClusterApi *GvaK8sClusterApi) DeleteK8sClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := GvaK8sClusterService.DeleteId(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (GvaK8sClusterApi *GvaK8sClusterApi) UpdateK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindJSON(&K8sCluster)
	if err := GvaK8sClusterService.Update(&K8sCluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags K8sCluster
// @Summary 用id查询K8sCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "用id查询K8sCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sCluster/findK8sCluster [get]
func (GvaK8sClusterApi *GvaK8sClusterApi) FindK8sCluster(c *gin.Context) {
	var K8sCluster kubernetes.K8sCluster
	_ = c.ShouldBindQuery(&K8sCluster)
	if reK8sCluster, err := GvaK8sClusterService.Get(K8sCluster.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"K8sCluster": reK8sCluster}, c)
	}
}

// @Tags K8sCluster
// @Summary 分页获取K8sCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sClusterSearch true "分页获取K8sCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sCluster/getK8sClusterList [get]
func (GvaK8sClusterApi *GvaK8sClusterApi) GetK8sClusterList(c *gin.Context) {
	var pageInfo kubernetesrequest.K8sClusterSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		fmt.Println("绑定查询失败: ", err.Error())
	}
	if list, total, err := GvaK8sClusterService.GetList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
