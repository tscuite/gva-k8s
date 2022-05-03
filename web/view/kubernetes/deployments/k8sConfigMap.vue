<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog('addApi')">创建deployment</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="selection" width="55"></el-table-column> 
        <el-table-column align="left" label="id" min-width="60" prop="id" sortable="custom"></el-table-column> 
        <el-table-column align="left" label="名称空间" min-width="150" prop="namespaces" sortable="custom"></el-table-column>
        <el-table-column align="left" label="configmap" min-width="150" prop="configmap" sortable="custom"></el-table-column> 
        <el-table-column align="left" label="config" min-width="150" prop="config" sortable="custom" >
           <template #default="scope">
            <el-button
              icon="Share"
              size="small"
              type="text"
              @click="editApiFunc(scope.row)"
            >查看</el-button>
          </template>
          </el-table-column> 
        <el-table-column align="left" label="创建时间" min-width="150" prop="createTime" sortable="custom"></el-table-column> 
        <el-table-column align="left" fixed="right" label="操作" width="200">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="text"
              @click="editApiFunc(scope.row)"
            >编辑</el-button>
            <el-button
              icon="delete"
              size="small"
              type="text"
              @click="deleteApiFunc(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>

    <el-dialog v-model="dialogVisible" :before-close="closeDialog" :title="dialogTitle" draggable destroy-on-close center >
      <!-- <el-dialog v-model="dialogVisible" :before-close="closeDialog" :title="dialogTitle" > -->
        <el-form ref="apiForm" :model="form" label-width="80px"> 
          <el-form-item label-width="120px" label="configmap:"  prop="configmap">
              <el-input v-model="form.configmap" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item label-width="120px" label="名称空间" prop="namespace">
            <el-select v-model="form.namespaces" class="m-2" placeholder="Select" size="large">
              <el-option
                v-for="item in nstableData"
                :key="item.namespace"
                :label="item.namespace"
                :value="item.namespace"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label-width="120px" label="config:" prop="config">
            <ul>
                <!-- key 为被遍历对象中的属性名， value 为对应的属性值-->
                <!-- <li v-for="(value, key, i) in form.config" v-bind:key="key" v-cloak> {{ key }}: {{ value }},{{ i }}</li> -->
                <li v-for="(value, key) in form.config" v-bind:key="key" v-cloak> {{ key }}: {{ value }}</li>
            </ul>
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
  </div>
</template>

<script>
export default {
  name: "K8sDeployment",
}
</script>

<script setup>
import {
    getConfigMap,
    findConfigMap,
} from "@/api/k8sConfigMap";
import { getK8sNamespacesList } from "@/api/k8sNamespaces"; 
import { ref } from 'vue'

import  {useRoute}  from "vue-router";
const route = useRoute()
const type = ref('')
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const nstableData = ref([])
const searchInfo = ref({})
const namespace = ref("default")

// 查询
const getTableData = async() => {
  const table = await getConfigMap({ page: page.value, pageSize: pageSize.value, ...searchInfo.value, namespace: namespace.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const nsgetTableData = async() => {
  const nstable = await getK8sNamespacesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (nstable.code === 0) {
    nstableData.value = nstable.data.list
    total.value = nstable.data.total
    page.value = nstable.data.page
    pageSize.value = nstable.data.pageSize
  }
}
nsgetTableData()

const apiForm = ref(null)
const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    configmap: "",
    namespaces: "",
    config: {},
  }
}

const form = ref({
    configmap: "",
    namespaces: "",
    config: {},
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
      dialogTitle.value = '编辑configmap'
      break
    default:
      break
  }
  type.value = key
  dialogVisible.value = true
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
const editApiFunc = async(row) => {
  const res = await findConfigMap({ ConfigMap: row.configmap,Namespaces: row.namespaces})
  form.value = res.data
  openDialog('edit')
}


//取消
const closeDialog = () => {
  initForm()
  dialogVisible.value = false
}

</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
  color: #dc143c;
}
</style>