<template>
  <div>
    <div class="search-term">
      <!-- <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">创建</el-button>
        </el-form-item>
      </el-form> -->
    </div>
    <el-table :data="tableData" style="width: 100%">
    <el-table-column type="selection" min-width="20%"></el-table-column>
    <el-table-column label="ID" prop="id" min-width="20%"></el-table-column>
    <el-table-column label="命名空间" prop="namespace" min-width="50%"></el-table-column> 
    <el-table-column label="运行状态" prop="status" min-width="30%"></el-table-column> 
    <el-table-column label="创建时间" prop="createTime" min-width="50%"></el-table-column>
    <el-table-column label="应用列表" prop="application" min-width="50%">
      <template #default="scope">
            <el-button
              size="small"
              icon="document"
              type="text"
              @click="toDetile(scope.row)"
            >详情</el-button>
      </template>
    </el-table-column> 
    </el-table>
  </div>
</template>

<script>
export default {
  name: "K8sNamespaces",
}
</script>
<script setup>
import {
    createK8sNamespaces,
    findK8sNamespaces,
    getK8sNamespacesList
} from "@/api/k8sNamespaces";
import { ref } from 'vue'
import  {useRouter}  from "vue-router";
const router = useRouter();

const type = ref('')
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const apiForm = ref(null)

const getTableData = async() => {
  const table = await getK8sNamespacesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const toDetile = (row) => {
  router.push({
    name: 'k8sDeployments',
    params: {
      namespace: row.namespace,
    },
  })
}


</script>

<style>
</style>