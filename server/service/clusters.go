package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	k8sRequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/request"
	k8sResponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/gva-k8s/server/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GvaK8sClusterService struct {
}

func (GvaK8sClusterService *GvaK8sClusterService) ClusterID(c *gin.Context) string {
	cluster, err := GvaK8sClusterService.BoolK8sCluster()
	if err != nil {
		global.GVA_LOG.Error("获取集群失败", zap.Any("err", err))
		response.FailWithMessage("获取集群失败", c)
	}
	return cluster.KubeConfig
}

func (GvaK8sClusterService *GvaK8sClusterService) Create(K8sCluster k8sResponse.K8sCluster) (err error) {
	err = global.GVA_DB.Create(&K8sCluster).Error
	return err
}

func (GvaK8sClusterService *GvaK8sClusterService) Delete(K8sCluster *k8sResponse.K8sCluster) (err error) {
	err = global.GVA_DB.Delete(K8sCluster).Error
	return err
}

func (GvaK8sClusterService *GvaK8sClusterService) DeleteId(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]k8sResponse.K8sCluster{}, "id in ?", ids.Ids).Error
	return err
}

func (GvaK8sClusterService *GvaK8sClusterService) Update(K8sCluster *k8sResponse.K8sCluster) (err error) {
	err = global.GVA_DB.Save(K8sCluster).Error
	return err
}

func (GvaK8sClusterService *GvaK8sClusterService) Get(id uint) (K8sCluster k8sResponse.K8sCluster, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&K8sCluster).Error
	return
}

func (GvaK8sClusterService *GvaK8sClusterService) BoolK8sCluster() (K8sCluster k8sResponse.K8sCluster, err error) {
	err = global.GVA_DB.Where("start = ?", 1).First(&K8sCluster).Error
	return
}

func (GvaK8sClusterService *GvaK8sClusterService) GetList(info k8sRequest.K8sClusterSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&k8sResponse.K8sCluster{})
	var K8sClusters []k8sResponse.K8sCluster
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&K8sClusters).Error
	return K8sClusters, total, err
}
