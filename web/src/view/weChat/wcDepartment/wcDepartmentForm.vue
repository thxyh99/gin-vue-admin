<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="部门ID:" prop="departmentId">
          <el-input v-model.number="formData.departmentId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="部门名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入部门名称" />
       </el-form-item>
        <el-form-item label="英文名称:" prop="nameEn">
          <el-input v-model="formData.nameEn" :clearable="true"  placeholder="请输入英文名称" />
       </el-form-item>
        <el-form-item label="部门负责人:" prop="departmentLeader">
          <el-input v-model="formData.departmentLeader" :clearable="true"  placeholder="请输入部门负责人" />
       </el-form-item>
        <el-form-item label="父部门ID:" prop="parentid">
          <el-input v-model.number="formData.parentid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:" prop="order">
          <el-input v-model.number="formData.order" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWcDepartment,
  updateWcDepartment,
  findWcDepartment
} from '@/api/weChat/wcDepartment'

defineOptions({
    name: 'WcDepartmentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            departmentId: 0,
            name: '',
            nameEn: '',
            departmentLeader: '',
            parentid: 0,
            order: 0,
        })
// 验证规则
const rule = reactive({
               departmentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               parentid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               order : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWcDepartment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcDepartment
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWcDepartment(formData.value)
               break
             case 'update':
               res = await updateWcDepartment(formData.value)
               break
             default:
               res = await createWcDepartment(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
