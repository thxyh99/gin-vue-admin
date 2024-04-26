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
        <el-form-item label="身份证号:" prop="idNumber">
          <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入身份证号" />
       </el-form-item>
        <el-form-item label="身份证地址:" prop="idAddress">
          <el-input v-model="formData.idAddress" :clearable="true"  placeholder="请输入身份证地址" />
       </el-form-item>
        <el-form-item label="户籍类型(1:本地城镇 2:本地农村 3:外地城镇[省内] 4:外地农村[省内] 5:外地城镇[省外] 6:外地农村[省外]):" prop="householdType">
          <el-input v-model.number="formData.householdType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="籍贯:" prop="nativePlace">
          <el-input v-model="formData.nativePlace" :clearable="true"  placeholder="请输入籍贯" />
       </el-form-item>
        <el-form-item label="民族:" prop="nationality">
          <el-input v-model="formData.nationality" :clearable="true"  placeholder="请输入民族" />
       </el-form-item>
        <el-form-item label="身高:" prop="height">
          <el-input-number v-model="formData.height" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="体重:" prop="weight">
          <el-input-number v-model="formData.weight" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="婚否(1:已婚 2:未婚 3:其他):" prop="marriage">
          <el-input v-model.number="formData.marriage" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="政治面貌(1:团员 2:党员 3:群众 0:其他):" prop="politicalOutlook">
          <el-input v-model.number="formData.politicalOutlook" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="常住地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入常住地址" />
       </el-form-item>
        <el-form-item label="社保电脑号:" prop="socialNumber">
          <el-input v-model="formData.socialNumber" :clearable="true"  placeholder="请输入社保电脑号" />
       </el-form-item>
        <el-form-item label="公积金账号:" prop="accountNumber">
          <el-input v-model="formData.accountNumber" :clearable="true"  placeholder="请输入公积金账号" />
       </el-form-item>
        <el-form-item label="社保公积金缴纳地:" prop="paymentPlace">
          <el-input v-model="formData.paymentPlace" :clearable="true"  placeholder="请输入社保公积金缴纳地" />
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
  createWcStaffInfo,
  updateWcStaffInfo,
  findWcStaffInfo
} from '@/api/weChat/wcStaffInfo'

defineOptions({
    name: 'WcStaffInfoForm'
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
            idNumber: '',
            idAddress: '',
            householdType: 0,
            birthday: new Date(),
            nativePlace: '',
            nationality: '',
            height: 0,
            weight: 0,
            marriage: 0,
            politicalOutlook: 0,
            address: '',
            socialNumber: '',
            accountNumber: '',
            paymentPlace: '',
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
               idNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               idAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               householdType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               birthday : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nativePlace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nationality : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               height : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               weight : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               marriage : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               politicalOutlook : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               address : [{
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
      const res = await findWcStaffInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffInfo
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
               res = await createWcStaffInfo(formData.value)
               break
             case 'update':
               res = await updateWcStaffInfo(formData.value)
               break
             default:
               res = await createWcStaffInfo(formData.value)
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
