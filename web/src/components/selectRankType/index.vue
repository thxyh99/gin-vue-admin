<template>
    <el-select  filterable
    >
      <el-option
        v-for="(item, index) in options"
        :key="index"
        :label="item.name"
        :value="item.ID"
      />
    </el-select>
</template>

<script setup>
import { ref } from 'vue'
import {getRankTypeList} from "@/api/weChat/wcRank";
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
  const table = await getRankTypeList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    options.value = table.data.list
  }

  // if (props.showOptionAll === true) {
  //   options.value.push({ label: '全部', value: '' })
  // }


}
initSelectOptions()

</script>

<style></style>
