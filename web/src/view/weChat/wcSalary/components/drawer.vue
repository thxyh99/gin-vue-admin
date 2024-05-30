<template>
	<el-drawer size="400" v-model="dialogFormVisible" :show-close="false">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">添 加</span>
				<div>
					<el-button type="primary">确 定</el-button>
					<el-button>取 消</el-button>
				</div>
			</div>
		</template>
		<el-form label-width="80" label-position="right">
			<el-form-item label="工资月份">
				<el-input />
			</el-form-item>
			<el-form-item label="工资模板">
				<el-select></el-select>
			</el-form-item>
			<el-form-item label="适用职级">
				<el-select></el-select>
			</el-form-item>
		</el-form>
	</el-drawer>
</template>

<script setup>
import { ref } from 'vue'

const dialogFormVisible = ref(false)

defineExpose({
	dialogFormVisible,
})
</script>
