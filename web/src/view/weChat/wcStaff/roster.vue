<template>
	<div>
		<div class="gva-search-box">
			<el-image
				style="width: 120px; height: 120px; border-radius: 10px"
				src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
			/>
			<span>xxx</span>
		</div>
		<div class="gva-table-box">
			<el-tabs v-model="activeName">
				<el-tab-pane label="基础信息" name="first">
					<baseInfo />
				</el-tab-pane>
				<el-tab-pane label="合同信息" name="second">
					<contractInfo />
				</el-tab-pane>
				<el-tab-pane label="材料附件" name="third">
					<material />
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import baseInfo from './components/baseInfo.vue'
import contractInfo from './components/contractInfo.vue'
import material from './components/material.vue'

import { getObtainEmployeeRoster } from '@/api/weChat/wcStaff'

const route = useRoute()

const activeName = ref('first')
const id = ref(route.query.id)

const getData = () => {
	getObtainEmployeeRoster({ id: id.value })
}

onMounted(() => {
	getData()
})
</script>

<style></style>
