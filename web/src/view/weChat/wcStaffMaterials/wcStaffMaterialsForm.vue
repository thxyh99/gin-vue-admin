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
        <el-form-item label="身份证(人像):" prop="idCardPortrait">
          <el-input v-model="formData.idCardPortrait" :clearable="true"  placeholder="请输入身份证(人像)" />
       </el-form-item>
        <el-form-item label="身份证(国徽):" prop="idCardNational">
          <el-input v-model="formData.idCardNational" :clearable="true"  placeholder="请输入身份证(国徽)" />
       </el-form-item>
        <el-form-item label="学历证书:" prop="educationCertificate">
          <el-input v-model="formData.educationCertificate" :clearable="true"  placeholder="请输入学历证书" />
       </el-form-item>
        <el-form-item label="学位证书:" prop="degreeCertificate">
          <el-input v-model="formData.degreeCertificate" :clearable="true"  placeholder="请输入学位证书" />
       </el-form-item>
        <el-form-item label="前公司离职证明:" prop="resignationCertificate">
          <el-input v-model="formData.resignationCertificate" :clearable="true"  placeholder="请输入前公司离职证明" />
       </el-form-item>
        <el-form-item label="入职补充登记表:" prop="onboardingForm">
          <el-input v-model="formData.onboardingForm" :clearable="true"  placeholder="请输入入职补充登记表" />
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
  createWcStaffMaterials,
  updateWcStaffMaterials,
  findWcStaffMaterials
} from '@/api/weChat/wcStaffMaterials'

defineOptions({
    name: 'WcStaffMaterialsForm'
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
            idCardPortrait: '',
            idCardNational: '',
            educationCertificate: '',
            degreeCertificate: '',
            resignationCertificate: '',
            onboardingForm: '',
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
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWcStaffMaterials({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffMaterials
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
               res = await createWcStaffMaterials(formData.value)
               break
             case 'update':
               res = await updateWcStaffMaterials(formData.value)
               break
             default:
               res = await createWcStaffMaterials(formData.value)
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
