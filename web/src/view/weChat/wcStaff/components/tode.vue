<template>
	<el-drawer size="800" v-model="todoShow" title="待办事项" destroy-on-close>
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">待办事项</span>
			</div>
		</template>
		<el-tabs>
			<el-tab-pane>
				<template #label>
					<div>全部（9）</div>
				</template>
			</el-tab-pane>
			<el-tab-pane>
				<template #label>
					<div>超期未处理（3）</div>
				</template>
			</el-tab-pane>
			<el-tab-pane>
				<template #label>
					<div>今日到期（3）</div>
				</template>
			</el-tab-pane>
			<el-tab-pane>
				<template #label>
					<div>提前提醒（3）</div>
				</template>
			</el-tab-pane>
		</el-tabs>
		<div class="todo-container">
			<template v-for="item in 4" :key="item">
				<div class="item p-2">
					<div class="font-bold mb-2">武则天合同到期</div>
					<div class="mb-4">合同到期日：2024-06-18</div>
					<el-button type="primary">前往处理</el-button>
				</div>
				<el-divider></el-divider>
			</template>
		</div>
	</el-drawer>
</template>

<script setup>
import { ref } from 'vue'

const todoShow = ref(false)

defineExpose({ todoShow })
</script>

<style scoped lang="scss">
.tode-container {
	.item:hover {
		background-color: #000;
	}
}

::v-deep(.el-divider) {
	margin: 12px 0;
}
</style>
