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
        <el-table-column align="left" label="deployment" min-width="150" prop="deployment" sortable="custom"></el-table-column> 
        <el-table-column align="left" label="replicas" min-width="150" prop="replicas" sortable="custom"></el-table-column> 
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
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle"  draggable destroy-on-close center >
      <warning-bar title="新增deployments.apps,需要手动写下方对应的容器名" /> 
      <el-form ref="apiForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="名字" prop="name">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="名称空间" prop="namespace">
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
        <el-form-item label="副本数" prop="replicas">
          <el-input v-model.number="form.replicas" autocomplete="off" />
        </el-form-item>
        <el-form-item label="port" >
         <div>
          <el-button size="small" type="primary" icon="edit" @click="addPort(form)" >添加port</el-button>
          <el-table :data="form.port" style="width: 100%">
            <el-table-column align="left" prop="name" label="name" width="120%">
              <template #default="scope">
                <div>
                  <el-input v-model="scope.row.name" />
                </div>
              </template>
            </el-table-column>
            <el-table-column align="left" prop="port" label="port" >
              <template #default="scope">
                <div>
                  <el-input v-model.number="scope.row.port" />
                </div>
              </template>
            </el-table-column>
            <el-table-column align="left" prop="images" label="需要手动写下方对应的容器名" width="120">
              <template #default="scope">
                <div>
                  <el-input v-model="scope.row.images" />
                </div>
              </template>
            </el-table-column>
            
            <el-table-column align="left" width="200" >
              <template #default="scope">
                <div>
                  <el-button type="danger" size="small" icon="delete" @click="remove(form.port,scope.$index)" >删除</el-button>
                </div>
              </template>
            </el-table-column>
            </el-table>
           </div>
          </el-form-item>
      </el-form>
      <div>
        <el-button size="small" type="primary" icon="edit" @click="addParameter(form)" >添加容器</el-button>
        <el-table :data="form.images" style="width: 100%">
          <el-table-column align="left" prop="name" label="设置" width="800">
            <template #default="props">
              <div>
                <el-form :data="form.images" label-width="125px">
                  <el-collapse label="left">
                  <el-collapse-item title="点击查看" name="1">
                    <el-form-item label="name" >
                      <el-input v-model.number="props.row.name"/>
                    </el-form-item>
                    <el-form-item label="images" >
                      <el-input v-model.number="props.row.images"/>
                    </el-form-item>
                    <el-form-item label="LimitsMemory(Mi)" >
                      <el-input v-model.number="props.row.LimitsMemory"/>
                    </el-form-item>
                    <el-form-item label="LimitsCPU(核)" >
                      <el-input v-model.number="props.row.LimitsCPU" />
                    </el-form-item>
                    <el-form-item label="RequestsCPU(核)" >
                      <el-input v-model.number="props.row.RequestsCPU" />
                    </el-form-item>
                    <el-form-item label="RequestsMemory(Mi)" >
                      <el-input v-model.number="props.row.RequestsMemory" />
                    </el-form-item>
                    <el-form-item label="健康检查端口" >
                      <el-input v-model.number="props.row.check" />
                    </el-form-item>
                  </el-collapse-item>
                  </el-collapse>
                </el-form>
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" width="200" >
            <template #default="scope">
              <div>
                <el-button type="danger" size="small" icon="delete" @click="remove(form.images,scope.$index)" >删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
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
  name: "K8sDeployment",
}
</script>

<script setup>
import {
    createK8sDeployment,
    getK8sDeploymentList,
    updateK8sDeployment,
    findK8sDeployment,
    deleteK8sDeployment,
} from "@/api/k8sDeployments";
import { getK8sNamespacesList } from "@/api/k8sNamespaces"; 
import { ref } from 'vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
import  {useRoute}  from "vue-router";
const route = useRoute()

const type = ref('')
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const nstableData = ref([])
const searchInfo = ref({})
const apiForm = ref(null)
const namespace = ref(route.params.namespace)

// 查询
const getTableData = async() => {
  const table = await getK8sDeploymentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value, namespace: namespace.value})
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


//增加按钮
const dialogTitle = ref('新增dep')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = '新增dep'
      break
    case 'edit':
      dialogTitle.value = '编辑dep'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}

//输入提示
const rules = ref({
  deployment: [{ required: true, message: '请输入deployment名字', trigger: 'blur' }],
  images: [
    { required: true, message: '请输入images', trigger: 'blur' }
  ],
  namespaces: [
    { required: true, message: '请选择namespaces', trigger: 'blur' }
  ],
  replicas: [
    { required: true, message: '请输入副本数', trigger: 'blur' }
  ]
})

const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    name: '',
    namespaces: '',
    port: [],
    images: [],
    replicas: Number(),
  }
}
const form = ref({
  name: '',
  namespaces: '',
  images: [],
  port: [],
  replicas: Number(),
})

//选项，后期改成获取到的名称空间
const methodOptions = ref([
  {
    value: 'default',
    label: 'default',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  },
  {
    value: 'PUT',
    label: '更新',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: '删除',
    type: 'danger'
  }
])

//取消
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
//确定
const enterDialog = async() => {
  apiForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addApi':
          {
            const res = await createK8sDeployment(form.value)
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
            const res = await updateK8sDeployment(form.value)
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
          // eslint-disable-next-line no-lone-blocks
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
  const res = await findK8sDeployment({ Name: row.deployment,Namespaces: row.namespaces})
  form.value = res.data
  openDialog('edit')
}

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteK8sDeployment({ Name: row.deployment,Namespaces: row.namespace})
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

// 新增参数
const addParameter = (form) => {
  if (!form.images) {
    form.value.images = []
  }
  form.images.push({
    name: '',
    images: '',
    check: Number(),
    LimitsCPU: Number(),
    LimitsMemory: Number(),
    RequestsCPU: Number(),
    RequestsMemory: Number(),
  })
}


// 新增参数2
const addPort = (form) => {
  if (!form.port) {
    form.value.port = []
  }
  form.port.push({
    name: '',
    images: '',
    port: Number(),
  })
}
// 删除参数1and2
const remove = (arr, index) => {
  arr.splice(index, 1)
}

// const authOptions = ref([])
// const setOptions = (authData) => {
//   authOptions.value = []
//   setAuthorityOptions(authData, authOptions.value)
// }

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