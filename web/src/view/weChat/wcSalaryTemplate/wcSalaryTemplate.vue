<template>
	<div>
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button type="primary" icon="plus" @click="openDialog">新建模板</el-button>
			</div>
			<el-table
				ref="multipleTable"
				v-loading="loading"
				style="width: 100%"
				tooltip-effect="dark"
				:data="tableData"
				row-key="ID"
			>
				<el-table-column type="selection" width="55" />
				<el-table-column align="left" label="工资单模板" prop="name" />
				<el-table-column align="left" label="工资类型" prop="typeText" />
				<el-table-column align="left" label="适用职级体系" prop="rankTypeText" />
				<el-table-column align="left" label="操作">
					<template #default="{ row }">
						<el-button type="primary" link @click="handleUpdate(row)">编辑</el-button>
						<el-button type="primary" link @click="handleDel(row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>
			<div class="gva-pagination">
				<el-pagination
					layout="total, sizes, prev, pager, next, jumper"
					:current-page="searchInfo.page"
					:page-size="searchInfo.pageSize"
					:page-sizes="[10, 30, 50, 100]"
					:total="total"
					@current-change="handleCurrentChange"
					@size-change="handleSizeChange"
				/>
			</div>
		</div>

		<drawer ref="drawerRef" :modeType="modeType" :rowInfo="rowInfo" @onSuccess="getWcSalaryTemplate" />
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import { ElMessage, ElMessageBox } from 'element-plus'

import { getWcSalaryTemplateList, deleteWcSalaryTemplateByIds } from '@/api/weChat/wcSalaryTemplate'

import drawer from './components/drawer.vue'

const searchInfo = ref({
	page: 1,
	pageSize: 10,
	keyword: '',
})

const total = ref(0)

const loading = ref(false)
const rowInfo = ref({})
const modeType = ref('add')

const drawerRef = ref()
const openDialog = () => {
	modeType.value = 'add'
	drawerRef.value.dialogFormVisible = true
}

const tableData = ref([])
const getWcSalaryTemplate = () => {
	loading.value = true
	getWcSalaryTemplateList(searchInfo.value).then((res) => {
		loading.value = false
		tableData.value = res.data.list || []
		total.value = res.data.total
	})
}

const handleCurrentChange = (page) => {
	searchInfo.value.page = page
	getWcSalaryTemplate()
}

const handleSizeChange = (pageSize) => {
	searchInfo.value.pageSize = pageSize
	searchInfo.value.page = 1
	getWcSalaryTemplate()
}

const handleUpdate = (row) => {
	rowInfo.value = row
	modeType.value = 'edit'
	drawerRef.value.dialogFormVisible = true
}

const handleDel = (row) => {
	ElMessageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	}).then(() => {
		deleteWcSalaryTemplateByIds({ IDs: [row.ID] }).then(() => {
			ElMessage.success('删除成功')
			getWcSalaryTemplate()
		})
	})
}

onMounted(() => {
	getWcSalaryTemplate()
})
</script>
