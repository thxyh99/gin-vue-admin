<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="入职标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入入职标题" />
       </el-form-item>
        <el-form-item label="姓名:" prop="staffName">
          <el-input v-model="formData.staffName" :clearable="true"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="入职日期:" prop="employmentDate">
          <el-date-picker v-model="formData.employmentDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="入职部门:" prop="employmentDepartment">
        <el-select v-model="formData.employmentDepartment" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="入职职位:" prop="jobPosition">
        <el-select v-model="formData.jobPosition" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="社保电脑号:" prop="socialNumber">
          <el-input v-model="formData.socialNumber" :clearable="true"  placeholder="请输入社保电脑号" />
       </el-form-item>
        <el-form-item label="公积金账号:" prop="accountNumber">
          <el-input v-model="formData.accountNumber" :clearable="true"  placeholder="请输入公积金账号" />
       </el-form-item>
        <el-form-item label="户籍类型:" prop="householdType">
        <el-select v-model="formData.householdType" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="社保公积金缴纳地:" prop="paymentPlace">
        <el-select v-model="formData.paymentPlace" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [255]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="性别(0未知1男2女):" prop="gender">
        <el-select v-model="formData.gender" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="出生年月:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="籍贯:" prop="nativePlace">
        <el-select v-model="formData.nativePlace" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [255]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="民族:" prop="nationality">
        <el-select v-model="formData.nationality" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [50]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="身高:" prop="height">
          <el-input-number v-model="formData.height" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="体重:" prop="weight">
          <el-input-number v-model="formData.weight" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="婚否(1:已婚 2:未婚 3:其他):" prop="marriage">
        <el-select v-model="formData.marriage" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="政治面貌(1:团员 2:党员 3:群众 0:其他):" prop="politicalOutlook">
        <el-select v-model="formData.politicalOutlook" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他):" prop="education">
        <el-select v-model="formData.education" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="专业:" prop="major">
          <el-input v-model="formData.major" :clearable="true"  placeholder="请输入专业" />
       </el-form-item>
        <el-form-item label="毕业院校:" prop="school">
          <el-input v-model="formData.school" :clearable="true"  placeholder="请输入毕业院校" />
       </el-form-item>
        <el-form-item label="毕业时间:" prop="graduationDate">
          <el-date-picker v-model="formData.graduationDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="职称技能证书:" prop="certificate">
           <SelectImage v-model="formData.certificate" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="身份证号码:" prop="idNumber">
          <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入身份证号码" />
       </el-form-item>
        <el-form-item label="身份证地址:" prop="idAddress">
          <el-input v-model="formData.idAddress" :clearable="true"  placeholder="请输入身份证地址" />
       </el-form-item>
        <el-form-item label="银行账号:" prop="cardNumber">
          <el-input v-model="formData.cardNumber" :clearable="true"  placeholder="请输入银行账号" />
       </el-form-item>
        <el-form-item label="开户支行信息:" prop="bank">
          <el-input v-model="formData.bank" :clearable="true"  placeholder="请输入开户支行信息" />
       </el-form-item>
        <el-form-item label="本人联系微信:" prop="contantWechat">
          <el-input v-model="formData.contantWechat" :clearable="true"  placeholder="请输入本人联系微信" />
       </el-form-item>
        <el-form-item label="手机:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入手机" />
       </el-form-item>
        <el-form-item label="常住地址:" prop="homeAddress">
          <el-input v-model="formData.homeAddress" :clearable="true"  placeholder="请输入常住地址" />
       </el-form-item>
        <el-form-item label="证件资料附件上传:" prop="licenseAttment">
           <SelectImage v-model="formData.licenseAttment" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="联系人关系(1:父母 2:配偶 3:子女 0:其他):" prop="relationShip">
          <el-switch v-model="formData.relationShip" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="紧急联系人姓名:" prop="relationName">
          <el-input v-model="formData.relationName" :clearable="true"  placeholder="请输入紧急联系人姓名" />
       </el-form-item>
        <el-form-item label="紧急联系人电话:" prop="relationMobile">
          <el-input v-model="formData.relationMobile" :clearable="true"  placeholder="请输入紧急联系人电话" />
       </el-form-item>
        <el-form-item label="联系人常住地址:" prop="relationAddress">
          <el-input v-model="formData.relationAddress" :clearable="true"  placeholder="请输入联系人常住地址" />
       </el-form-item>
        <el-form-item label="是否经过首席执行官:" prop="isCeopass">
          <el-switch v-model="formData.isCeopass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="入职体检是否正常:" prop="isBodychecknormal">
          <el-switch v-model="formData.isBodychecknormal" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="职级:" prop="jobLevel">
        <el-select v-model="formData.jobLevel" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他):" prop="tryPeriod">
        <el-select v-model="formData.tryPeriod" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="用人部门意见:" prop="employingOpinion">
          <el-input v-model="formData.employingOpinion" :clearable="true"  placeholder="请输入用人部门意见" />
       </el-form-item>
        <el-form-item label="人资部意见:" prop="humanOpinion">
          <el-input v-model="formData.humanOpinion" :clearable="true"  placeholder="请输入人资部意见" />
       </el-form-item>
        <el-form-item label="通知紧急程度:" prop="urgencyNotification">
          <el-input v-model="formData.urgencyNotification" :clearable="true"  placeholder="请输入通知紧急程度" />
       </el-form-item>
        <el-form-item label="入职意见:" prop="onboardingOpinion">
          <el-input v-model="formData.onboardingOpinion" :clearable="true"  placeholder="请输入入职意见" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
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
  createWcStaffEmploymentApplication,
  updateWcStaffEmploymentApplication,
  findWcStaffEmploymentApplication
} from '@/api/employee/wcStaffEmploymentApplication'

defineOptions({
    name: 'WcStaffEmploymentApplicationForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            title: '',
            staffName: '',
            employmentDate: new Date(),
            socialNumber: '',
            accountNumber: '',
            birthday: new Date(),
            height: 0,
            weight: 0,
            major: '',
            school: '',
            graduationDate: new Date(),
            certificate: [],
            idNumber: '',
            idAddress: '',
            cardNumber: '',
            bank: '',
            contantWechat: '',
            mobile: '',
            homeAddress: '',
            licenseAttment: [],
            relationShip: false,
            relationName: '',
            relationMobile: '',
            relationAddress: '',
            isCeopass: false,
            isBodychecknormal: false,
            employingOpinion: '',
            humanOpinion: '',
            urgencyNotification: '',
            onboardingOpinion: '',
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               staffName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               employmentDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               socialNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               accountNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               householdType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               paymentPlace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
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
               education : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               major : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               school : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               graduationDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               certificate : [{
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
               cardNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               bank : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mobile : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               homeAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               licenseAttment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tryPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findWcStaffEmploymentApplication({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffEmploymentApplication
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
               res = await createWcStaffEmploymentApplication(formData.value)
               break
             case 'update':
               res = await updateWcStaffEmploymentApplication(formData.value)
               break
             default:
               res = await createWcStaffEmploymentApplication(formData.value)
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
