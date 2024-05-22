<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">修改</span>
				<div>
					<el-button type="primary" @click="enterDialog">确 定</el-button>
					<el-button @click="closeDialog">取 消</el-button>
				</div>
			</div>
		</template>

		<el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
			<el-form-item label="银行卡号:" prop="cardNumber">
				<el-input v-model="formData.cardNumber" :clearable="true" placeholder="请输入银行卡号" />
			</el-form-item>
			<el-form-item label="开户行:" prop="bank">
				<el-input v-model="formData.bank" :clearable="true" placeholder="请输入开户行" />
			</el-form-item>
		</el-form>
	</el-drawer>
</template>

<script setup>
import { ref, reactive, watchEffect } from 'vue'
import { updateWcStaffBank } from '@/api/weChat/wcStaffBank'

const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	rewcStaffBank: {
		type: Object,
		default: () => {
			return {}
		},
	},
})

const emits = defineEmits(['updateInfoSuccess'])

const dialogFormVisible = ref(false)

const rule = reactive({
	staffId: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	cardNumber: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
		{
			whitespace: true,
			message: '不能只输入空格',
			trigger: ['input', 'blur'],
		},
	],
	bank: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
		{
			whitespace: true,
			message: '不能只输入空格',
			trigger: ['input', 'blur'],
		},
	],
})

const formData = ref({})
const elFormRef = ref()
watchEffect(() => {
	formData.value = { ...props.rewcStaffBank }
})

const closeDialog = () => {
	formData.value = { ...props.rewcStaffBank }
	dialogFormVisible.value = false
}

const enterDialog = () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		const res = await updateWcStaffBank({
			...formData.value,
			staffId: Number(props.id),
			ID: props.rewcStaffBank.ID,
		})
		if (res.code === 0) {
			emits('updateInfoSuccess')
			dialogFormVisible.value = false
		}
	})
}

defineExpose({
	dialogFormVisible,
})
</script>
