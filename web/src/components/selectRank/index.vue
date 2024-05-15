<template>
    <el-select  filterable
    >
      <el-option
        v-for="item in options"
        :key="item.ID"
        :label="item.name"
        :value="item.ID"
      />
    </el-select>
</template>

<script setup>
import { ref } from 'vue'
import {getRankListByRankType} from "@/api/weChat/wcRank";
const page = ref(1)
const pageSize = ref(1000000)


const props = defineProps({
  // 是否显示"全部"选项
  showOptionAll: {
    type: Boolean,
    default: true
  }
})

const options = ref([])
const initSelectOptions = async (type) => {
  const table = await getRankListByRankType({page: page.value, pageSize: pageSize.value, type})
  if (table.code === 0) {
    options.value = table.data.list
  }

  // if (props.showOptionAll === true) {
  //   options.value.push({ label: '全部', value: '' })
  // }


}
// initSelectOptions()
defineExpose({
  initSelectOptions
})
</script>

<style></style>
