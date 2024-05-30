<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金):" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="员工ID:" prop="staffId">
          <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="一级部门:" prop="departmentFirst">
          <el-input v-model="formData.departmentFirst" :clearable="true"  placeholder="请输入一级部门" />
       </el-form-item>
        <el-form-item label="二级部门:" prop="departmentSecond">
          <el-input v-model="formData.departmentSecond" :clearable="true"  placeholder="请输入二级部门" />
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
        <el-form-item label="职位/职级:" prop="position">
          <el-input v-model="formData.position" :clearable="true"  placeholder="请输入职位/职级" />
       </el-form-item>
        <el-form-item label="学历:" prop="education">
          <el-input v-model="formData.education" :clearable="true"  placeholder="请输入学历" />
       </el-form-item>
        <el-form-item label="等级工资:" prop="gradeSalary">
          <el-input-number v-model="formData.gradeSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="岗位工资/职称技能津贴:" prop="skillSalary">
          <el-input-number v-model="formData.skillSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="保密工资:" prop="secrecySalary">
          <el-input-number v-model="formData.secrecySalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="CPI工资:" prop="cpiSalary">
          <el-input-number v-model="formData.cpiSalary" :precision="2" :clearable="true"></el-input-number>
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
        <el-form-item label="高温补贴:" prop="highTemperatureSubsidy">
          <el-input-number v-model="formData.highTemperatureSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="全勤:" prop="fullSalary">
          <el-input-number v-model="formData.fullSalary" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="组长补助:" prop="leaderSubsidy">
          <el-input-number v-model="formData.leaderSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="分包装间:" prop="packagingSubsidy">
          <el-input-number v-model="formData.packagingSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="岗位补助:" prop="jobSubsidy">
          <el-input-number v-model="formData.jobSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="宿舍补助:" prop="dormitorySubsidy">
          <el-input-number v-model="formData.dormitorySubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="餐补:" prop="mealSubsidy">
          <el-input-number v-model="formData.mealSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="电话补助:" prop="telephoneSubsidy">
          <el-input-number v-model="formData.telephoneSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="交通补助:" prop="transportSubsidy">
          <el-input-number v-model="formData.transportSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="出差补助:" prop="travelSubsidy">
          <el-input-number v-model="formData.travelSubsidy" :precision="2" :clearable="true"></el-input-number>
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
        <el-form-item label="专项抵扣:" prop="specialTax">
          <el-input-number v-model="formData.specialTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="本次应缴税所得额:" prop="thisPayTax">
          <el-input-number v-model="formData.thisPayTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="上次累计应缴预缴所得额:" prop="lastTotalPayTax">
          <el-input-number v-model="formData.lastTotalPayTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="本次累计应缴预缴所得额:" prop="thisTotalPayTax">
          <el-input-number v-model="formData.thisTotalPayTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="累计税额:" prop="totalPayTax">
          <el-input-number v-model="formData.totalPayTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="上次累计已扣个税:" prop="lastTotalPaidTax">
          <el-input-number v-model="formData.lastTotalPaidTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="本次累计已缴个税:" prop="thisTotalPaidTax">
          <el-input-number v-model="formData.thisTotalPaidTax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="本次应扣个税:" prop="thisShouldPayTax">
          <el-input-number v-model="formData.thisShouldPayTax" :precision="2" :clearable="true"></el-input-number>
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
        <el-form-item label="月度提成补发:" prop="monthlyCommission">
          <el-input-number v-model="formData.monthlyCommission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="年度提成:" prop="annualCommission">
          <el-input-number v-model="formData.annualCommission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="子品牌经营绩效奖励:" prop="subbrandCommission">
          <el-input-number v-model="formData.subbrandCommission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="奖金点:" prop="bonusPoints">
          <el-input-number v-model="formData.bonusPoints" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="出勤系数:" prop="attendanceRatio">
          <el-input-number v-model="formData.attendanceRatio" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="月度评分:" prop="monthlyRating">
          <el-input-number v-model="formData.monthlyRating" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="个人计点:" prop="personalCount">
          <el-input-number v-model="formData.personalCount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="迟到乐捐:" prop="lateDonation">
          <el-input-number v-model="formData.lateDonation" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="月度奖金:" prop="monthlyBonus">
          <el-input-number v-model="formData.monthlyBonus" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="区域:" prop="region">
          <el-input v-model="formData.region" :clearable="true"  placeholder="请输入区域" />
       </el-form-item>
        <el-form-item label="全年任务（万）:" prop="annualTask">
          <el-input-number v-model="formData.annualTask" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="月度提成点:" prop="monthlyCommissionRating">
          <el-input-number v-model="formData.monthlyCommissionRating" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="提成比例:" prop="commissionRatio">
          <el-input-number v-model="formData.commissionRatio" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="所责市场:" prop="market">
          <el-input v-model="formData.market" :clearable="true"  placeholder="请输入所责市场" />
       </el-form-item>
        <el-form-item label="月度出货指标:" prop="monthlyShipmentTarget">
          <el-input-number v-model="formData.monthlyShipmentTarget" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="当月实际出货:" prop="actualShipment">
          <el-input-number v-model="formData.actualShipment" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="月度提成基数:" prop="monthlyCommissionBase">
          <el-input-number v-model="formData.monthlyCommissionBase" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="完成率%:" prop="completionRate">
          <el-input-number v-model="formData.completionRate" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="绩效考评分数:" prop="performanceScore">
          <el-input-number v-model="formData.performanceScore" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="不在岗缺勤:" prop="absenceWork">
          <el-input-number v-model="formData.absenceWork" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="出货提成:" prop="shippingCommission">
          <el-input-number v-model="formData.shippingCommission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="个人销售提成:" prop="personalSalesCommission">
          <el-input-number v-model="formData.personalSalesCommission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="市场差补服务费:" prop="marketServiceFee">
          <el-input-number v-model="formData.marketServiceFee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="月度奖励:" prop="monthlyReward">
          <el-input-number v-model="formData.monthlyReward" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="品牌负责人月度补贴:" prop="monthlyLeaderSubsidy">
          <el-input-number v-model="formData.monthlyLeaderSubsidy" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="大店开发提成:" prop="storeCommission">
          <el-input-number v-model="formData.storeCommission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="项目数:" prop="projectNumber">
          <el-input v-model.number="formData.projectNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手工费:" prop="handicraftCosts">
          <el-input-number v-model="formData.handicraftCosts" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="业绩:" prop="performance">
          <el-input-number v-model="formData.performance" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="提成:" prop="commission">
          <el-input-number v-model="formData.commission" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备用字段1:" prop="stringValue1">
          <el-input v-model="formData.stringValue1" :clearable="true"  placeholder="请输入备用字段1" />
       </el-form-item>
        <el-form-item label="备用字段2:" prop="stringValue2">
          <el-input v-model="formData.stringValue2" :clearable="true"  placeholder="请输入备用字段2" />
       </el-form-item>
        <el-form-item label="备用字段3:" prop="stringValue3">
          <el-input v-model="formData.stringValue3" :clearable="true"  placeholder="请输入备用字段3" />
       </el-form-item>
        <el-form-item label="备用字段4:" prop="stringValue4">
          <el-input v-model="formData.stringValue4" :clearable="true"  placeholder="请输入备用字段4" />
       </el-form-item>
        <el-form-item label="备用字段5:" prop="stringValue5">
          <el-input v-model="formData.stringValue5" :clearable="true"  placeholder="请输入备用字段5" />
       </el-form-item>
        <el-form-item label="备用字段1:" prop="decimalValue1">
          <el-input-number v-model="formData.decimalValue1" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备用字段2:" prop="decimalValue2">
          <el-input-number v-model="formData.decimalValue2" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备用字段3:" prop="decimalValue3">
          <el-input-number v-model="formData.decimalValue3" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备用字段4:" prop="decimalValue4">
          <el-input-number v-model="formData.decimalValue4" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备用字段5:" prop="decimalValue5">
          <el-input-number v-model="formData.decimalValue5" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备用字段1:" prop="intValue1">
          <el-input v-model.number="formData.intValue1" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备用字段2:" prop="intValue2">
          <el-input v-model.number="formData.intValue2" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备用字段3:" prop="intValue3">
          <el-input v-model.number="formData.intValue3" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备用字段4:" prop="intValue4">
          <el-input v-model.number="formData.intValue4" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备用字段5:" prop="intValue5">
          <el-input v-model.number="formData.intValue5" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备用字段1:" prop="dateValue1">
          <el-date-picker v-model="formData.dateValue1" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="备用字段2:" prop="dateValue2">
          <el-date-picker v-model="formData.dateValue2" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="备用字段3:" prop="dateValue3">
          <el-date-picker v-model="formData.dateValue3" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="备用字段4:" prop="dateValue4">
          <el-date-picker v-model="formData.dateValue4" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="备用字段5:" prop="dateValue5">
          <el-date-picker v-model="formData.dateValue5" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
            type: 0,
            staffId: 0,
            departmentFirst: '',
            departmentSecond: '',
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
            secrecySalary: 0,
            cpiSalary: 0,
            allowancesSalary: 0,
            ageAllowanceSalary: 0,
            educationSalary: 0,
            filialSalary: 0,
            performanceSalary: 0,
            festivalSalary: 0,
            birthdaySalary: 0,
            highTemperatureSubsidy: 0,
            fullSalary: 0,
            leaderSubsidy: 0,
            packagingSubsidy: 0,
            jobSubsidy: 0,
            dormitorySubsidy: 0,
            mealSubsidy: 0,
            telephoneSubsidy: 0,
            transportSubsidy: 0,
            travelSubsidy: 0,
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
            specialTax: 0,
            thisPayTax: 0,
            lastTotalPayTax: 0,
            thisTotalPayTax: 0,
            totalPayTax: 0,
            lastTotalPaidTax: 0,
            thisTotalPaidTax: 0,
            thisShouldPayTax: 0,
            preFilialSalary: 0,
            other: 0,
            sent: 0,
            actualPay: 0,
            monthlyCommission: 0,
            annualCommission: 0,
            subbrandCommission: 0,
            bonusPoints: 0,
            attendanceRatio: 0,
            monthlyRating: 0,
            personalCount: 0,
            lateDonation: 0,
            monthlyBonus: 0,
            remark: '',
            region: '',
            annualTask: 0,
            monthlyCommissionRating: 0,
            commissionRatio: 0,
            market: '',
            monthlyShipmentTarget: 0,
            actualShipment: 0,
            monthlyCommissionBase: 0,
            completionRate: 0,
            performanceScore: 0,
            absenceWork: 0,
            shippingCommission: 0,
            personalSalesCommission: 0,
            marketServiceFee: 0,
            monthlyReward: 0,
            monthlyLeaderSubsidy: 0,
            storeCommission: 0,
            projectNumber: 0,
            handicraftCosts: 0,
            performance: 0,
            commission: 0,
            stringValue1: '',
            stringValue2: '',
            stringValue3: '',
            stringValue4: '',
            stringValue5: '',
            decimalValue1: 0,
            decimalValue2: 0,
            decimalValue3: 0,
            decimalValue4: 0,
            decimalValue5: 0,
            intValue1: 0,
            intValue2: 0,
            intValue3: 0,
            intValue4: 0,
            intValue5: 0,
            dateValue1: new Date(),
            dateValue2: new Date(),
            dateValue3: new Date(),
            dateValue4: new Date(),
            dateValue5: new Date(),
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
