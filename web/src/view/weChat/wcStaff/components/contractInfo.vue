<template>
	<div>
		<template v-if="rewcStaffAgreement && rewcStaffAgreement.length">
			<el-descriptions title="合同信息" border v-for="(item, index) in rewcStaffAgreement" :key="index">
				<template #extra>
					<el-button type="primary" size="small" @click="handleEdit(item, index)">编 辑</el-button>
					<el-button type="primary" size="small" @click="handleDelete(item.ID)">删 除</el-button>
				</template>
				<el-descriptions-item label="合同公司" label-align="right">{{ item.companyText }}</el-descriptions-item>
				<el-descriptions-item label="合同类型" label-align="right">{{ item.typeText }}</el-descriptions-item>
				<el-descriptions-item label="合同起始日" label-align="right">{{
					formatDate(item.startDay)
				}}</el-descriptions-item>
				<el-descriptions-item label="合同到期日" label-align="right">{{
					formatDate(item.endDay)
				}}</el-descriptions-item>
				<el-descriptions-item label="续签次数" label-align="right">{{ item.timesText }}</el-descriptions-item>
				<el-descriptions-item label="合同期限（月）"  label-align="right" >{{ item.period }}</el-descriptions-item>
				<el-descriptions-item label="特殊备注"  label-align="right" >{{ item.remark }}</el-descriptions-item>
				<el-descriptions-item label="合同附件" label-align="right" :span="3">
					<el-button type="text" @click="handlePreview(item.attachment.privateAccessURL)">{{
						item.attachment.name
					}}</el-button>
				</el-descriptions-item>
			</el-descriptions>
		</template>
		<div class="btn" @click="handleAdd">
			<span>+ 添加合同</span>
		</div>
		<el-drawer size="800" v-model="dialogFormVisible" :show-close="false" @closed="closeDialog">
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
          <el-select v-model="formData.company" placeholder="请选择合同公司">
            <el-option v-for="agreementCompany in agreementCompanies" :key="agreementCompany.value" :label="agreementCompany.label" :value="agreementCompany.value"></el-option>
          </el-select>
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
          <el-select v-model="formData.times" placeholder="请选择续签次数">
            <el-option v-for="time in times" :key="time.value" :label="time.label" :value="time.value"></el-option>
          </el-select>
				</el-form-item>
        <el-form-item label="特殊备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入特殊备注" />
        </el-form-item>
				<el-form-item label="合同附件:" prop="fileId">
					<el-upload
						action="/api/wcFile/upload"
						:file-list="fileList"
						:data="{ staffId: props.id, type: 1 }"
						:before-upload="beforeAvatarUpload"
						:on-remove="handleRemove"
						:on-success="handleUploadSuccess"
					>
						<el-button type="primary" :disabled="fileList.length">点击上传</el-button>
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
const ind = ref(0)
const formData = ref({})

const fileList = ref([])

const types = [
	{ label: '固定期限劳动合同', value: 1 },
	{ label: '无固定期限劳动合同', value: 2 },
	{ label: '实习协议', value: 3 },
	{ label: '外包协议', value: 4 },
	{ label: '劳务派遣合同', value: 5 },
	{ label: '返聘协议', value: 6 },
	{ label: '培训协议', value: 7 },
]

const agreementCompanies = [
  {label: "深圳市容大生物科技股份有限公司",value:1},
  {label: "深圳市容大生物科技股份有限公司东莞分公司",value:2},
  {label: "深圳市容大霜道云连锁健康互联有限公司",value:3},
  {label: "东莞市容大生物科技有限公司",value:4},
  {label: "广东国韵健康科技有限公司",value:5},
  {label: "深圳市青春纳斯健康科技有限公司",value:6},
  {label: "广东金生缘健康科技有限公司",value:7},
  {label: "广东金因康健康科技有限公司",value:8},
  {label: "深圳市容大文化传媒有限公司",value:9},
]

const times = [
  { label: null, value: 0 },
  { label: '1', value: 1 },
  { label: '2', value: 2 },
  { label: '3', value: 3 },
  { label: '其他', value: 4 },
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

const handleEdit = (item, index) => {
	ind.value = index
	modeType.value = false
	for (const key in item) {
		formData.value[key] = item[key]
	}
	if (item.attachment.ID) {
		fileList.value = [
			{
				name: item.attachment.name,
				url: item.attachment.privateAccessURL,
				id: item.attachment.ID,
			},
		]
	}

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
			if (res.code === 0) {
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

const handleUploadSuccess = (res, uploadFiles) => {
	console.log(uploadFiles)
	fileList.value.push({
		name: uploadFiles.name,
		url: uploadFiles.privateAccessURL,
		ID: uploadFiles.ID,
	})
	formData.value.fileId = res.data.ID
}

const beforeAvatarUpload = () => {
	if (fileList.value.length > 0) {
		ElMessage.error('一份合同只能一个附件，请删除后重新上传！')
		return false
	}
}

const handlePreview = (url) => {
	window.open(url)
}

const handleRemove = (file) => {
	fileList.value = []
}

const closeDialog = () => {
	formData.value = {}
	fileList.value = []
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
			ID: modeType.value ? undefined : props.rewcStaffAgreement[ind.value].ID,
		})
		if (res.code === 0) {
			emits('updateInfoSuccess')
			dialogFormVisible.value = false
			formData.value = {}
			fileList.value = []
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
