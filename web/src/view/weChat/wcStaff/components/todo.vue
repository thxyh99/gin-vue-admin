<template>
	<el-drawer size="800" v-model="todoShow" title="待办事项" destroy-on-close>
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">待办事项</span>
			</div>
		</template>
		<el-tabs v-model="activeName">
			<el-tab-pane v-for="(item, index) in props.todoList" :key="index">
				<template #label>
					<div>{{ item.label }}（{{ item.list.length }}）</div>
				</template>
				<div class="todo-container">
					<template v-for="(itm, ind) in item.list" :key="ind">
						<div class="item p-2">
							<div class="font-bold mb-2">{{ itm.title }}</div>
							<div class="mb-4">{{ itm.label }}：{{ itm.value }}</div>
							<!-- <el-button type="primary">前往处理</el-button> -->
						</div>
						<el-divider></el-divider>
					</template>
				</div>
			</el-tab-pane>
		</el-tabs>
	</el-drawer>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
	todoList: {
		type: Array,
		default: () => [],
	},
})

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
