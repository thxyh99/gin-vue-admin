<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">添 加</span>
				<div>
					<el-button type="primary">确 定</el-button>
					<el-button>取 消</el-button>
				</div>
			</div>
		</template>
		<el-form label-width="120" label-position="right">
			<el-form-item label="工资单模板名称">
				<el-input />
			</el-form-item>
			<el-form-item label="工资类型">
				<el-select></el-select>
			</el-form-item>
			<el-form-item label="适用职级体系">
				<el-select></el-select>
			</el-form-item>
		</el-form>
		<el-table :data="tableData" style="width: 100%">
			<el-table-column align="left" label="工资单字段" prop="label" />
			<el-table-column align="left" label="员工可见">
				<template #default="{ row }">
					<el-switch></el-switch>
				</template>
			</el-table-column>
		</el-table>
	</el-drawer>
</template>

<script setup>
import { ref } from 'vue'

const dialogFormVisible = ref(false)

const tableData = ref([
	{ label: '等级工资', value: 1 },
	{ label: '职称技能工资', value: 1 },
	{ label: '保密工资', value: 1 },
	{ label: 'CPI工资', value: 1 },
	{ label: '津贴', value: 1 },
	{ label: '司龄津贴', value: 1 },
	{ label: '学历工资', value: 1 },
	{ label: '全勤', value: 1 },
	{ label: '尽孝基金', value: 1 },
	{ label: '节日金', value: 1 },
])

defineExpose({
	dialogFormVisible,
})
</script>
