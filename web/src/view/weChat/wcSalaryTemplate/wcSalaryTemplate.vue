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
				<el-table-column align="left" label="工资单模板" prop="" />
				<el-table-column align="left" label="工资类型" prop="" />
				<el-table-column align="left" label="适用职级体系" prop="" />
				<el-table-column align="left" label="操作">
					<template #default="scope">
						<el-button type="primary" link>编辑</el-button>
						<el-button type="primary" link>删除</el-button>
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

		<drawer ref="drawerRef" :modeType="modeType" :rowInfo="rowInfo" />
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import { getWcSalaryTemplateList } from '@/api/weChat/wcSalaryTemplate'

import drawer from './components/drawer.vue'

const searchInfo = ref({
	page: 1,
	pageSize: 10,
	keyword: '',
})

const loading = ref(false)
const rowInfo = ref({})
const modeType = ref('add')

const drawerRef = ref()
const openDialog = () => {
	drawerRef.value.dialogFormVisible = true
}

const tableData = ref([])
const getWcSalaryTemplate = () => {
	loading.value = true
	getWcSalaryTemplateList(searchInfo.value).then((res) => {
		loading.value = false
		tableData.value = res.data.list || []
	})
}

const handleCurrentChange = () => {
	getWcSalaryTemplate()
}

const handleSizeChange = () => {
	searchInfo.value.page = 1
	getWcSalaryTemplate()
}

onMounted(() => {
	getWcSalaryTemplate()
})
</script>
