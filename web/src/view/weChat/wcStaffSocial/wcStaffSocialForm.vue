<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="账号(社保:电脑号|个人社保号 公积金:个人账号):" prop="account">
          <el-input v-model="formData.account" :clearable="true"  placeholder="请输入账号(社保:电脑号|个人社保号 公积金:个人账号)" />
       </el-form-item>
        <el-form-item label="类型(1:深圳社保 2:深圳公积金 3:东莞社保 4:东莞公积金):" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="证件类型(1:身份证 2:港澳通行证):" prop="credentialType">
          <el-input v-model.number="formData.credentialType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="证件号码:" prop="credentialNumber">
          <el-input v-model="formData.credentialNumber" :clearable="true"  placeholder="请输入证件号码" />
       </el-form-item>
        <el-form-item label="社保缴费合计:" prop="totalSocial">
          <el-input-number v-model="formData.totalSocial" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="社保缴费个人合计:" prop="totalSocialSelf">
          <el-input-number v-model="formData.totalSocialSelf" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="社保缴费单位合计:" prop="totalSocialUnit">
          <el-input-number v-model="formData.totalSocialUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="养老缴纳基数:" prop="pensionBase">
          <el-input-number v-model="formData.pensionBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="养老个人缴费:" prop="pensionSelf">
          <el-input-number v-model="formData.pensionSelf" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="养老单位缴费:" prop="pensionUnit">
          <el-input-number v-model="formData.pensionUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="医疗缴纳基数:" prop="medicalBase">
          <el-input-number v-model="formData.medicalBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="医疗个人缴费:" prop="medicalSelf">
          <el-input-number v-model="formData.medicalSelf" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="医疗单位缴费:" prop="medicalUnit">
          <el-input-number v-model="formData.medicalUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="失业缴纳基数:" prop="unemployedBase">
          <el-input-number v-model="formData.unemployedBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="失业个人缴费:" prop="unemployedSelf">
          <el-input-number v-model="formData.unemployedSelf" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="失业单位缴费:" prop="unemployedUnit">
          <el-input-number v-model="formData.unemployedUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="工商保险缴费基数:" prop="injuryInsuranceBase">
          <el-input-number v-model="formData.injuryInsuranceBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="工商保险单位缴费:" prop="injuryInsuranceUnit">
          <el-input-number v-model="formData.injuryInsuranceUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="生育医疗缴费基数:" prop="birthBase">
          <el-input-number v-model="formData.birthBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="生育医疗单位缴费:" prop="birthUnit">
          <el-input-number v-model="formData.birthUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="公积金缴费合计:" prop="totalHousing">
          <el-input-number v-model="formData.totalHousing" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="公积金个人缴费合计:" prop="totalHousingSelf">
          <el-input-number v-model="formData.totalHousingSelf" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="公积金公司缴费合计:" prop="totalHousingUnit">
          <el-input-number v-model="formData.totalHousingUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="缴存基数:" prop="housingBase">
          <el-input-number v-model="formData.housingBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="个人缴存比例:" prop="housingRatioSelf">
          <el-input-number v-model="formData.housingRatioSelf" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="单位缴存比例:" prop="housingRatioUnit">
          <el-input-number v-model="formData.housingRatioUnit" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="费款所属期起:" prop="periodStart">
          <el-input v-model="formData.periodStart" :clearable="true"  placeholder="请输入费款所属期起" />
       </el-form-item>
        <el-form-item label="费款所属期止:" prop="periodEnd">
          <el-input v-model="formData.periodEnd" :clearable="true"  placeholder="请输入费款所属期止" />
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
  createWcStaffSocial,
  updateWcStaffSocial,
  findWcStaffSocial
} from '@/api/weChat/wcStaffSocial'

defineOptions({
    name: 'WcStaffSocialForm'
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
            account: '',
            type: 0,
            staffId: 0,
            name: '',
            credentialType: 0,
            credentialNumber: '',
            totalSocial: 0,
            totalSocialSelf: 0,
            totalSocialUnit: 0,
            pensionBase: 0,
            pensionSelf: 0,
            pensionUnit: 0,
            medicalBase: 0,
            medicalSelf: 0,
            medicalUnit: 0,
            unemployedBase: 0,
            unemployedSelf: 0,
            unemployedUnit: 0,
            injuryInsuranceBase: 0,
            injuryInsuranceUnit: 0,
            birthBase: 0,
            birthUnit: 0,
            totalHousing: 0,
            totalHousingSelf: 0,
            totalHousingUnit: 0,
            housingBase: 0,
            housingRatioSelf: 0,
            housingRatioUnit: 0,
            periodStart: '',
            periodEnd: '',
        })
// 验证规则
const rule = reactive({
               account : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               staffId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               credentialType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               credentialNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalSocial : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalSocialSelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalSocialUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pensionBase : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pensionSelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pensionUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               medicalBase : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               medicalSelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               medicalUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               unemployedBase : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               unemployedSelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               unemployedUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
                injuryInsuranceBase : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
                injuryInsuranceUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               birthBase : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               birthUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalHousing : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalHousingSelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalHousingUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               housingBase : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               housingRatioSelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               housingRatioUnit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               periodStart : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               periodEnd : [{
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
      const res = await findWcStaffSocial({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffSocial
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
               res = await createWcStaffSocial(formData.value)
               break
             case 'update':
               res = await updateWcStaffSocial(formData.value)
               break
             default:
               res = await createWcStaffSocial(formData.value)
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
