<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">添加</span>
				<div>
					<el-button type="primary" @click="enterDialog">确 定</el-button>
					<el-button @click="closeDialog">取 消</el-button>
				</div>
			</div>
		</template>
		<el-form :model="formData" label-width="120px" label-position="right" ref="elFormRef">
			<el-form-item label="工资单模板标题" prop="">
				<el-input style="width: 100%" />
			</el-form-item>
			<el-form-item label="适用职级" prop="">
				<el-select style="width: 100%"></el-select>
			</el-form-item>
		</el-form>
		<el-table style="width: 100%" tooltip-effect="dark" :data="tableData">
			<el-table-column align="left" label="工资单字段" prop="label" />
			<el-table-column align="left" label="员工可见" prop="value">
				<template #default="{ row }">
					<el-switch v-model="row.value" inline-prompt active-text="是" inactive-text="否" />
				</template>
			</el-table-column>
		</el-table>
	</el-drawer>
</template>

<script setup>
import { ref } from 'vue'

const dialogFormVisible = ref(false)

const tableData = ref([
	{
		label: '等级工资',
		value: 1,
	},
	{
		label: '职称技能工资',
		value: 1,
	},
	{
		label: '保密工资',
		value: 1,
	},
	{
		label: 'CPI工资',
		value: 1,
	},
	{
		label: '津贴',
		value: 1,
	},
	{
		label: '司龄津贴',
		value: 1,
	},
	{
		label: '学历工资',
		value: 1,
	},
	{
		label: '全勤',
		value: 1,
	},
	{
		label: '尽孝基金',
		value: 1,
	},
	{
		label: '节日金',
		value: 1,
	},
])

defineExpose({
	dialogFormVisible,
})
</script>
