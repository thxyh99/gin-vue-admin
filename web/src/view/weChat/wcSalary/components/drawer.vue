<template>
	<el-drawer size="400" v-model="dialogFormVisible" :show-close="false">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">添加</span>
				<div>
					<el-button type="primary" @click="enterDialog">确 定</el-button>
					<el-button @click="closeDialog">取 消</el-button>
				</div>
			</div>
		</template>
		<el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule">
			<el-form-item label="工资月份" prop="">
				<el-date-picker style="width: 100%" type="month" />
			</el-form-item>
			<el-form-item label="工资类型" prop="">
				<el-select style="width: 100%"></el-select>
			</el-form-item>
			<el-form-item label="适用职级" prop="">
				<el-select style="width: 100%"></el-select>
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
