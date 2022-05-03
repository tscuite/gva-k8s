<template>
  <div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column type="selection" min-width="55%"></el-table-column>
        <el-table-column label="ID" prop="id" min-width="55%"></el-table-column>
        <el-table-column label="容器" prop="podName" min-width="270%"></el-table-column> 
        <el-table-column label="容器IP" prop="podIP" min-width="120%"></el-table-column>
        <el-table-column label="主机IP" prop="hostIP" min-width="150%"></el-table-column>
        <el-table-column label="状态" prop="status" min-width="100%"></el-table-column> 
        <el-table-column label="启动时间" prop="startTime" min-width="200%"></el-table-column>
        <el-table-column label="重启次数" prop="restartCount" min-width="80%"></el-table-column> 
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  name: "K8sDeployment",
}
</script>

<script  setup>
import {
    getK8sPodsList
} from "@/api/k8sPods";
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'


const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})


// 查询
const getTableData = async() => {
  const table = await getK8sPodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()


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