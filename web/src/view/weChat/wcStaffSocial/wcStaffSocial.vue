<template>
	<div class="gva-search-box">
		<el-tabs v-model="activeName" class="demo-tabs">
			<el-tab-pane label="深圳社保" name="first">
				<componentA />
			</el-tab-pane>
			<el-tab-pane label="深圳公积金" name="second">
				<componentB />
			</el-tab-pane>
			<el-tab-pane label="东莞社保" name="third">
				<componentC />
			</el-tab-pane>
			<el-tab-pane label="东莞公积金" name="fourth">
				<componentD />
			</el-tab-pane>
		</el-tabs>
	</div>
</template>

<script setup>
import { ref } from 'vue'
import componentA from './components/componentA.vue'
import componentB from './components/componentB.vue'
import componentC from './components/componentC.vue'
import componentD from './components/componentD.vue'

const activeName = ref('first')
</script>
