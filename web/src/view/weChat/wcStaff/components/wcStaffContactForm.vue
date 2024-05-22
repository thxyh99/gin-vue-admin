<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
				<div>
					<el-button type="primary" @click="enterDialog">确 定</el-button>
					<el-button @click="closeDialog">取 消</el-button>
				</div>
			</div>
		</template>

		<el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
			<el-form-item label="紧急联系人姓名:" prop="name">
				<el-input v-model="formData.name" :clearable="true" placeholder="请输入紧急联系人姓名" />
			</el-form-item>
			<el-form-item label="联系人关系:" prop="relationship">
				<el-select v-model="formData.relationship" placeholder="选择联系人关系">
					<el-option
						v-for="relationship in relationships"
						:key="relationship.value"
						:label="relationship.label"
						:value="relationship.value"
					></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="联系人电话:" prop="mobile">
				<el-input v-model="formData.mobile" :clearable="true" placeholder="请输入联系人电话" />
			</el-form-item>
			<el-form-item label="联系人常住地址:" prop="address">
				<el-input v-model="formData.address" :clearable="true" placeholder="请输入联系人常住地址" />
			</el-form-item>
		</el-form>
	</el-drawer>
</template>

<script setup>
import { ref, reactive, watchEffect } from 'vue'
import { updateWcStaffContact } from '@/api/weChat/wcStaffContact'

const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	rewcStaffContact: {
		type: Object,
		default: () => {
			return {}
		},
	},
})

const relationships = [
	{ label: '其他', value: 0 },
	{ label: '父母', value: 1 },
	{ label: '配偶', value: 2 },
	{ label: '子女', value: 3 },
]

const rule = reactive({
	staffId: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	name: [
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
	relationship: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	mobile: [
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

const emits = defineEmits(['updateInfoSuccess'])

const dialogFormVisible = ref(false)

const formData = ref({})
const elFormRef = ref()
watchEffect(() => {
	formData.value = { ...props.rewcStaffContact }
})

const closeDialog = () => {
	formData.value = { ...props.rewcStaffContact }
	dialogFormVisible.value = false
}

const enterDialog = () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		const res = await updateWcStaffContact({
			...formData.value,
			staffId: Number(props.id),
			ID: props.rewcStaffContact.ID,
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
