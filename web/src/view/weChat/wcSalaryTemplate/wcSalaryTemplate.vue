<template>
	<div>
		<!-- <div class="gva-search-box">
			<el-form
				ref="elSearchFormRef"
				:inline="true"
				:model="searchInfo"
				class="demo-form-inline"
				@keyup.enter="onSubmit"
			>
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
		</div> -->
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button type="primary" icon="plus" @click="openDialog">新建模板</el-button>
			</div>
			<el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID">
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
		</div>

		<drawer ref="drawerRef" />
	</div>
</template>

<script setup>
import { ref } from 'vue'
import drawer from './components/drawer.vue'

const searchInfo = ref({})

const onSubmit = () => {}

const drawerRef = ref()
const openDialog = () => {
	drawerRef.value.dialogFormVisible = true
}
const tableData = ref([])
</script>
