<template>
  <el-button
    type="primary"
    icon="download"
    @click="exportExcelFunc"
  >导出花名册</el-button>
</template>

<script setup>
const props = defineProps({
  templateId: {
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
  staffId: {
    type: Number,
    default: ''
  },
  historyDate: {
    type: String,
    default: ''
  },
  employmentDateRange: {
    type: String,
    default: ''
  },
  departmentIds: {
    type: String,
    default: ''
  },
  keyword: {
    type: String,
    default: ''
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
  const url = `${baseUrl}/wcStaff/exportExcel?templateID=${props.templateId}&staffId=${props.staffId}&historyDate=${props.historyDate}&employmentDateRange=${props.employmentDateRange}&departmentIds=${props.departmentIds}&keyword=${props.keyword}${params ? '&' + params : ''}`
  // /wcStaff/exportExcel?templateID=roster&staffId=10&historyDate=2024-06-30&employmentDateRange=2024-01-01,2025-07-01&departmentIds=2,3&keyword=RD_SZ&limit=9999
  window.open(url, '_blank')
}
</script>
