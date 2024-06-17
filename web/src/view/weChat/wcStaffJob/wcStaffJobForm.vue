<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID(SSO):" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="企微成员UserID:" prop="userid">
          <el-input v-model="formData.userid" :clearable="true"  placeholder="请输入企微成员UserID" />
       </el-form-item>
        <el-form-item label="员工类型(1:全职 2:兼职 3:实习 4:劳务外包 5:无类型):" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="员工状态(1:试用 2:正式 3:待离职 0:无状态):" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他):" prop="tryPeriod">
          <el-input v-model.number="formData.tryPeriod" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="入职容大日期:" prop="employmentDate">
          <el-date-picker v-model="formData.employmentDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="入职总部日期:" prop="employmentDate">
          <el-date-picker v-model="formData.employmentHeadquarterDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="转正日期:" prop="formalDate">
          <el-date-picker v-model="formData.formalDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="职位名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入职位名称" />
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
  createWcStaffJob,
  updateWcStaffJob,
  findWcStaffJob
} from '@/api/weChat/wcStaffJob'

defineOptions({
    name: 'WcStaffJobForm'
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
            userId: 0,
            userid: '',
            type: 0,
            status: 0,
            tryPeriod: 0,
            employmentDate: new Date(),
            formalDate: new Date(),
            name: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tryPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               formalDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
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
      const res = await findWcStaffJob({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffJob
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
               res = await createWcStaffJob(formData.value)
               break
             case 'update':
               res = await updateWcStaffJob(formData.value)
               break
             default:
               res = await createWcStaffJob(formData.value)
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
