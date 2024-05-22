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
			<el-form-item label="学历:" prop="education">
				<el-select v-model="formData.education" placeholder="选择学历">
					<el-option
						v-for="education in educations"
						:key="education.value"
						:label="education.label"
						:value="education.value"
					></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="毕业院校:" prop="school">
				<el-input v-model="formData.school" :clearable="true" placeholder="请输入毕业院校" />
			</el-form-item>
			<el-form-item label="毕业日期:" prop="date">
				<el-date-picker
					v-model="formData.date"
					type="date"
					style="width: 100%"
					placeholder="选择日期"
					:clearable="true"
				/>
			</el-form-item>
			<el-form-item label="专业:" prop="major">
				<el-input v-model="formData.major" :clearable="true" placeholder="请输入专业" />
			</el-form-item>
			<el-form-item label="学历津贴:" prop="educationPay">
				<el-input-number v-model="formData.educationPay" style="width: 100%" :precision="2" :clearable="true" />
				<!-- <el-input v-model="formData.educationPay" :clearable="true" placeholder="请输入学历津贴" /> -->
			</el-form-item>
			<el-form-item label="职称技能津贴:" prop="skillPay">
				<el-input-number v-model="formData.skillPay" style="width: 100%" :precision="2" :clearable="true" />
				<!-- <el-input v-model="formData.skillPay" :clearable="true" placeholder="请输入职称技能津贴" /> -->
			</el-form-item>
			<el-form-item label="职称/技能证书:" prop="certificate">
				<el-input v-model="formData.certificate" :clearable="true" placeholder="请输入职称/技能证书" />
				<!-- <el-upload
					v-model:file-list="fileList"
					action="/api/wcFile/upload"
					multiple
					:data="{ staffId: props.id, type: 10 }"
					:on-preview="handlePreview"
					:on-success="handleUploadSuccess"
				>
					<el-button type="primary">点击上传</el-button>
				</el-upload> -->
			</el-form-item>
		</el-form>
	</el-drawer>
</template>

<script setup>
import { ref, reactive, watchEffect } from 'vue'
import { updateWcStaffEducation } from '@/api/weChat/wcStaffEducation'

const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	rewcStaffEducation: {
		type: Object,
		default: () => {
			return {}
		},
	},
})

const emits = defineEmits(['updateInfoSuccess'])

const educations = [
	{ label: '大专以下', value: 1 },
	{ label: '大专', value: 2 },
	{ label: '专升本/自考本科', value: 3 },
	{ label: '本科', value: 4 },
	{ label: '重点本科', value: 5 },
	{ label: '硕士', value: 6 },
	{ label: '重点硕士', value: 7 },
	{ label: '博士', value: 8 },
]

const rule = reactive({
	staffId: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	education: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	school: [
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
	date: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	major: [
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
	educationPay: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	skillPay: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
})

const dialogFormVisible = ref(false)

const formData = ref({})
const elFormRef = ref()

watchEffect(() => {
	formData.value = { ...props.rewcStaffEducation }
})

const fileList = ref([])
const handleUploadSuccess = (response, uploadFile) => {
	console.log(response);
	console.log(fileList.value);
	fileList.value.map(item => {
		return {
			name: item.response.data.name,
			url: item.response.data.privateAccessURL
		}
	})
}
const handlePreview = (uploadFile) => {
	window.open(uploadFile.response.data.privateAccessURL)
}

const closeDialog = () => {
	formData.value = { ...props.rewcStaffEducation }
	dialogFormVisible.value = false
}

const enterDialog = () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		const res = await updateWcStaffEducation({
			...formData.value,
			staffId: Number(props.id),
			ID: props.rewcStaffEducation.ID,
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
