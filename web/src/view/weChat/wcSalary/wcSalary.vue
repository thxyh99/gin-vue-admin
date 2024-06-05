<template>
	<div>
		<div class="gva-search-box">
			<el-form ref="elFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
				<el-form-item label="发放职级">
					<el-select style="width: 220px" placeholder="请选择发放职级"></el-select>
				</el-form-item>
				<el-form-item label="工资月份">
					<el-date-picker
						style="width: 220px"
						type="monthrange"
						range-separator="至"
						start-placeholder="开始时间"
						end-placeholder="结束时间"
					/>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
					<el-button icon="refresh" @click="onReset">重置</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button type="primary" icon="plus" @click="openDialog">新建工资单</el-button>
			</div>
			<el-table
				ref="multipleTable"
				style="width: 100%"
				tooltip-effect="dark"
				:data="tableData"
				row-key="ID"
				v-loading="loading"
			>
				<el-table-column type="selection" width="55" />
				<el-table-column align="left" label="工资月份" prop="" />
				<el-table-column align="left" label="工资类型" prop="" />
				<el-table-column align="left" label="发放职级" prop="" />
				<el-table-column align="left" label="操作">
					<template #default="{ row }">
						<el-button type="primary" link>导入基本工资</el-button>
						<el-button type="primary" link>导入记录</el-button>
						<el-button type="primary" link @click="haneleDel(row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
			<div class="gva-pagination">
				<el-pagination
					layout="total, sizes, prev, pager, next, jumper"
					:current-page="searchInfo.page"
					:page-size="searchInfo.pageSize"
					:page-sizes="[10, 30, 50, 100]"
					:total="0"
					@current-change="handleCurrentChange"
					@size-change="handleSizeChange"
				/>
			</div>
		</div>

		<drawer ref="drawerRef" :modeType="modeType" @onSuccess="getWcSalary" />
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import drawer from './components/drawer.vue'

import { getWcSalaryList, deleteWcSalary } from '@/api/weChat/wcSalary'

import { ElMessage, ElMessageBox } from 'element-plus'

const searchInfo = ref({
	page: 1,
	pageSize: 10,
	monthStart: '',
	monthEnd: '',
	templateId: '',
})

const loading = ref(false)

const onSubmit = () => {}

const drawerRef = ref()
const modeType = ref('add')
const openDialog = () => {
	modeType.value = 'add'
	drawerRef.value.dialogFormVisible = true
}

const tableData = ref([])
const getWcSalary = () => {
	loading.value = true
	getWcSalaryList(searchInfo.value).then((res) => {
		tableData.value = res.data.list || []
		loading.value = false
	})
}

const handleCurrentChange = (page) => {
	searchInfo.value.page = page
	getWcSalary()
}

const handleSizeChange = (pageSize) => {
	searchInfo.value.pageSize = pageSize
	searchInfo.value.page = 1
	getWcSalary()
}

const haneleDel = (row) => {
	ElMessageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	}).then(() => {
		deleteWcSalary({ ID: row.ID }).then(() => {
			ElMessage.success('删除成功')
			getWcSalary()
		})
	})
}

onMounted(() => {
	getWcSalary()
})
</script>
