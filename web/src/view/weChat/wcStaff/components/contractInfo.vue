<template>
	<div>
		<template v-if="rewcStaffAgreement && rewcStaffAgreement.length">
			<el-descriptions title="合同信息" border v-for="(item, index) in rewcStaffAgreement" :key="index">
				<template #extra>
					<el-button type="primary" size="small">编 辑</el-button>
					<el-button type="primary" size="small" @click="handleDelete(item.ID)">删 除</el-button>
				</template>
				<el-descriptions-item label="合同公司" label-align="right">{{ item.company }}</el-descriptions-item>
				<el-descriptions-item label="合同类型" label-align="right">{{ item.typeText }}</el-descriptions-item>
				<el-descriptions-item label="合同起始日" label-align="right">{{
					formatDate(item.startDay)
				}}</el-descriptions-item>
				<el-descriptions-item label="合同到期日" label-align="right">{{
					formatDate(item.endDay)
				}}</el-descriptions-item>
				<el-descriptions-item label="续签次数" label-align="right">{{ item.times }}</el-descriptions-item>
				<el-descriptions-item label-align="right"></el-descriptions-item>
				<el-descriptions-item label="合同附件" label-align="right" :span="3">
					<el-button type="text">{{ item.attachment.name }}</el-button>
				</el-descriptions-item>
			</el-descriptions>
		</template>

		<div class="btn" @click="handleAdd">
			<span>+ 添加合同</span>
		</div>
		<el-drawer size="800" v-model="dialogFormVisible" :show-close="false">
			<template #title>
				<div class="flex justify-between items-center">
					<span class="text-lg">{{ modeType ? '添加' : '修改' }}</span>
					<div>
						<el-button type="primary" @click="enterDialog">确 定</el-button>
						<el-button @click="closeDialog">取 消</el-button>
					</div>
				</div>
			</template>

			<el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
				<el-form-item label="合同公司:" prop="company">
					<el-input v-model="formData.company" :clearable="true" placeholder="请输入合同公司" />
				</el-form-item>
				<el-form-item label="合同类型:" prop="type">
					<el-select v-model="formData.type" placeholder="选择合同类型">
						<el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="合同起始日:" prop="startDay">
					<el-date-picker
						v-model="formData.startDay"
						type="date"
						style="width: 100%"
						placeholder="选择日期"
						:clearable="true"
					/>
				</el-form-item>
				<el-form-item label="合同到期日:" prop="endDay">
					<el-date-picker
						v-model="formData.endDay"
						type="date"
						style="width: 100%"
						placeholder="选择日期"
						:clearable="true"
					/>
				</el-form-item>
				<el-form-item label="续签次数:" prop="times">
					<el-input v-model.number="formData.times" :clearable="true" placeholder="请输入续签次数" />
				</el-form-item>
				<el-form-item label="合同附件:" prop="fileId">
					<el-upload
						v-model:file-list="fileList"
						action="/api/wcFile/upload"
						multiple
						:data="{ staffId: props.id, type: 1 }"
						:on-preview="handlePreview"
						:on-success="handleUploadSuccess"
					>
						<el-button type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-form>
		</el-drawer>
	</div>
</template>

<script setup>
import { ref, reactive } from 'vue'

import { updateWcStaffAgreement, deleteWcStaffAgreement } from '@/api/weChat/wcStaffAgreement'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	rewcStaffAgreement: {
		type: Object,
		default: () => {
			return {}
		},
	},
})

const emits = defineEmits(['updateInfoSuccess'])

const dialogFormVisible = ref(false)
const modeType = ref(true)
const formData = ref({})

const types = [
	{ label: '固定期限劳动合同', value: 1 },
	{ label: '无固定期限劳动合同', value: 2 },
	{ label: '实习协议', value: 3 },
	{ label: '外包协议', value: 4 },
	{ label: '劳务派遣合同', value: 5 },
	{ label: '返聘协议', value: 6 },
	{ label: '培训协议', value: 7 },
]
const rule = reactive({
	staffId: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	company: [
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
	type: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	startDay: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	endDay: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
})
const handleAdd = () => {
	modeType.value = true
	dialogFormVisible.value = true
}

const handleDelete = (ID) => {
	ElMessageBox.confirm('是否删除该合同?', '提示', {
		confirmButtonText: '删除',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(async () => {
			const res = await deleteWcStaffAgreement({ ID })
			if(res.code === 0) {
				emits('updateInfoSuccess')
			}
		})
		.catch(() => {
			ElMessage({
				type: 'info',
				message: '取消删除',
			})
		})
}

const handleUploadSuccess = (res) => {
	console.log(res)
	formData.value.fileId = res.data.ID
}

const closeDialog = () => {
	dialogFormVisible.value = false
}

const elFormRef = ref()
const enterDialog = () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		const res = await updateWcStaffAgreement({
			...formData.value,
			fileId: formData.value.fileId ? formData.value.fileId : 0,
			staffId: Number(props.id),
			// ID: props.rewcStaffBank.ID,
		})
		if (res.code === 0) {
			emits('updateInfoSuccess')
			dialogFormVisible.value = false
		}
	})
}
</script>

<style scoped>
.btn {
	margin-top: 20px;
	padding: 10px;
	width: 100%;
	border: 1px dashed #ccc;
	display: flex;
	align-items: center;
	justify-content: center;
	cursor: pointer;
}
</style>
