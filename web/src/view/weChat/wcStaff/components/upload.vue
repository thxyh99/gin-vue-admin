<template>
	<el-upload
		action="/api/wcFile/upload"
		:fileList="uploadList"
		:data="{ staffId: props.staffId, type: props.type }"
		:on-remove="handleRemove"
		:on-preview="handlePreview"
		:on-success="handleSuccess"
	>
		<el-button size="small" type="primary" :disabled="props.type !== 10 && uploadList.length">点击上传</el-button>
	</el-upload>
</template>

<script setup>
import { ref, watchEffect } from 'vue'

import { ElMessage } from 'element-plus'
import { deleteFile } from '@/api/weChat/wcStaff'
const props = defineProps({
	staffId: {
		type: String,
		default: '',
	},
	type: {
		type: Number,
		default: 1,
	},
	fileList: {
		type: Array,
		default: () => [],
	},
})

const uploadList = ref([])

watchEffect(() => {
	uploadList.value = (props.fileList || []).map((item) => {
		return {
			name: item.name,
			url: item.privateAccessURL,
			ID: item.ID,
		}
	})
})

const emits = defineEmits(['update:fileList', 'updateInfoSuccess'])

const handleRemove = async (uploadFile, uploadFiles) => {
	const res = await deleteFile({ ID: uploadFile.ID })
	if (res.code === 0) {
		emits('update:fileList', uploadFiles)
		emits('updateInfoSuccess')
		ElMessage.success('删除成功')
	}
}
const handlePreview = (item) => {
	window.open(item.url)
}

const handleSuccess = (response, uploadFiles) => {
	if (response.code === 0) {
		uploadList.value.push({
			name: uploadFiles.name,
			url: uploadFiles.privateAccessURL,
			ID: uploadFiles.ID,
		})
		emits('update:fileList', uploadList.value)
		emits('updateInfoSuccess')
		ElMessage.success('上传成功')
	}
}
</script>
