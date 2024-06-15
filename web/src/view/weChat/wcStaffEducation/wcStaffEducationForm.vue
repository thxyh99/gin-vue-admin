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
        <el-form-item label="学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他):" prop="education">
          <el-input v-model.number="formData.education" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="毕业院校:" prop="school">
          <el-input v-model="formData.school" :clearable="true"  placeholder="请输入毕业院校" />
       </el-form-item>
        <el-form-item label="毕业日期:" prop="date">
          <el-date-picker v-model="formData.date" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="专业:" prop="major">
          <el-input v-model="formData.major" :clearable="true"  placeholder="请输入专业" />
       </el-form-item>
        <el-form-item label="职称证书:" prop="professionalCertificate">
          <el-input v-model="formData.professionalCertificate" :clearable="true"  placeholder="请输入职称证书" />
       </el-form-item>
        <el-form-item label="技能证书:" prop="skillCertificate">
          <el-input v-model="formData.skillCertificate" :clearable="true"  placeholder="请输入技能证书" />
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
  createWcStaffEducation,
  updateWcStaffEducation,
  findWcStaffEducation
} from '@/api/weChat/wcStaffEducation'

defineOptions({
    name: 'WcStaffEducationForm'
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
            education: 0,
            school: '',
            date: new Date(),
            major: '',
  professionalCertificate: '',
  skillCertificate: '',
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
               education : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               school : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               date : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               major : [{
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
      const res = await findWcStaffEducation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffEducation
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
               res = await createWcStaffEducation(formData.value)
               break
             case 'update':
               res = await updateWcStaffEducation(formData.value)
               break
             default:
               res = await createWcStaffEducation(formData.value)
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
