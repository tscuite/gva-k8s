<template>
    <div class="gva-table-box">
      <div class="search-term">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item>
            <div v-auth="btnAuth.a" >
              <el-button size="small" type="primary" icon="plus" @click="openDialog('addApi')">注册集群</el-button>
            </div>
          </el-form-item>
          <el-form-item>
          </el-form-item>
        </el-form>
      </div>
      <el-table :data="tableData">
        <el-table-column label="id" prop="ID" min-width="20%"></el-table-column>
        <el-table-column label="集群名称" prop="clusterName" min-width="50%" ></el-table-column>
        <el-table-column label="Kubeconfig内容" prop="kubeConfig" min-width="100%" :show-overflow-tooltip="true" ></el-table-column>
        <el-table-column label="集群版本" prop="clusterVersion" min-width="20%"></el-table-column>
        <el-table-column min-width="26%" label="只能选一个( ture )">
          <template #default="scope">
          <el-tag :type="scope.row.start === false ? '' : 'success'" disable-transitions >{{ scope.row.start }}</el-tag>
         </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="text"
              @click="zupdateK8sCluster(scope.row)"
            >编辑</el-button>
            <el-button
              icon="delete"
              size="small"
              type="text"
              @click="cdeleteK8sCluster(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog v-model="dialogVisible" :before-close="closeDialog" :title="dialogTitle" draggable destroy-on-close center >
      <!-- <el-dialog v-model="dialogVisible" :before-close="closeDialog" :title="dialogTitle" > -->
        <el-form ref="apiForm" :model="form" :rules="rules" label-width="80px"> 
          <el-form-item label-width="120px" label="集群名称:"  prop="clusterName">
              <el-input v-model="form.clusterName" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item  label-width="120px" label="kubeConfig:"  prop="kubeConfig">
              <el-input  v-model="form.kubeConfig" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item label-width="120px" label="k8s版本:" prop="clusterVersion">
              <el-input v-model="form.clusterVersion" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label-width="120px" label="启用:" prop="start">
          <div>
            <el-checkbox v-model="form.start"> </el-checkbox>
          </div>
          </el-form-item>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button size="small" @click="closeDialog">取 消</el-button>
            <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
</template>

<script>
export default {
    name: "K8sCluster",
}
</script>

<script setup>
import {
    createK8sCluster,
    deleteK8sCluster,
    updateK8sCluster,
    findK8sCluster,
    getK8sClusterList,
} from "@/api/k8sCluster";
import warningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useBtnAuth } from '@/utils/btnAuth'
const btnAuth = useBtnAuth()
const type = ref('')
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 查询
const getTableData = async() => {
  const table = await getK8sClusterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const apiForm = ref(null)
const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    clusterName: "",
    kubeConfig: "",
    clusterVersion: "",
    start: false,
  }
}

const form = ref({
    clusterName: "",
    kubeConfig: "",
    clusterVersion: "",
    start: false,
})
          
const rules = ref({
    clusterName: [{ required: true, message: "请输入集群名称", trigger: "blur",},],
    kubeConfig: [{ required: true, message: "请输入KubeConfig文件内容", trigger: "blur",},],
    k8sVersion: [{ required: true, message: "请输入k8sVersion版本", trigger: "blur",},],
})
 

//增加按钮
const dialogTitle = ref('新增clusters')
const dialogVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = '新增clusters'
      break
    case 'edit':
      dialogTitle.value = '编辑clusters'
      break
    default:
      break
  }
  type.value = key
  dialogVisible.value = true
}

//取消
const closeDialog = () => {
  initForm()
  dialogVisible.value = false
}
//确定
const enterDialog = async() => {
  apiForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addApi':
          {
            const res = await createK8sCluster(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }
          break
        case 'edit':
          {
            const res = await updateK8sCluster(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '编辑成功',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }
          break
        default:
          {
            ElMessage({
              type: 'error',
              message: '未知操作',
              showClose: true
            })
          }
          break
      }
    }
  })
}

//更新
const zupdateK8sCluster = async(row) => {
  const res = await findK8sCluster({ ID: row.ID})
  form.value = res.data.K8sCluster
  openDialog('edit')
}

const cdeleteK8sCluster = async(row) => {
  // ElMessage({
  //             type: 'error',
  //             message: row,
  //             showClose: true
  //           })
  ElMessageBox.confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteK8sCluster({ ID: row.ID})
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
}

</script>

<style>
</style>