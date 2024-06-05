<template>
	<el-drawer
		v-loading="loading"
		size="800"
		v-model="dialogFormVisible"
		:show-close="false"
		@opened="handleOpen"
		@closed="resetForm(ruleFormRef)"
	>
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
				<el-select v-model="formData.type" @change="getWcSalaryFields" :disabled="props.modeType === 'edit'">
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
				<el-select v-model="formData.rankType" :disabled="props.modeType === 'edit'">
					<el-option v-for="(item, index) in typeList" :key="index" :label="item.name" :value="item.ID" />
				</el-select>
			</el-form-item>
		</el-form>
		<el-table :data="formData.fields" style="width: 100%">
			<el-table-column align="left" label="工资单字段" prop="name" />
			<el-table-column align="left" label="员工可见">
				<template #default="{ row }">
					<el-switch v-model="row.isVisible" :active-value="1" :inactive-value="0"></el-switch>
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
	findWcSalaryTemplate,
} from '@/api/weChat/wcSalaryTemplate'

import { ElMessage } from 'element-plus'

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

const emits = defineEmits(['onSuccess'])

const ruleFormRef = ref()
const dialogFormVisible = ref(false)
const loading = ref(false)
const formData = ref({
	name: '',
	type: undefined,
	rankType: undefined,
	fields: [],
})

const rules = reactive({
	name: [{ required: true, message: '', trigger: 'blur' }],
	type: [
		{
			required: true,
			message: '',
			trigger: 'change',
		},
	],
	rankType: [
		{
			required: true,
			message: '',
			trigger: 'change',
		},
	],
})

const typeList = ref([])

const handleOpen = () => {
	if (props.modeType !== 'add') {
		findWcSalaryTemplate({ ID: props.rowInfo.ID }).then((res) => {
			formData.value.name = res.data.rewcSalaryTemplate.name
			formData.value.type = res.data.rewcSalaryTemplate.type
			formData.value.rankType = res.data.rewcSalaryTemplate.rankType
			formData.value.fields = res.data.rewcSalaryTemplate.fields
		})
	}
}

const getRankFormTypeList = async () => {
	const res = await getRankTypeList()
	typeList.value = res.data.list
}

const getWcSalaryFields = async (type) => {
	loading.value = true
	const res = await getWcSalaryFieldsByType({ type })
	loading.value = false
	formData.value.fields = res.data.list
}

const submitForm = async (formEl) => {
	if (!formEl) return
	await formEl.validate((valid) => {
		if (valid) {
			loading.value = true
			if (props.modeType === 'add') {
				// 新增
				createWcSalaryTemplate(formData.value).then((res) => {
					ElMessage.success('新增成功')
					emits('onSuccess')
					resetForm(formEl)
				})
			} else {
				// 编辑
				updateWcSalaryTemplate({ ID: props.rowInfo.ID, ...formData.value }).then((res) => {
					ElMessage.success('编辑成功')
					emits('onSuccess')
					resetForm(formEl)
				})
			}
		}
	})
}

const resetForm = (formEl) => {
	if (!formEl) return
	formEl.resetFields()
	formData.value.type = undefined
	formData.value.fields = []
	dialogFormVisible.value = false
}

onMounted(() => {
	getRankFormTypeList()
})

defineExpose({
	dialogFormVisible,
})
</script>
