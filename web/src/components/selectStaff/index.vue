<template>
    <el-select filterable
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
import {getSimpleStaffList} from "@/api/weChat/wcStaff";
const page = ref(1)
const pageSize = ref(10000)
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
  const table = await getSimpleStaffList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    table.data.list.forEach(item => {
        options.value.push({
          label: item.name+"【"+item.jobNum+"】",
          value: item.ID
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
