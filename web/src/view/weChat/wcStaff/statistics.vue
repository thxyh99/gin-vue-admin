<template>
	<div>
		<div class="gva-search-box">
			<el-form ref="elSearchFormRef" :inline="true" :model="searchInfo">
				<el-form-item label="统计月份">
					<el-date-picker
						style="width: 300px"
						type="monthrange"
						range-separator="至"
						start-placeholder="开始月份"
						end-placeholder="结束月份"
						value-format="YYYYMM"
					/>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" icon="search">查询</el-button>
					<el-button icon="refresh">重置</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="gva-table-box">
			<el-row>
				<el-col :span="24">
					<staffNumber ref="staffNumberRef" />
				</el-col>
				<el-col :span="24">
					<entryNumber ref="entryNumberRef" />
				</el-col>
				<el-col :span="24">
					<regularNumber ref="regularNumberRef" />
				</el-col>
				<el-col :span="24">
					<transferNumber ref="transferNumberRef" />
				</el-col>
				<el-col :span="24">
					<resigneNumber ref="resigneNumberRef" />
				</el-col>
			</el-row>
		</div>
	</div>
</template>

<script setup>
import { ref } from 'vue'
import staffNumber from './components/staffNumber.vue'
import entryNumber from './components/entryNumber.vue'
import regularNumber from './components/regularNumber.vue'
import transferNumber from './components/transferNumber.vue'
import resigneNumber from './components/resigneNumber.vue'

const searchInfo = ref({})

const staffNumberRef = ref(null)
const entryNumberRef = ref(null)
const regularNumberRef = ref(null)
const transferNumberRef = ref(null)
const resigneNumberRef = ref(null)
</script>

<style scoped lang="scss"></style>
