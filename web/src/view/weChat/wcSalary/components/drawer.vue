<template>
	<el-drawer
		size="400"
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
		<el-form ref="ruleFormRef" :model="formData" :rules="rules" label-width="80" label-position="right">
			<el-form-item label="工资月份" prop="month">
				<el-date-picker style="width: 100%" v-model="formData.month" type="month" value-format="YYYYMM" />
			</el-form-item>
			<el-form-item label="工资模板" prop="templateId">
				<el-select v-model="formData.templateId">
					<el-option v-for="item in templateList" :key="item.ID" :label="item.name" :value="item.ID" />
				</el-select>
			</el-form-item>
		</el-form>
	</el-drawer>
</template>

<script setup>
import { ref } from 'vue'

import { createWcSalary } from '@/api/weChat/wcSalary'
import { getWcSalaryTemplateList } from '@/api/weChat/wcSalaryTemplate'

import { ElMessage } from 'element-plus'

const props = defineProps({
	modeType: {
		type: String,
		default: 'add',
	},
})

const emits = defineEmits(['onSuccess'])

const dialogFormVisible = ref(false)

const ruleFormRef = ref(null)

const formData = ref({
	month: '',
	templateId: '',
})

const rules = ref({
	month: [{ required: true, message: '请选择月份', trigger: 'change' }],
	templateId: [{ required: true, message: '请选择模板', trigger: 'change' }],
})

const templateList = ref([])
const getWcSalaryTemplate = () => {
	getWcSalaryTemplateList().then((res) => {
		templateList.value = res.data.list
	})
}

const handleOpen = () => {
	getWcSalaryTemplate()
}

const resetForm = (formEl) => {
	if (!formEl) return
	formEl.resetFields()
	dialogFormVisible.value = false
}

const submitForm = async (formEl) => {
	if (!formEl) return
	await formEl.validate((valid) => {
		if (valid) {
			if (props.modeType === 'add') {
				createWcSalary(formData.value).then((res) => {
					ElMessage.success('新增成功')
					emits('onSuccess')
					resetForm(formEl)
				})
			} else {
			}
		}
	})
}

defineExpose({
	dialogFormVisible,
})
</script>
