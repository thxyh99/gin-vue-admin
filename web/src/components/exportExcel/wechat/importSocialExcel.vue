<template>
  <el-upload
    :action="url"
    :show-file-list="false"
    :on-success="handleSuccess"
    :multiple="false"
  >
    <el-button
      type="primary"
      icon="upload"
    >{{ btnName }}</el-button>
  </el-upload>

</template>

<script setup>
import { ElMessage } from 'element-plus'

const baseUrl = import.meta.env.VITE_BASE_API

const props = defineProps({
  templateId: {
    type: String,
    required: true
  },
  type: {
    type: Number,
    required: true
  },
  btnName: {
    type: String,
    required: true,
    default: "导入"
  }
})

const emit = defineEmits(['on-success'])

const url = `${baseUrl}/wcStaffSocial/importExcel?templateID=${props.templateId}&type=${props.type}`

const handleSuccess = (res) => {
  if (res.code === 0) {
    ElMessage.success('导入成功')
    emit('on-success')
  } else {
    ElMessage.error(res.msg)
  }
}
</script>
