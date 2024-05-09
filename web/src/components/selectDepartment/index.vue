<template>
    <el-select multiple
    >
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
</template>

<script setup>
import { ref } from 'vue'
import {getAllFullDepartmentList} from "@/api/weChat/wcDepartment";
const page = ref(1)
const pageSize = ref(1000000)
const searchInfo = ref({})


const props = defineProps({
  // 是否显示"全部"选项
  showOptionAll: {
    type: Boolean,
    default: true
  }
})

const options = ref([])
const initSelectOptions = async () => {
  const table = await getAllFullDepartmentList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    table.data.list.forEach(item => {
        options.value.push({
          label: item.name,
          value: item.id
        })
    })
  }

  // if (props.showOptionAll === true) {
  //   options.value.push({ label: '全部', value: '' })
  // }


}
initSelectOptions()

</script>

<style></style>
