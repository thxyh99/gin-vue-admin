<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="员工ID" prop="staffId" width="120" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
         <el-table-column align="left" label="入职时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.employmentDate) }}</template>
         </el-table-column>
         <el-table-column align="left" label="截止日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.deadline) }}</template>
         </el-table-column>
        <el-table-column align="left" label="司龄" prop="companyAge" width="120" />
        <el-table-column align="left" label="发薪天数" prop="salaryDays" width="120" />
        <el-table-column align="left" label="加班情况" prop="overtime" width="120" />
        <el-table-column align="left" label="缺勤情况" prop="absence" width="120" />
        <el-table-column align="left" label="费用科目" prop="expense" width="120" />
        <el-table-column align="left" label="职位" prop="position" width="120" />
        <el-table-column align="left" label="学历" prop="education" width="120" />
        <el-table-column align="left" label="等级工资" prop="gradeSalary" width="120" />
        <el-table-column align="left" label="职称技能津贴" prop="skillSalary" width="120" />
        <el-table-column align="left" label="津贴" prop="allowancesSalary" width="120" />
        <el-table-column align="left" label="司龄津贴" prop="ageAllowanceSalary" width="120" />
        <el-table-column align="left" label="学历工资" prop="educationSalary" width="120" />
        <el-table-column align="left" label="全勤" prop="fullSalary" width="120" />
        <el-table-column align="left" label="尽孝基金" prop="filialSalary" width="120" />
        <el-table-column align="left" label="业绩奖" prop="performanceSalary" width="120" />
        <el-table-column align="left" label="节日金" prop="festivalSalary" width="120" />
        <el-table-column align="left" label="生日金" prop="birthdaySalary" width="120" />
        <el-table-column align="left" label="岗位补助" prop="jobSubsidy" width="120" />
        <el-table-column align="left" label="组长补助" prop="leaderSubsidy" width="120" />
        <el-table-column align="left" label="电话补助" prop="telephoneSubsidy" width="120" />
        <el-table-column align="left" label="宿舍补助" prop="dormitorySubsidy" width="120" />
        <el-table-column align="left" label="交通补助" prop="transportSubsidy" width="120" />
        <el-table-column align="left" label="餐补" prop="mealSubsidy" width="120" />
        <el-table-column align="left" label="出差补助" prop="travelSubsidy" width="120" />
        <el-table-column align="left" label="分包装间" prop="packagingSubsidy" width="120" />
        <el-table-column align="left" label="宿舍费" prop="dormSubsidy" width="120" />
        <el-table-column align="left" label="其他(补助)" prop="otherSubsidy" width="120" />
        <el-table-column align="left" label="合计(补助)" prop="totalSubsidy" width="120" />
        <el-table-column align="left" label="加班费" prop="overtimeFee" width="120" />
        <el-table-column align="left" label="缺勤" prop="absenceFee" width="120" />
        <el-table-column align="left" label="合计(加班及缺勤)" prop="totalFee" width="120" />
        <el-table-column align="left" label="应发工资" prop="shouldPay" width="120" />
        <el-table-column align="left" label="养老保险" prop="pensionPay" width="120" />
        <el-table-column align="left" label="医疗保险" prop="medicalPay" width="120" />
        <el-table-column align="left" label="失业保险" prop="unemploymentPay" width="120" />
        <el-table-column align="left" label="住房公积金" prop="housingPay" width="120" />
        <el-table-column align="left" label="预存尽孝基金" prop="preFilialSalary" width="120" />
        <el-table-column align="left" label="其他" prop="other" width="120" />
        <el-table-column align="left" label="已发" prop="sent" width="120" />
        <el-table-column align="left" label="实发小计" prop="actualPay" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffSalaryFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="员工ID:"  prop="staffId" >
              <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入员工ID" />
            </el-form-item>
            <el-form-item label="姓名:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="入职时间:"  prop="employmentDate" >
              <el-date-picker v-model="formData.employmentDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="截止日期:"  prop="deadline" >
              <el-date-picker v-model="formData.deadline" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="司龄:"  prop="companyAge" >
              <el-input v-model.number="formData.companyAge" :clearable="true" placeholder="请输入司龄" />
            </el-form-item>
            <el-form-item label="发薪天数:"  prop="salaryDays" >
              <el-input v-model.number="formData.salaryDays" :clearable="true" placeholder="请输入发薪天数" />
            </el-form-item>
            <el-form-item label="加班情况:"  prop="overtime" >
              <el-input-number v-model="formData.overtime"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="缺勤情况:"  prop="absence" >
              <el-input-number v-model="formData.absence"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="费用科目:"  prop="expense" >
              <el-input v-model="formData.expense" :clearable="true"  placeholder="请输入费用科目" />
            </el-form-item>
            <el-form-item label="职位:"  prop="position" >
              <el-input v-model="formData.position" :clearable="true"  placeholder="请输入职位" />
            </el-form-item>
            <el-form-item label="学历:"  prop="education" >
              <el-input v-model="formData.education" :clearable="true"  placeholder="请输入学历" />
            </el-form-item>
            <el-form-item label="等级工资:"  prop="gradeSalary" >
              <el-input-number v-model="formData.gradeSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="职称技能津贴:"  prop="skillSalary" >
              <el-input-number v-model="formData.skillSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="津贴:"  prop="allowancesSalary" >
              <el-input-number v-model="formData.allowancesSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="司龄津贴:"  prop="ageAllowanceSalary" >
              <el-input-number v-model="formData.ageAllowanceSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="学历工资:"  prop="educationSalary" >
              <el-input-number v-model="formData.educationSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="全勤:"  prop="fullSalary" >
              <el-input-number v-model="formData.fullSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="尽孝基金:"  prop="filialSalary" >
              <el-input-number v-model="formData.filialSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="业绩奖:"  prop="performanceSalary" >
              <el-input-number v-model="formData.performanceSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="节日金:"  prop="festivalSalary" >
              <el-input-number v-model="formData.festivalSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="生日金:"  prop="birthdaySalary" >
              <el-input-number v-model="formData.birthdaySalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="岗位补助:"  prop="jobSubsidy" >
              <el-input-number v-model="formData.jobSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="组长补助:"  prop="leaderSubsidy" >
              <el-input-number v-model="formData.leaderSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="电话补助:"  prop="telephoneSubsidy" >
              <el-input-number v-model="formData.telephoneSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="宿舍补助:"  prop="dormitorySubsidy" >
              <el-input-number v-model="formData.dormitorySubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="交通补助:"  prop="transportSubsidy" >
              <el-input-number v-model="formData.transportSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="餐补:"  prop="mealSubsidy" >
              <el-input-number v-model="formData.mealSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="出差补助:"  prop="travelSubsidy" >
              <el-input-number v-model="formData.travelSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="分包装间:"  prop="packagingSubsidy" >
              <el-input-number v-model="formData.packagingSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="宿舍费:"  prop="dormSubsidy" >
              <el-input-number v-model="formData.dormSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="其他(补助):"  prop="otherSubsidy" >
              <el-input-number v-model="formData.otherSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="合计(补助):"  prop="totalSubsidy" >
              <el-input-number v-model="formData.totalSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="加班费:"  prop="overtimeFee" >
              <el-input-number v-model="formData.overtimeFee"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="缺勤:"  prop="absenceFee" >
              <el-input-number v-model="formData.absenceFee"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="合计(加班及缺勤):"  prop="totalFee" >
              <el-input-number v-model="formData.totalFee"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="应发工资:"  prop="shouldPay" >
              <el-input-number v-model="formData.shouldPay"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="养老保险:"  prop="pensionPay" >
              <el-input-number v-model="formData.pensionPay"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="医疗保险:"  prop="medicalPay" >
              <el-input-number v-model="formData.medicalPay"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="失业保险:"  prop="unemploymentPay" >
              <el-input-number v-model="formData.unemploymentPay"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="住房公积金:"  prop="housingPay" >
              <el-input-number v-model="formData.housingPay"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="预存尽孝基金:"  prop="preFilialSalary" >
              <el-input-number v-model="formData.preFilialSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="其他:"  prop="other" >
              <el-input-number v-model="formData.other"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="已发:"  prop="sent" >
              <el-input-number v-model="formData.sent"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="实发小计:"  prop="actualPay" >
              <el-input-number v-model="formData.actualPay"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="员工ID">
                        {{ formData.staffId }}
                </el-descriptions-item>
                <el-descriptions-item label="姓名">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="入职时间">
                      {{ formatDate(formData.employmentDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="截止日期">
                      {{ formatDate(formData.deadline) }}
                </el-descriptions-item>
                <el-descriptions-item label="司龄">
                        {{ formData.companyAge }}
                </el-descriptions-item>
                <el-descriptions-item label="发薪天数">
                        {{ formData.salaryDays }}
                </el-descriptions-item>
                <el-descriptions-item label="加班情况">
                        {{ formData.overtime }}
                </el-descriptions-item>
                <el-descriptions-item label="缺勤情况">
                        {{ formData.absence }}
                </el-descriptions-item>
                <el-descriptions-item label="费用科目">
                        {{ formData.expense }}
                </el-descriptions-item>
                <el-descriptions-item label="职位">
                        {{ formData.position }}
                </el-descriptions-item>
                <el-descriptions-item label="学历">
                        {{ formData.education }}
                </el-descriptions-item>
                <el-descriptions-item label="等级工资">
                        {{ formData.gradeSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="职称技能津贴">
                        {{ formData.skillSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="津贴">
                        {{ formData.allowancesSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="司龄津贴">
                        {{ formData.ageAllowanceSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="学历工资">
                        {{ formData.educationSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="全勤">
                        {{ formData.fullSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="尽孝基金">
                        {{ formData.filialSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="业绩奖">
                        {{ formData.performanceSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="节日金">
                        {{ formData.festivalSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="生日金">
                        {{ formData.birthdaySalary }}
                </el-descriptions-item>
                <el-descriptions-item label="岗位补助">
                        {{ formData.jobSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="组长补助">
                        {{ formData.leaderSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="电话补助">
                        {{ formData.telephoneSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="宿舍补助">
                        {{ formData.dormitorySubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="交通补助">
                        {{ formData.transportSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="餐补">
                        {{ formData.mealSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="出差补助">
                        {{ formData.travelSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="分包装间">
                        {{ formData.packagingSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="宿舍费">
                        {{ formData.dormSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="其他(补助)">
                        {{ formData.otherSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="合计(补助)">
                        {{ formData.totalSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="加班费">
                        {{ formData.overtimeFee }}
                </el-descriptions-item>
                <el-descriptions-item label="缺勤">
                        {{ formData.absenceFee }}
                </el-descriptions-item>
                <el-descriptions-item label="合计(加班及缺勤)">
                        {{ formData.totalFee }}
                </el-descriptions-item>
                <el-descriptions-item label="应发工资">
                        {{ formData.shouldPay }}
                </el-descriptions-item>
                <el-descriptions-item label="养老保险">
                        {{ formData.pensionPay }}
                </el-descriptions-item>
                <el-descriptions-item label="医疗保险">
                        {{ formData.medicalPay }}
                </el-descriptions-item>
                <el-descriptions-item label="失业保险">
                        {{ formData.unemploymentPay }}
                </el-descriptions-item>
                <el-descriptions-item label="住房公积金">
                        {{ formData.housingPay }}
                </el-descriptions-item>
                <el-descriptions-item label="预存尽孝基金">
                        {{ formData.preFilialSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="其他">
                        {{ formData.other }}
                </el-descriptions-item>
                <el-descriptions-item label="已发">
                        {{ formData.sent }}
                </el-descriptions-item>
                <el-descriptions-item label="实发小计">
                        {{ formData.actualPay }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaffSalary,
  deleteWcStaffSalary,
  deleteWcStaffSalaryByIds,
  updateWcStaffSalary,
  findWcStaffSalary,
  getWcStaffSalaryList
} from '@/api/weChat/wcStaffSalary'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WcStaffSalary'
})

// 自动化生成的字典（可能为空）以及字段
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
               },
              ],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               employmentDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               deadline : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               companyAge : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               salaryDays : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               overtime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               absence : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               expense : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               position : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               education : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               gradeSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               skillSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               allowancesSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               ageAllowanceSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               educationSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               fullSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               filialSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               performanceSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               festivalSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               birthdaySalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               jobSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               leaderSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               telephoneSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               dormitorySubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               transportSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               mealSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               travelSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               packagingSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               dormSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               otherSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               totalSubsidy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               overtimeFee : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               absenceFee : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               totalFee : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               shouldPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               pensionPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               medicalPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               unemploymentPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               housingPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               preFilialSalary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               other : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               sent : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               actualPay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWcStaffSalaryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteWcStaffSalaryFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteWcStaffSalaryByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWcStaffSalaryFunc = async(row) => {
    const res = await findWcStaffSalary({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffSalary
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffSalaryFunc = async (row) => {
    const res = await deleteWcStaffSalary({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWcStaffSalary({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffSalary
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
