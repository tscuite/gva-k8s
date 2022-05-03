<template>
  <div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" min-width="55%"></el-table-column>   
    <el-table-column label="ID" prop="id" min-width="55%"></el-table-column>  
    <el-table-column label="容器" prop="podName" min-width="270%"></el-table-column> 
    <el-table-column label="容器IP" prop="podIP" min-width="120%"></el-table-column>
    <el-table-column label="主机IP" prop="hostIP" min-width="150%"></el-table-column>
    <el-table-column label="状态" prop="status" min-width="100%"></el-table-column> 
    <el-table-column label="启动时间" prop="startTime" min-width="200%"></el-table-column>   
    <el-table-column label="重启次数" prop="restartCount" min-width="80%"></el-table-column> 
    
      <el-table-column label="操作" min-width="250%">
        <template #default="scope">
          <el-button class="table-button" @click="findK8sPods(scope.row)" size="small" type="primary" icon="el-icon-edit">查看日志</el-button>
          <el-button class="table-button" @click="execContainer(scope.row)" size="small" type="primary" icon="el-icon-s-promotion">进入容器</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 查看日志 -->
    <el-dialog
      width="30%"
      title="查看日志"
      :visible.sync="innerVisible"
      append-to-body
    >
      {{ messageDialog }}
    </el-dialog>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import {
    createK8sPods,
    deleteK8sPods,
    deleteK8sPodsByIds,
    updateK8sPods,
    findK8sPods,
    execContainer,
    getK8sPodsList
} from "@/api/k8sPods";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";

export default {
  name: "K8sPods",
  mixins: [infoList],
  data() {
    return {
      listApi: getK8sPodsList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
            id:0,
            pod:"",      
      },
      messageDialog: "暂无日志返回",
      innerVisible: false
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10         
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteK8sPodsByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async findK8sPods(row) {
      // const res = await findK8sPods({ ID: row.ID });
      const { namespace, clusterID } = this.selectItem;
      const { podName } = row;
      const res = await findK8sPods({ namespace, podName, clusterID, line: row.id })
      this.type = "find";
      if (res.code == 0) {
        // 暂时自己定义，和下面的dialogFormVisible 具体意义是否一直暂时为避免错误，单独加了个flag
        this.messageDialog = res.data;
        this.innerVisible = true;
        this.formData = res.data.rek8sPods;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          id:0,
          pod:"",
          
      };
    },
    async deleteK8sPods(row) {
      this.visible = false;
      const res = await deleteK8sPods({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async execContainer(row) {
      this.visible = false;
      const res = await execContainer({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "进入容器成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createK8sPods(this.formData);
          break;
        case "update":
          res = await updateK8sPods(this.formData);
          break;
        case "find":
          res = await findK8sPods(this.formData);
          break;
        default:
          res = await createK8sPods(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  },
  computed: {
    ...mapGetters("user", ["selectItem"]),
  }
};
</script>

<style>
</style>