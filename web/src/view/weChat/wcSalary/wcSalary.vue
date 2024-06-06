<template>
	<div>
		<div class="gva-search-box">
			<el-form ref="elFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="getWcSalary">
				<el-form-item label="工资月份">
					<el-date-picker
						style="width: 220px"
						type="monthrange"
						range-separator="至"
						start-placeholder="开始时间"
						end-placeholder="结束时间"
						v-model="dateRange"
						value-format="YYYYMM"
						@change="handleDataChange"
					/>
				</el-form-item>
				<el-form-item label="发放模板">
					<el-select style="width: 220px" v-model="searchInfo.templateId">
						<el-option v-for="item in templateList" :key="item.ID" :label="item.name" :value="item.ID" />
					</el-select>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" icon="search" @click="getWcSalary">查询</el-button>
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
				<el-table-column align="left" label="工资月份" prop="month" />
				<el-table-column align="left" label="工资类型" prop="typeText" />
				<el-table-column align="left" label="发放职级" prop="rankTypeText" />
				<el-table-column align="left" label="发放模板" prop="name" />
				<el-table-column align="left" label="操作">
					<template #default="{ row }">
						<el-button type="primary" link>
              <el-button type="primary" link>
                <ImportExcel template-id="staffSalary234" type="3" month="202402" rankType="1" @on-success="getTableData" btnName="导入" />
              </el-button>
            </el-button>
						<el-button type="primary" link>查看明细</el-button>
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
					:total="total"
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
import { getWcSalaryTemplateList } from '@/api/weChat/wcSalaryTemplate'

import { ElMessage, ElMessageBox } from 'element-plus'
import ImportExcel from "@/components/exportExcel/wechat/importSalaryExcel.vue";

const searchInfo = ref({
	page: 1,
	pageSize: 10,
	monthStart: '',
	monthEnd: '',
	templateId: undefined,
})

const loading = ref(false)
const dateRange = ref([])
const templateList = ref([])

const onReset = () => {
	searchInfo.value = {
		page: 1,
		pageSize: 10,
		monthStart: '',
		monthEnd: '',
		templateId: '',
	}
	dateRange.value = []
	getWcSalary()
}

const drawerRef = ref()
const modeType = ref('add')
const openDialog = () => {
	modeType.value = 'add'
	drawerRef.value.dialogFormVisible = true
}

const tableData = ref([])
const total = ref(0)
const getWcSalary = () => {
	loading.value = true
	getWcSalaryList(searchInfo.value).then((res) => {
		tableData.value = res.data.list || []
		total.value = res.data.total
		loading.value = false
	})
}

const handleDataChange = (e) => {
	if (e && e.length) {
		searchInfo.value.monthStart = e[0]
		searchInfo.value.monthEnd = e[1]
	} else {
		searchInfo.value.monthStart = ''
		searchInfo.value.monthEnd = ''
	}
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
	getWcSalaryTemplateList().then((res) => {
		templateList.value = res.data.list
	})
})
</script>
