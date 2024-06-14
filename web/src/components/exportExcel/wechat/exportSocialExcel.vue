<template>
  <el-button
    type="primary"
    icon="download"
    @click="exportExcelFunc"
  >{{ btnName }}</el-button>
</template>

<script setup>
const props = defineProps({
  templateId: {
    type: String,
    required: true
  },
  type: {
    type: String,
    required: true
  },
  condition: {
    type: Object,
    default: () => ({})
  },
  limit: {
    type: Number,
    default: 0
  },
  offset: {
    type: Number,
    default: 0
  },
  order: {
    type: String,
    default: ''
  },
  btnName: {
    type: String,
    required: true,
    default: "导出"
  }
})

import { ElMessage } from 'element-plus'

const exportExcelFunc = async() => {
  if (props.templateId === '') {
    ElMessage.error('组件未设置模板ID')
    return
  }
  const baseUrl = import.meta.env.VITE_BASE_API
  const paramsCopy = JSON.parse(JSON.stringify(props.condition))
  if (props.limit) {
    paramsCopy.limit = props.limit
  }
  if (props.offset) {
    paramsCopy.offset = props.offset
  }
  if (props.order) {
    paramsCopy.order = props.order
  }
  const params = Object.entries(paramsCopy)
    .map(([key, value]) => `${encodeURIComponent(key)}=${encodeURIComponent(value)}`)
    .join('&')
  const url = `${baseUrl}/wcStaffSocial/exportExcel?templateID=${props.templateId}&type=${props.type}${params ? '&' + params : ''}`

  window.open(url, '_blank')
}
</script>
