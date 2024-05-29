<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="入职时间:" prop="employmentDate">
          <el-date-picker v-model="formData.employmentDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="截止日期:" prop="deadline">
          <el-date-picker v-model="formData.deadline" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="司龄:" prop="companyAge">
          <el-input v-model.number="formData.companyAge" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发薪天数:" prop="salaryDays">
          <el-input v-model.number="formData.salaryDays" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="加班情况:" prop="overtime">
          <el-input-number v-model="formData.overtime" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="缺勤情况:" prop="absence">
          <el-input-number v-model="formData.absence" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="费用科目:" prop="expense">
          <el-input v-model="formData.expense" :clearable="true"  placeholder="请输入费用科目" />
       </el-form-item>
        <el-form-item label="职位:" prop="position">
          <el-input v-model="formData.position" :clearable="true"  placeholder="请输入职位" />
       </el-form-item>
        <el-form-item label="学历:" prop="education">
          <el-input v-model="formData.education" :clearable="true"  placeholder="请输入学历" />
       </el-form-item>
        <el-form-item label="等级工资:" prop="gradeSalary">
          <el-input-number v-model="formData.gradeSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="职称技能津贴:" prop="skillSalary">
          <el-input-number v-model="formData.skillSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="津贴:" prop="allowancesSalary">
          <el-input-number v-model="formData.allowancesSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="司龄津贴:" prop="ageAllowanceSalary">
          <el-input-number v-model="formData.ageAllowanceSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="学历工资:" prop="educationSalary">
          <el-input-number v-model="formData.educationSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="全勤:" prop="fullSalary">
          <el-input-number v-model="formData.fullSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="尽孝基金:" prop="filialSalary">
          <el-input-number v-model="formData.filialSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="业绩奖:" prop="performanceSalary">
          <el-input-number v-model="formData.performanceSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="节日金:" prop="festivalSalary">
          <el-input-number v-model="formData.festivalSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="生日金:" prop="birthdaySalary">
          <el-input-number v-model="formData.birthdaySalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="岗位补助:" prop="jobSubsidy">
          <el-input-number v-model="formData.jobSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="组长补助:" prop="leaderSubsidy">
          <el-input-number v-model="formData.leaderSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="电话补助:" prop="telephoneSubsidy">
          <el-input-number v-model="formData.telephoneSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="宿舍补助:" prop="dormitorySubsidy">
          <el-input-number v-model="formData.dormitorySubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="交通补助:" prop="transportSubsidy">
          <el-input-number v-model="formData.transportSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="餐补:" prop="mealSubsidy">
          <el-input-number v-model="formData.mealSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="出差补助:" prop="travelSubsidy">
          <el-input-number v-model="formData.travelSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="分包装间:" prop="packagingSubsidy">
          <el-input-number v-model="formData.packagingSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="宿舍费:" prop="dormSubsidy">
          <el-input-number v-model="formData.dormSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="其他(补助):" prop="otherSubsidy">
          <el-input-number v-model="formData.otherSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="合计(补助):" prop="totalSubsidy">
          <el-input-number v-model="formData.totalSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="加班费:" prop="overtimeFee">
          <el-input-number v-model="formData.overtimeFee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="缺勤:" prop="absenceFee">
          <el-input-number v-model="formData.absenceFee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="合计(加班及缺勤):" prop="totalFee">
          <el-input-number v-model="formData.totalFee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="应发工资:" prop="shouldPay">
          <el-input-number v-model="formData.shouldPay" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="养老保险:" prop="pensionPay">
          <el-input-number v-model="formData.pensionPay" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="医疗保险:" prop="medicalPay">
          <el-input-number v-model="formData.medicalPay" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="失业保险:" prop="unemploymentPay">
          <el-input-number v-model="formData.unemploymentPay" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="住房公积金:" prop="housingPay">
          <el-input-number v-model="formData.housingPay" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="预存尽孝基金:" prop="preFilialSalary">
          <el-input-number v-model="formData.preFilialSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="其他:" prop="other">
          <el-input-number v-model="formData.other" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="已发:" prop="sent">
          <el-input-number v-model="formData.sent" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="实发小计:" prop="actualPay">
          <el-input-number v-model="formData.actualPay" :precision="2" :clearable="true"></el-input-number>
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
  createWcStaffSalary,
  updateWcStaffSalary,
  findWcStaffSalary
} from '@/api/weChat/wcStaffSalary'

defineOptions({
    name: 'WcStaffSalaryForm'
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
            staffId: 0,
            name: '',
            employmentDate: new Date(),
            deadline: new Date(),
            companyAge: 0,
            salaryDays: 0,
            overtime: 0,
            absence: 0,
            expense: '',
            position: '',
            education: '',
            gradeSalary: 0,
            skillSalary: 0,
            allowancesSalary: 0,
            ageAllowanceSalary: 0,
            educationSalary: 0,
            fullSalary: 0,
            filialSalary: 0,
            performanceSalary: 0,
            festivalSalary: 0,
            birthdaySalary: 0,
            jobSubsidy: 0,
            leaderSubsidy: 0,
            telephoneSubsidy: 0,
            dormitorySubsidy: 0,
            transportSubsidy: 0,
            mealSubsidy: 0,
            travelSubsidy: 0,
            packagingSubsidy: 0,
            dormSubsidy: 0,
            otherSubsidy: 0,
            totalSubsidy: 0,
            overtimeFee: 0,
            absenceFee: 0,
            totalFee: 0,
            shouldPay: 0,
            pensionPay: 0,
            medicalPay: 0,
            unemploymentPay: 0,
            housingPay: 0,
            preFilialSalary: 0,
            other: 0,
            sent: 0,
            actualPay: 0,
        })
// 验证规则
const rule = reactive({
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
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               deadline : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               companyAge : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               salaryDays : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               overtime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               absence : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               expense : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               position : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               education : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gradeSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               skillSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               allowancesSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ageAllowanceSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               educationSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fullSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               filialSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               performanceSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               festivalSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               birthdaySalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               leaderSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               telephoneSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dormitorySubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               transportSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mealSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               travelSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               packagingSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dormSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               otherSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               overtimeFee : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               absenceFee : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalFee : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               shouldPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pensionPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               medicalPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               unemploymentPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               housingPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               preFilialSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               other : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sent : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               actualPay : [{
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
      const res = await findWcStaffSalary({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewcStaffSalary
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
               res = await createWcStaffSalary(formData.value)
               break
             case 'update':
               res = await updateWcStaffSalary(formData.value)
               break
             default:
               res = await createWcStaffSalary(formData.value)
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
