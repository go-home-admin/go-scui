<template>
	<el-container>
		<el-header>
			<div class="left-panel">
				<el-button icon="el-icon-Filter" @click="add">筛选</el-button>
			</div>
			<div class="right-panel">
				<div class="right-panel-search">
					<el-button type="primary" icon="el-icon-plus" @click="add">新增</el-button>
					<el-dropdown split-button type="primary">
						导出
						<template #dropdown>
							<el-dropdown-menu>
								<el-dropdown-item>全部</el-dropdown-item>
								<el-dropdown-item>当前页</el-dropdown-item>
								<el-dropdown-item>选中的行</el-dropdown-item>
							</el-dropdown-menu>
						</template>
					</el-dropdown>
					<el-button type="danger" plain icon="el-icon-delete" :disabled="selection.length==0" @click="batch_del"></el-button>
				</div>
			</div>
		</el-header>
		<el-main class="nopadding">
			<div style="margin: 15px 0px 0px 0px; border-bottom: 1px solid var(--el-border-color-light);">
				<el-row>
					<el-col :span="24">
						<el-form :model="form" label-width="120px">
							<el-form-item label="Activity name">
								<el-input v-model="form.name" />
							</el-form-item>
							<el-form-item label="Activity zone">
								<el-select v-model="form.region" placeholder="please select your zone">
									<el-option label="Zone one" value="shanghai" />
									<el-option label="Zone two" value="beijing" />
								</el-select>
							</el-form-item>
							<el-form-item>
								<el-button type="primary" @click="onSubmit">搜索</el-button>
								<el-button>重置</el-button>
							</el-form-item>
						</el-form>
					</el-col>
				</el-row>
			</div>
			<scTable ref="table" :apiObj="list.apiObj" row-key="id" @selection-change="selectionChange" stripe>
				<el-table-column type="selection" width="50"></el-table-column>
				<el-table-column label="姓名" prop="name" width="180"></el-table-column>
				<el-table-column label="性别" prop="sex" width="150"></el-table-column>
				<el-table-column label="邮箱" prop="email" width="250"></el-table-column>
				<el-table-column label="状态" prop="boolean" width="60">
					<template #default="scope">
						<sc-status-indicator v-if="scope.row.boolean" type="success"></sc-status-indicator>
						<sc-status-indicator v-if="!scope.row.boolean" pulse type="danger"></sc-status-indicator>
					</template>
				</el-table-column>
				<el-table-column label="评分" prop="num" width="150"></el-table-column>
				<el-table-column label="操作" fixed="right" align="right" width="300">
					<template #default="scope">
						<el-button plain size="small" @click="table_show(scope.row)">查看</el-button>
						<el-button type="primary" plain size="small" @click="table_edit(scope.row)">编辑</el-button>
						<el-popconfirm title="确定删除吗？" @confirm="table_del(scope.row, scope.$index)">
							<template #reference>
								<el-button plain type="danger" size="small">删除</el-button>
							</template>
						</el-popconfirm>
						<el-dropdown>
							<el-button icon="el-icon-more" size="small"></el-button>
							<template #dropdown>
								<el-dropdown-menu>
									<el-dropdown-item>配额</el-dropdown-item>
									<el-dropdown-item divided>重启</el-dropdown-item>
									<el-dropdown-item >停机</el-dropdown-item>
									<el-dropdown-item divided>释放主机</el-dropdown-item>
								</el-dropdown-menu>
							</template>
						</el-dropdown>
					</template>
				</el-table-column>
			</scTable>
		</el-main>
	</el-container>

	<save-dialog v-if="dialog.save" ref="saveDialog" @success="handleSaveSuccess" @closed="dialog.save=false"></save-dialog>

	<el-drawer v-model="dialog.info" :size="800" title="详细" custom-class="drawerBG" direction="rtl" destroy-on-close>
		<info ref="infoDialog"></info>
	</el-drawer>

</template>

<script>
	import saveDialog from './save'
	import info from './info'

	export default {
		name: 'listCrud',
		components: {
			saveDialog,
			info
		},
		data() {
			return {
				dialog:{
					save: false,
					info: false
				},
				list: {
					apiObj: this.$API.demo.list
				},
				selection: [],
				form: {
					name: '',
					region: '',
				}
			}
		},
		mounted() {

		},
		methods: {
			//点击筛选按钮
			onFilter() {

			},
			//窗口新增
			add(){
				this.dialog.save = true
				this.$nextTick(() => {
					this.$refs.saveDialog.open()
				})
			},
			//窗口编辑
			table_edit(row){
				this.dialog.save = true
				this.$nextTick(() => {
					this.$refs.saveDialog.open('edit').setData(row)
				})
			},
			//页面新增
			addPage(){
				this.$router.push({
					path: '/template/list/crud/detail',
				})
			},
			//查看
			table_show(row){
				this.dialog.info = true
				this.$nextTick(() => {
					this.$refs.infoDialog.setData(row)
				})
			},
			//删除明细
			async table_del(row, index){
				var reqData = {id: row.id}
				var res = await this.$API.demo.post.post(reqData);
				if(res.code == 200){
					this.$refs.table.removeIndex(index)
					this.$message.success("删除成功")
				}else{
					this.$alert(res.message, "提示", {type: 'error'})
				}
			},
			//批量删除
			async batch_del(){
				var confirmRes = await this.$confirm(`确定删除选中的 ${this.selection.length} 项吗？`, '提示', {
					type: 'warning',
					confirmButtonText: '删除',
					confirmButtonClass: 'el-button--danger'
				}).catch(() => {})

				if(!confirmRes){
					return false
				}

				var ids = this.selection.map(v => v.id)
				this.$refs.table.removeKeys(ids)
				this.$message.success("操作成功")

			},
			//表格选择后回调事件
			selectionChange(selection){
				this.selection = selection
			},
			//本地更新数据
			handleSaveSuccess(data, mode){
				//为了减少网络请求，直接变更表格内存数据
				if(mode=='add'){
					this.$refs.table.unshiftRow(data)
				}else if(mode=='edit'){
					this.$refs.table.updateKey(data)
				}

				//当然也可以暴力的直接刷新表格
				// this.$refs.table.refresh()
			}
		}
	}
</script>

<style>
</style>
