<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item>
                    <el-button @click="openDialog" type="primary">注册集群</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-table
            :data="tableData"
            @selection-change="handleSelectionChange"
            border
            ref="multipleTable"
            stripe
            style="width: 100%"
            tooltip-effect="dark"
            :row-style="{ height: 40 + 'px' }"
            :cell-style="{ padding: 0 + 'px' }"
        >
            <el-table-column type="selection" min-width="20%"></el-table-column>
            <el-table-column label="ID" prop="ID" min-width="20%"></el-table-column>
            <el-table-column label="集群名称" prop="clusterName" min-width="50%" ></el-table-column>
            <el-table-column label="Kubeconfig内容" prop="kubeConfig" min-width="100%" :show-overflow-tooltip="true" ></el-table-column>
            <el-table-column label="集群版本" prop="clusterVersion" min-width="20%"></el-table-column>
            <el-table-column label="操作">
                <template #default="scope">
                    <el-button  
                     class="table-button" 
                     @click="updateK8sCluster(scope.row)" 
                     size="small" 
                     type="primary" 
                     icon="el-icon-edit"
                     >编辑</el-button
                    >
                    <el-button  
                     class="table-button" 
                     @click="deleteK8sCluster(scope.row)" 
                     size="small" 
                     type="primary" 
                     icon="el-icon-edit"
                     >删除</el-button
                    >
                </template>
            </el-table-column>
        </el-table>
        <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :style="{ float: 'right', padding: '20px' }"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            layout="total, sizes, prev, pager, next, jumper"
        ></el-pagination>
        <el-dialog v-model="dialogFormVisible" modal:before-close="closeDialog" width="60%">
            <el-form  :model="formData"  label-position="right" label-width="80px" :rules="rules" ref="form">
                <el-form-item label="集群名称:" label-width="120px"  prop="clusterName">
                    <el-input v-model="formData.clusterName" clearable placeholder="请输入集群名称"></el-input>
                </el-form-item>
                <el-form-item  label-width="120px" label="kubeConfig:"  prop="KubeConfig">
                    <el-input  v-model="formData.KubeConfig"  type="textarea" placeholder="请输入KubeConfig文件内容" :autosize="{ minRows: 4, maxRows: 4 }" :style="{ width: '100%' }"></el-input>
                </el-form-item>
                <el-form-item  label-width="120px"  label="k8s版本:"  prop="clusterVersion">
                    <el-input v-model="formData.clusterVersion" placeholder="请输入k8s版本" :style="{ width: '100%' }"></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" default="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="sureHandle" type="primary">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
import {
    createK8sCluster,
    deleteK8sCluster,
    deleteK8sClusterByIds,
    updateK8sCluster,
    findK8sCluster,
    getK8sClusterList,
} from "@/api/k8sCluster";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
    name: "K8sCluster",
    mixins: [infoList],
    data() {
        return {
            listApi: getK8sClusterList,
            dialogFormVisible: false,
            visible: false,
            type: "",
            deleteVisible: false,
            multipleSelection: [],
            formData: {
                id: 0,
                clusterName: "",
                KubeConfig: "",
                clusterVersion: "",
                k8sVersio: "k8sVersion",
            },
            rules: {
                clusterName: [{ required: true, message: "请输入集群名称", trigger: "blur",},],
                KubeConfig: [{ required: true, message: "请输入KubeConfig文件内容", trigger: "blur",},],
                k8sVersion: [{ required: true, message: "请输入k8sVersion版本", trigger: "blur",},],
            },
        };
    },
    async created() {
        await this.getTableData();
    },
    methods: {
        //条件搜索前端看此方法
        onSubmit() {
            this.page = 1;
            this.pageSize = 10;
            this.getTableData();
        },
        handleSelectionChange(val) {
            this.multipleSelection = val;
        },
        async onDelete() {
            const ids = [];
            if (this.multipleSelection.length == 0) {
                this.$message({
                    type: "warning",
                    message: "请选择要删除的数据",
                });
                return;
            }
            this.multipleSelection &&
                this.multipleSelection.map((item) => {
                    ids.push(item.ID);
                });
            const res = await deleteK8sClusterByIds({ ids });
            if (res.code == 0) {
                this.$message({
                    type: "success",
                    message: "删除成功",
                });
                this.deleteVisible = false;
                this.getTableData();
            }
        },
        async updateK8sCluster(row) {
            const res = await findK8sCluster({  ID: row.ID });
            this.type = "update";
            if (res.code == 0) {
                this.formData = res.data.reK8sCluster;
                this.dialogFormVisible = true;
            }
        },
        closeDialog() {
            this.dialogFormVisible = false;
            this.formData = {
                id: 0,
                clusterName: "",
                kubeConfig: "",
                clusterVersion: "",
            };
        },
        deleteK8sCluster(row) {
          this.$confirm('此操作将永久删除所有角色下该账号, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          })
            .then(async() => {
              const res = await deleteK8sCluster({ ID: row.ID })
              if (res.code === 0) {
                this.$message({
                  type: 'success',
                  message: '删除成功!'
                })
                if (this.tableData.length === 1 && this.page > 1) {
                  this.page--
                }
                this.getTableData()
              }
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '已取消删除'
              })
            })
        },
        sureHandle() {
            this.$refs.form.validate((valid) => {
                if (valid) {
                    this.enterDialog();
                }
            });
        },
        async enterDialog() {
            let res;
            switch (this.type) {
                case "create":
                    res = await createK8sCluster(this.formData);
                    break;
                case "update":
                    res = await updateK8sCluster(this.formData);
                    break;
                default:
                    res = await createK8sCluster(this.formData);
                    break;
            }
            if (res.code == 0) {
                this.$message({
                    type: "success",
                    message: "创建/更改成功",
                });
                this.closeDialog();
                this.getTableData();
            }
        },
        openDialog() {
            this.type = "create";
            this.dialogFormVisible = true;
        },
    },
    filters: {
        formatDate: function (time) {
            if (time != null && time != "") {
                var date = new Date(time);
                return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
            } else {
                return "";
            }
        },
        formatBoolean: function (bool) {
            if (bool != null) {
                return bool ? "是" : "否";
            } else {
                return "";
            }
        },
    },
};
</script>

<style>
</style>