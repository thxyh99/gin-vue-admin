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
        
        <el-table-column align="left" label="工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金)" prop="type" width="120" />
        <el-table-column align="left" label="员工ID" prop="staffId" width="120" />
        <el-table-column align="left" label="一级部门" prop="departmentFirst" width="120" />
        <el-table-column align="left" label="二级部门" prop="departmentSecond" width="120" />
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
        <el-table-column align="left" label="职位/职级" prop="position" width="120" />
        <el-table-column align="left" label="学历" prop="education" width="120" />
        <el-table-column align="left" label="等级工资" prop="gradeSalary" width="120" />
        <el-table-column align="left" label="岗位工资/职称技能津贴" prop="skillSalary" width="120" />
        <el-table-column align="left" label="保密工资" prop="secrecySalary" width="120" />
        <el-table-column align="left" label="CPI工资" prop="cpiSalary" width="120" />
        <el-table-column align="left" label="津贴" prop="allowancesSalary" width="120" />
        <el-table-column align="left" label="司龄津贴" prop="ageAllowanceSalary" width="120" />
        <el-table-column align="left" label="学历工资" prop="educationSalary" width="120" />
        <el-table-column align="left" label="尽孝基金" prop="filialSalary" width="120" />
        <el-table-column align="left" label="业绩奖" prop="performanceSalary" width="120" />
        <el-table-column align="left" label="节日金" prop="festivalSalary" width="120" />
        <el-table-column align="left" label="生日金" prop="birthdaySalary" width="120" />
        <el-table-column align="left" label="高温补贴" prop="highTemperatureSubsidy" width="120" />
        <el-table-column align="left" label="全勤" prop="fullSalary" width="120" />
        <el-table-column align="left" label="组长补助" prop="leaderSubsidy" width="120" />
        <el-table-column align="left" label="分包装间" prop="packagingSubsidy" width="120" />
        <el-table-column align="left" label="岗位补助" prop="jobSubsidy" width="120" />
        <el-table-column align="left" label="宿舍补助" prop="dormitorySubsidy" width="120" />
        <el-table-column align="left" label="餐补" prop="mealSubsidy" width="120" />
        <el-table-column align="left" label="电话补助" prop="telephoneSubsidy" width="120" />
        <el-table-column align="left" label="交通补助" prop="transportSubsidy" width="120" />
        <el-table-column align="left" label="出差补助" prop="travelSubsidy" width="120" />
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
        <el-table-column align="left" label="专项抵扣" prop="specialTax" width="120" />
        <el-table-column align="left" label="本次应缴税所得额" prop="thisPayTax" width="120" />
        <el-table-column align="left" label="上次累计应缴预缴所得额" prop="lastTotalPayTax" width="120" />
        <el-table-column align="left" label="本次累计应缴预缴所得额" prop="thisTotalPayTax" width="120" />
        <el-table-column align="left" label="累计税额" prop="totalPayTax" width="120" />
        <el-table-column align="left" label="上次累计已扣个税" prop="lastTotalPaidTax" width="120" />
        <el-table-column align="left" label="本次累计已缴个税" prop="thisTotalPaidTax" width="120" />
        <el-table-column align="left" label="本次应扣个税" prop="thisShouldPayTax" width="120" />
        <el-table-column align="left" label="预存尽孝基金" prop="preFilialSalary" width="120" />
        <el-table-column align="left" label="其他" prop="other" width="120" />
        <el-table-column align="left" label="已发" prop="sent" width="120" />
        <el-table-column align="left" label="实发小计" prop="actualPay" width="120" />
        <el-table-column align="left" label="月度提成补发" prop="monthlyCommission" width="120" />
        <el-table-column align="left" label="年度提成" prop="annualCommission" width="120" />
        <el-table-column align="left" label="子品牌经营绩效奖励" prop="subBrandCommission" width="120" />
        <el-table-column align="left" label="奖金点" prop="bonusPoints" width="120" />
        <el-table-column align="left" label="出勤系数" prop="attendanceRatio" width="120" />
        <el-table-column align="left" label="月度评分" prop="monthlyRating" width="120" />
        <el-table-column align="left" label="个人计点" prop="personalCount" width="120" />
        <el-table-column align="left" label="迟到乐捐" prop="lateDonation" width="120" />
        <el-table-column align="left" label="月度奖金" prop="monthlyBonus" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="区域" prop="region" width="120" />
        <el-table-column align="left" label="全年任务（万）" prop="annualTask" width="120" />
        <el-table-column align="left" label="月度提成点" prop="monthlyCommissionRating" width="120" />
        <el-table-column align="left" label="提成比例" prop="commissionRatio" width="120" />
        <el-table-column align="left" label="所责市场" prop="market" width="120" />
        <el-table-column align="left" label="月度出货指标" prop="monthlyShipmentTarget" width="120" />
        <el-table-column align="left" label="当月实际出货" prop="actualShipment" width="120" />
        <el-table-column align="left" label="月度提成基数" prop="monthlyCommissionBase" width="120" />
        <el-table-column align="left" label="完成率%" prop="completionRate" width="120" />
        <el-table-column align="left" label="绩效考评分数" prop="performanceScore" width="120" />
        <el-table-column align="left" label="不在岗缺勤" prop="absenceWork" width="120" />
        <el-table-column align="left" label="出货提成" prop="shippingCommission" width="120" />
        <el-table-column align="left" label="个人销售提成" prop="personalSalesCommission" width="120" />
        <el-table-column align="left" label="市场差补服务费" prop="marketServiceFee" width="120" />
        <el-table-column align="left" label="月度奖励" prop="monthlyReward" width="120" />
        <el-table-column align="left" label="品牌负责人月度补贴" prop="monthlyLeaderSubsidy" width="120" />
        <el-table-column align="left" label="大店开发提成" prop="storeCommission" width="120" />
        <el-table-column align="left" label="项目数" prop="projectNumber" width="120" />
        <el-table-column align="left" label="手工费" prop="handicraftCosts" width="120" />
        <el-table-column align="left" label="业绩" prop="performance" width="120" />
        <el-table-column align="left" label="提成" prop="commission" width="120" />
        <el-table-column align="left" label="备用字段1" prop="stringValue1" width="120" />
        <el-table-column align="left" label="备用字段2" prop="stringValue2" width="120" />
        <el-table-column align="left" label="备用字段3" prop="stringValue3" width="120" />
        <el-table-column align="left" label="备用字段4" prop="stringValue4" width="120" />
        <el-table-column align="left" label="备用字段5" prop="stringValue5" width="120" />
        <el-table-column align="left" label="备用字段1" prop="decimalValue1" width="120" />
        <el-table-column align="left" label="备用字段2" prop="decimalValue2" width="120" />
        <el-table-column align="left" label="备用字段3" prop="decimalValue3" width="120" />
        <el-table-column align="left" label="备用字段4" prop="decimalValue4" width="120" />
        <el-table-column align="left" label="备用字段5" prop="decimalValue5" width="120" />
        <el-table-column align="left" label="备用字段1" prop="intValue1" width="120" />
        <el-table-column align="left" label="备用字段2" prop="intValue2" width="120" />
        <el-table-column align="left" label="备用字段3" prop="intValue3" width="120" />
        <el-table-column align="left" label="备用字段4" prop="intValue4" width="120" />
        <el-table-column align="left" label="备用字段5" prop="intValue5" width="120" />
         <el-table-column align="left" label="备用字段1" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateValue1) }}</template>
         </el-table-column>
         <el-table-column align="left" label="备用字段2" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateValue2) }}</template>
         </el-table-column>
         <el-table-column align="left" label="备用字段3" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateValue3) }}</template>
         </el-table-column>
         <el-table-column align="left" label="备用字段4" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateValue4) }}</template>
         </el-table-column>
         <el-table-column align="left" label="备用字段5" width="180">
            <template #default="scope">{{ formatDate(scope.row.dateValue5) }}</template>
         </el-table-column>
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
            <el-form-item label="工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金):"  prop="type" >
              <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金)" />
            </el-form-item>
            <el-form-item label="员工ID:"  prop="staffId" >
              <el-input v-model.number="formData.staffId" :clearable="true" placeholder="请输入员工ID" />
            </el-form-item>
            <el-form-item label="一级部门:"  prop="departmentFirst" >
              <el-input v-model="formData.departmentFirst" :clearable="true"  placeholder="请输入一级部门" />
            </el-form-item>
            <el-form-item label="二级部门:"  prop="departmentSecond" >
              <el-input v-model="formData.departmentSecond" :clearable="true"  placeholder="请输入二级部门" />
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
            <el-form-item label="职位/职级:"  prop="position" >
              <el-input v-model="formData.position" :clearable="true"  placeholder="请输入职位/职级" />
            </el-form-item>
            <el-form-item label="学历:"  prop="education" >
              <el-input v-model="formData.education" :clearable="true"  placeholder="请输入学历" />
            </el-form-item>
            <el-form-item label="等级工资:"  prop="gradeSalary" >
              <el-input-number v-model="formData.gradeSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="岗位工资/职称技能津贴:"  prop="skillSalary" >
              <el-input-number v-model="formData.skillSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="保密工资:"  prop="secrecySalary" >
              <el-input-number v-model="formData.secrecySalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="CPI工资:"  prop="cpiSalary" >
              <el-input-number v-model="formData.cpiSalary"  style="width:100%" :precision="2" :clearable="true"  />
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
            <el-form-item label="高温补贴:"  prop="highTemperatureSubsidy" >
              <el-input-number v-model="formData.highTemperatureSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="全勤:"  prop="fullSalary" >
              <el-input-number v-model="formData.fullSalary"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="组长补助:"  prop="leaderSubsidy" >
              <el-input-number v-model="formData.leaderSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="分包装间:"  prop="packagingSubsidy" >
              <el-input-number v-model="formData.packagingSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="岗位补助:"  prop="jobSubsidy" >
              <el-input-number v-model="formData.jobSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="宿舍补助:"  prop="dormitorySubsidy" >
              <el-input-number v-model="formData.dormitorySubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="餐补:"  prop="mealSubsidy" >
              <el-input-number v-model="formData.mealSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="电话补助:"  prop="telephoneSubsidy" >
              <el-input-number v-model="formData.telephoneSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="交通补助:"  prop="transportSubsidy" >
              <el-input-number v-model="formData.transportSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="出差补助:"  prop="travelSubsidy" >
              <el-input-number v-model="formData.travelSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
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
            <el-form-item label="专项抵扣:"  prop="specialTax" >
              <el-input-number v-model="formData.specialTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="本次应缴税所得额:"  prop="thisPayTax" >
              <el-input-number v-model="formData.thisPayTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="上次累计应缴预缴所得额:"  prop="lastTotalPayTax" >
              <el-input-number v-model="formData.lastTotalPayTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="本次累计应缴预缴所得额:"  prop="thisTotalPayTax" >
              <el-input-number v-model="formData.thisTotalPayTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="累计税额:"  prop="totalPayTax" >
              <el-input-number v-model="formData.totalPayTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="上次累计已扣个税:"  prop="lastTotalPaidTax" >
              <el-input-number v-model="formData.lastTotalPaidTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="本次累计已缴个税:"  prop="thisTotalPaidTax" >
              <el-input-number v-model="formData.thisTotalPaidTax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="本次应扣个税:"  prop="thisShouldPayTax" >
              <el-input-number v-model="formData.thisShouldPayTax"  style="width:100%" :precision="2" :clearable="true"  />
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
            <el-form-item label="月度提成补发:"  prop="monthlyCommission" >
              <el-input-number v-model="formData.monthlyCommission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="年度提成:"  prop="annualCommission" >
              <el-input-number v-model="formData.annualCommission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="子品牌经营绩效奖励:"  prop="subBrandCommission" >
              <el-input-number v-model="formData.subBrandCommission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="奖金点:"  prop="bonusPoints" >
              <el-input-number v-model="formData.bonusPoints"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="出勤系数:"  prop="attendanceRatio" >
              <el-input-number v-model="formData.attendanceRatio"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="月度评分:"  prop="monthlyRating" >
              <el-input-number v-model="formData.monthlyRating"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="个人计点:"  prop="personalCount" >
              <el-input-number v-model="formData.personalCount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="迟到乐捐:"  prop="lateDonation" >
              <el-input-number v-model="formData.lateDonation"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="月度奖金:"  prop="monthlyBonus" >
              <el-input-number v-model="formData.monthlyBonus"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
            <el-form-item label="区域:"  prop="region" >
              <el-input v-model="formData.region" :clearable="true"  placeholder="请输入区域" />
            </el-form-item>
            <el-form-item label="全年任务（万）:"  prop="annualTask" >
              <el-input-number v-model="formData.annualTask"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="月度提成点:"  prop="monthlyCommissionRating" >
              <el-input-number v-model="formData.monthlyCommissionRating"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="提成比例:"  prop="commissionRatio" >
              <el-input-number v-model="formData.commissionRatio"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="所责市场:"  prop="market" >
              <el-input v-model="formData.market" :clearable="true"  placeholder="请输入所责市场" />
            </el-form-item>
            <el-form-item label="月度出货指标:"  prop="monthlyShipmentTarget" >
              <el-input-number v-model="formData.monthlyShipmentTarget"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="当月实际出货:"  prop="actualShipment" >
              <el-input-number v-model="formData.actualShipment"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="月度提成基数:"  prop="monthlyCommissionBase" >
              <el-input-number v-model="formData.monthlyCommissionBase"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="完成率%:"  prop="completionRate" >
              <el-input-number v-model="formData.completionRate"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="绩效考评分数:"  prop="performanceScore" >
              <el-input-number v-model="formData.performanceScore"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="不在岗缺勤:"  prop="absenceWork" >
              <el-input-number v-model="formData.absenceWork"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="出货提成:"  prop="shippingCommission" >
              <el-input-number v-model="formData.shippingCommission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="个人销售提成:"  prop="personalSalesCommission" >
              <el-input-number v-model="formData.personalSalesCommission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="市场差补服务费:"  prop="marketServiceFee" >
              <el-input-number v-model="formData.marketServiceFee"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="月度奖励:"  prop="monthlyReward" >
              <el-input-number v-model="formData.monthlyReward"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="品牌负责人月度补贴:"  prop="monthlyLeaderSubsidy" >
              <el-input-number v-model="formData.monthlyLeaderSubsidy"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="大店开发提成:"  prop="storeCommission" >
              <el-input-number v-model="formData.storeCommission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="项目数:"  prop="projectNumber" >
              <el-input v-model.number="formData.projectNumber" :clearable="true" placeholder="请输入项目数" />
            </el-form-item>
            <el-form-item label="手工费:"  prop="handicraftCosts" >
              <el-input-number v-model="formData.handicraftCosts"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="业绩:"  prop="performance" >
              <el-input-number v-model="formData.performance"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="提成:"  prop="commission" >
              <el-input-number v-model="formData.commission"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段1:"  prop="stringValue1" >
              <el-input v-model="formData.stringValue1" :clearable="true"  placeholder="请输入备用字段1" />
            </el-form-item>
            <el-form-item label="备用字段2:"  prop="stringValue2" >
              <el-input v-model="formData.stringValue2" :clearable="true"  placeholder="请输入备用字段2" />
            </el-form-item>
            <el-form-item label="备用字段3:"  prop="stringValue3" >
              <el-input v-model="formData.stringValue3" :clearable="true"  placeholder="请输入备用字段3" />
            </el-form-item>
            <el-form-item label="备用字段4:"  prop="stringValue4" >
              <el-input v-model="formData.stringValue4" :clearable="true"  placeholder="请输入备用字段4" />
            </el-form-item>
            <el-form-item label="备用字段5:"  prop="stringValue5" >
              <el-input v-model="formData.stringValue5" :clearable="true"  placeholder="请输入备用字段5" />
            </el-form-item>
            <el-form-item label="备用字段1:"  prop="decimalValue1" >
              <el-input-number v-model="formData.decimalValue1"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段2:"  prop="decimalValue2" >
              <el-input-number v-model="formData.decimalValue2"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段3:"  prop="decimalValue3" >
              <el-input-number v-model="formData.decimalValue3"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段4:"  prop="decimalValue4" >
              <el-input-number v-model="formData.decimalValue4"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段5:"  prop="decimalValue5" >
              <el-input-number v-model="formData.decimalValue5"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段1:"  prop="intValue1" >
              <el-input v-model.number="formData.intValue1" :clearable="true" placeholder="请输入备用字段1" />
            </el-form-item>
            <el-form-item label="备用字段2:"  prop="intValue2" >
              <el-input v-model.number="formData.intValue2" :clearable="true" placeholder="请输入备用字段2" />
            </el-form-item>
            <el-form-item label="备用字段3:"  prop="intValue3" >
              <el-input v-model.number="formData.intValue3" :clearable="true" placeholder="请输入备用字段3" />
            </el-form-item>
            <el-form-item label="备用字段4:"  prop="intValue4" >
              <el-input v-model.number="formData.intValue4" :clearable="true" placeholder="请输入备用字段4" />
            </el-form-item>
            <el-form-item label="备用字段5:"  prop="intValue5" >
              <el-input v-model.number="formData.intValue5" :clearable="true" placeholder="请输入备用字段5" />
            </el-form-item>
            <el-form-item label="备用字段1:"  prop="dateValue1" >
              <el-date-picker v-model="formData.dateValue1" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段2:"  prop="dateValue2" >
              <el-date-picker v-model="formData.dateValue2" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段3:"  prop="dateValue3" >
              <el-date-picker v-model="formData.dateValue3" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段4:"  prop="dateValue4" >
              <el-date-picker v-model="formData.dateValue4" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="备用字段5:"  prop="dateValue5" >
              <el-date-picker v-model="formData.dateValue5" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
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
                <el-descriptions-item label="工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金)">
                        {{ formData.type }}
                </el-descriptions-item>
                <el-descriptions-item label="员工ID">
                        {{ formData.staffId }}
                </el-descriptions-item>
                <el-descriptions-item label="一级部门">
                        {{ formData.departmentFirst }}
                </el-descriptions-item>
                <el-descriptions-item label="二级部门">
                        {{ formData.departmentSecond }}
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
                <el-descriptions-item label="职位/职级">
                        {{ formData.position }}
                </el-descriptions-item>
                <el-descriptions-item label="学历">
                        {{ formData.education }}
                </el-descriptions-item>
                <el-descriptions-item label="等级工资">
                        {{ formData.gradeSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="岗位工资/职称技能津贴">
                        {{ formData.skillSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="保密工资">
                        {{ formData.secrecySalary }}
                </el-descriptions-item>
                <el-descriptions-item label="CPI工资">
                        {{ formData.cpiSalary }}
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
                <el-descriptions-item label="高温补贴">
                        {{ formData.highTemperatureSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="全勤">
                        {{ formData.fullSalary }}
                </el-descriptions-item>
                <el-descriptions-item label="组长补助">
                        {{ formData.leaderSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="分包装间">
                        {{ formData.packagingSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="岗位补助">
                        {{ formData.jobSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="宿舍补助">
                        {{ formData.dormitorySubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="餐补">
                        {{ formData.mealSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="电话补助">
                        {{ formData.telephoneSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="交通补助">
                        {{ formData.transportSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="出差补助">
                        {{ formData.travelSubsidy }}
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
                <el-descriptions-item label="专项抵扣">
                        {{ formData.specialTax }}
                </el-descriptions-item>
                <el-descriptions-item label="本次应缴税所得额">
                        {{ formData.thisPayTax }}
                </el-descriptions-item>
                <el-descriptions-item label="上次累计应缴预缴所得额">
                        {{ formData.lastTotalPayTax }}
                </el-descriptions-item>
                <el-descriptions-item label="本次累计应缴预缴所得额">
                        {{ formData.thisTotalPayTax }}
                </el-descriptions-item>
                <el-descriptions-item label="累计税额">
                        {{ formData.totalPayTax }}
                </el-descriptions-item>
                <el-descriptions-item label="上次累计已扣个税">
                        {{ formData.lastTotalPaidTax }}
                </el-descriptions-item>
                <el-descriptions-item label="本次累计已缴个税">
                        {{ formData.thisTotalPaidTax }}
                </el-descriptions-item>
                <el-descriptions-item label="本次应扣个税">
                        {{ formData.thisShouldPayTax }}
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
                <el-descriptions-item label="月度提成补发">
                        {{ formData.monthlyCommission }}
                </el-descriptions-item>
                <el-descriptions-item label="年度提成">
                        {{ formData.annualCommission }}
                </el-descriptions-item>
                <el-descriptions-item label="子品牌经营绩效奖励">
                        {{ formData.subBrandCommission }}
                </el-descriptions-item>
                <el-descriptions-item label="奖金点">
                        {{ formData.bonusPoints }}
                </el-descriptions-item>
                <el-descriptions-item label="出勤系数">
                        {{ formData.attendanceRatio }}
                </el-descriptions-item>
                <el-descriptions-item label="月度评分">
                        {{ formData.monthlyRating }}
                </el-descriptions-item>
                <el-descriptions-item label="个人计点">
                        {{ formData.personalCount }}
                </el-descriptions-item>
                <el-descriptions-item label="迟到乐捐">
                        {{ formData.lateDonation }}
                </el-descriptions-item>
                <el-descriptions-item label="月度奖金">
                        {{ formData.monthlyBonus }}
                </el-descriptions-item>
                <el-descriptions-item label="备注">
                        {{ formData.remark }}
                </el-descriptions-item>
                <el-descriptions-item label="区域">
                        {{ formData.region }}
                </el-descriptions-item>
                <el-descriptions-item label="全年任务（万）">
                        {{ formData.annualTask }}
                </el-descriptions-item>
                <el-descriptions-item label="月度提成点">
                        {{ formData.monthlyCommissionRating }}
                </el-descriptions-item>
                <el-descriptions-item label="提成比例">
                        {{ formData.commissionRatio }}
                </el-descriptions-item>
                <el-descriptions-item label="所责市场">
                        {{ formData.market }}
                </el-descriptions-item>
                <el-descriptions-item label="月度出货指标">
                        {{ formData.monthlyShipmentTarget }}
                </el-descriptions-item>
                <el-descriptions-item label="当月实际出货">
                        {{ formData.actualShipment }}
                </el-descriptions-item>
                <el-descriptions-item label="月度提成基数">
                        {{ formData.monthlyCommissionBase }}
                </el-descriptions-item>
                <el-descriptions-item label="完成率%">
                        {{ formData.completionRate }}
                </el-descriptions-item>
                <el-descriptions-item label="绩效考评分数">
                        {{ formData.performanceScore }}
                </el-descriptions-item>
                <el-descriptions-item label="不在岗缺勤">
                        {{ formData.absenceWork }}
                </el-descriptions-item>
                <el-descriptions-item label="出货提成">
                        {{ formData.shippingCommission }}
                </el-descriptions-item>
                <el-descriptions-item label="个人销售提成">
                        {{ formData.personalSalesCommission }}
                </el-descriptions-item>
                <el-descriptions-item label="市场差补服务费">
                        {{ formData.marketServiceFee }}
                </el-descriptions-item>
                <el-descriptions-item label="月度奖励">
                        {{ formData.monthlyReward }}
                </el-descriptions-item>
                <el-descriptions-item label="品牌负责人月度补贴">
                        {{ formData.monthlyLeaderSubsidy }}
                </el-descriptions-item>
                <el-descriptions-item label="大店开发提成">
                        {{ formData.storeCommission }}
                </el-descriptions-item>
                <el-descriptions-item label="项目数">
                        {{ formData.projectNumber }}
                </el-descriptions-item>
                <el-descriptions-item label="手工费">
                        {{ formData.handicraftCosts }}
                </el-descriptions-item>
                <el-descriptions-item label="业绩">
                        {{ formData.performance }}
                </el-descriptions-item>
                <el-descriptions-item label="提成">
                        {{ formData.commission }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段1">
                        {{ formData.stringValue1 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段2">
                        {{ formData.stringValue2 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段3">
                        {{ formData.stringValue3 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段4">
                        {{ formData.stringValue4 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段5">
                        {{ formData.stringValue5 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段1">
                        {{ formData.decimalValue1 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段2">
                        {{ formData.decimalValue2 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段3">
                        {{ formData.decimalValue3 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段4">
                        {{ formData.decimalValue4 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段5">
                        {{ formData.decimalValue5 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段1">
                        {{ formData.intValue1 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段2">
                        {{ formData.intValue2 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段3">
                        {{ formData.intValue3 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段4">
                        {{ formData.intValue4 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段5">
                        {{ formData.intValue5 }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段1">
                      {{ formatDate(formData.dateValue1) }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段2">
                      {{ formatDate(formData.dateValue2) }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段3">
                      {{ formatDate(formData.dateValue3) }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段4">
                      {{ formatDate(formData.dateValue4) }}
                </el-descriptions-item>
                <el-descriptions-item label="备用字段5">
                      {{ formatDate(formData.dateValue5) }}
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
        subBrandCommission: 0,
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
          subBrandCommission: 0,
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
        subBrandCommission: 0,
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
