<template>
	<div class="gva-table-box">
		<div class="gva-btn-list">
			<el-button type="primary" icon="plus" @click="openDialog">新建模板</el-button>
		</div>
		<el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID">
			<el-table-column type="selection" width="55" />
			<el-table-column align="left" label="工资单模板" prop="" />
			<el-table-column align="left" label="工资类型" prop="" />
			<el-table-column align="left" label="适用职级" prop="" />
			<el-table-column align="left" label="操作">
				<template #default="scope">
					<el-button type="primary" link>编辑</el-button>
					<el-button type="primary" link>删除</el-button>
				</template>
			</el-table-column>
		</el-table>

		<drawer ref="drawerRef" />
	</div>
</template>

<script setup>
import { ref } from 'vue'
import drawer from './components/drawer.vue'

const tableData = ref([])

const drawerRef = ref()
const openDialog = () => {
	drawerRef.value.dialogFormVisible = true
}
</script>
