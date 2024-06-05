<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">添 加</span>
				<div>
					<el-button type="primary" @click="submitForm(ruleFormRef)">确 定</el-button>
					<el-button @click="resetForm(ruleFormRef)">取 消</el-button>
				</div>
			</div>
		</template>
		<el-form ref="ruleFormRef" :model="formData" :rules="rules" label-width="120" label-position="right">
			<el-form-item label="工资单模板名称" prop="name">
				<el-input v-model="formData.name" />
			</el-form-item>
			<el-form-item label="工资类型">
				<el-select v-model="formData.type" @change="getWcSalaryFields">
					<el-option label="基本工资" :value="1"></el-option>
					<el-option label="集团经营绩效奖励" :value="2"></el-option>
					<el-option label="节日金" :value="3"></el-option>
					<el-option label="半年奖" :value="4"></el-option>
					<el-option label="年度奖金" :value="5"></el-option>
					<el-option label="总部职能体系月度奖金" :value="6"></el-option>
					<el-option label="总部金纳斯市场体系月度奖金" :value="7"></el-option>
					<el-option label="总部调理中心体系月度奖金" :value="8"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="适用职级体系" prop="rankType">
				<el-select v-model="formData.rankType">
					<el-option v-for="(item, index) in typeList" :key="index" :label="item.name" :value="item.ID" />
				</el-select>
			</el-form-item>
		</el-form>
		<el-table :data="formData.fields" style="width: 100%" v-loading="loading">
			<el-table-column align="left" label="工资单字段" prop="name" />
			<el-table-column align="left" label="员工可见">
				<template #default="{ row }">
					<el-switch v-model="row.IsRequired" :active-value="1" :inactive-value="0"></el-switch>
				</template>
			</el-table-column>
		</el-table>
	</el-drawer>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'

import {
	getRankTypeList,
	getWcSalaryFieldsByType,
	createWcSalaryTemplate,
	updateWcSalaryTemplate,
} from '@/api/weChat/wcSalaryTemplate'

const props = defineProps({
	modeType: {
		type: String,
		default: 'add',
	},
	rowInfo: {
		type: Object,
		default: () => ({}),
	},
})

const ruleFormRef = ref()
const dialogFormVisible = ref(false)
const loading = ref(false)
const formData = ref({
	name: '',
	type: 1,
	rankType: undefined,
	fields: [],
})

const rules = reactive({
	name: [{ required: true, message: '', trigger: 'blur' }],
	rankType: [
		{
			required: true,
			message: '',
			trigger: 'change',
		},
	],
})

const typeList = ref([])

const getRankFormTypeList = async () => {
	const res = await getRankTypeList()
	typeList.value = res.data.list
}

const getWcSalaryFields = async () => {
	loading.value = true
	const res = await getWcSalaryFieldsByType({ type: formData.value.type })
	loading.value = false
	formData.value.fields = res.data.list
}

const submitForm = async (formEl) => {
	if (!formEl) return
	await formEl.validate(async (valid) => {
		if (valid) {
			if (props.modeType === 'add') {
				// 新增
				const res = await createWcSalaryTemplate(formData.value)
			} else {
				// 编辑
				await updateWcSalaryTemplate({ ID: props.rowInfo.ID, ...formData })
			}
		}
	})
}

const resetForm = (formEl) => {
	if (!formEl) return
	formEl.resetFields()
	dialogFormVisible.value = false
}

onMounted(() => {
	getRankFormTypeList()
	getWcSalaryFields()
})

defineExpose({
	dialogFormVisible,
})
</script>
