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
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="入职标题" prop="title" width="120" />
        <el-table-column sortable align="left" label="姓名" prop="staffName" width="120" />
         <el-table-column sortable align="left" label="入职日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.employmentDate) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="入职部门" prop="employmentDepartment" width="120" />
        <el-table-column sortable align="left" label="入职职位" prop="jobPosition" width="120" />
        <el-table-column align="left" label="社保电脑号" prop="socialNumber" width="120" />
        <el-table-column align="left" label="公积金账号" prop="accountNumber" width="120" />
        <el-table-column sortable align="left" label="户籍类型" prop="householdType" width="120" />
        <el-table-column sortable align="left" label="社保公积金缴纳地" prop="paymentPlace" width="120" />
        <el-table-column sortable align="left" label="性别(0未知1男2女)" prop="gender" width="120" />
         <el-table-column sortable align="left" label="出生年月" width="180">
            <template #default="scope">{{ formatDate(scope.row.birthday) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="籍贯" prop="nativePlace" width="120" />
        <el-table-column sortable align="left" label="民族" prop="nationality" width="120" />
        <el-table-column align="left" label="身高" prop="height" width="120" />
        <el-table-column align="left" label="体重" prop="weight" width="120" />
        <el-table-column sortable align="left" label="婚否(1:已婚 2:未婚 3:其他)" prop="marriage" width="120" />
        <el-table-column sortable align="left" label="政治面貌(1:团员 2:党员 3:群众 0:其他)" prop="politicalOutlook" width="120" />
        <el-table-column sortable align="left" label="学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他)" prop="education" width="120" />
        <el-table-column sortable align="left" label="专业" prop="major" width="120" />
        <el-table-column align="left" label="毕业院校" prop="school" width="120" />
         <el-table-column align="left" label="毕业时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.graduationDate) }}</template>
         </el-table-column>
           <el-table-column label="职称技能证书" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.certificate" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
        <el-table-column align="left" label="身份证号码" prop="idNumber" width="120" />
        <el-table-column align="left" label="身份证地址" prop="idAddress" width="120" />
        <el-table-column align="left" label="银行账号" prop="cardNumber" width="120" />
        <el-table-column align="left" label="开户支行信息" prop="bank" width="120" />
        <el-table-column align="left" label="本人联系微信" prop="contantWechat" width="120" />
        <el-table-column align="left" label="手机" prop="mobile" width="120" />
        <el-table-column align="left" label="常住地址" prop="homeAddress" width="120" />
           <el-table-column label="证件资料附件上传" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.licenseAttment" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
        <el-table-column align="left" label="联系人关系(1:父母 2:配偶 3:子女 0:其他)" prop="relationShip" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.relationShip) }}</template>
        </el-table-column>
        <el-table-column align="left" label="紧急联系人姓名" prop="relationName" width="120" />
        <el-table-column align="left" label="紧急联系人电话" prop="relationMobile" width="120" />
        <el-table-column align="left" label="联系人常住地址" prop="relationAddress" width="120" />
        <el-table-column align="left" label="是否经过首席执行官" prop="isCeopass" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isCeopass) }}</template>
        </el-table-column>
        <el-table-column align="left" label="入职体检是否正常" prop="isBodychecknormal" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isBodychecknormal) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="职级" prop="jobLevel" width="120" />
        <el-table-column align="left" label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)" prop="tryPeriod" width="120" />
        <el-table-column align="left" label="用人部门意见" prop="employingOpinion" width="120" />
        <el-table-column align="left" label="人资部意见" prop="humanOpinion" width="120" />
        <el-table-column align="left" label="通知紧急程度" prop="urgencyNotification" width="120" />
        <el-table-column align="left" label="入职意见" prop="onboardingOpinion" width="120" />
        <el-table-column sortable align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffEmploymentApplicationFunc(scope.row)">变更</el-button>
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
            <el-form-item label="入职标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入入职标题" />
            </el-form-item>
            <el-form-item label="姓名:"  prop="staffName" >
              <el-input v-model="formData.staffName" :clearable="true"  placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="入职日期:"  prop="employmentDate" >
              <el-date-picker v-model="formData.employmentDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="入职部门:"  prop="employmentDepartment" >
                <el-select v-model="formData.employmentDepartment" placeholder="请选择入职部门" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [10]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="入职职位:"  prop="jobPosition" >
                <el-select v-model="formData.jobPosition" placeholder="请选择入职职位" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="社保电脑号:"  prop="socialNumber" >
              <el-input v-model="formData.socialNumber" :clearable="true"  placeholder="请输入社保电脑号" />
            </el-form-item>
            <el-form-item label="公积金账号:"  prop="accountNumber" >
              <el-input v-model="formData.accountNumber" :clearable="true"  placeholder="请输入公积金账号" />
            </el-form-item>
            <el-form-item label="户籍类型:"  prop="householdType" >
                <el-select v-model="formData.householdType" placeholder="请选择户籍类型" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="社保公积金缴纳地:"  prop="paymentPlace" >
                <el-select v-model="formData.paymentPlace" placeholder="请选择社保公积金缴纳地" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [255]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="性别(0未知1男2女):"  prop="gender" >
                <el-select v-model="formData.gender" placeholder="请选择性别(0未知1男2女)" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="出生年月:"  prop="birthday" >
              <el-date-picker v-model="formData.birthday" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="籍贯:"  prop="nativePlace" >
                <el-select v-model="formData.nativePlace" placeholder="请选择籍贯" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [255]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="民族:"  prop="nationality" >
                <el-select v-model="formData.nationality" placeholder="请选择民族" style="width:100%" :clearable="true" >
                   <el-option v-for="item in [50]" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="身高:"  prop="height" >
              <el-input-number v-model="formData.height"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="体重:"  prop="weight" >
              <el-input-number v-model="formData.weight"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="婚否(1:已婚 2:未婚 3:其他):"  prop="marriage" >
                <el-select v-model="formData.marriage" placeholder="请选择婚否(1:已婚 2:未婚 3:其他)" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="政治面貌(1:团员 2:党员 3:群众 0:其他):"  prop="politicalOutlook" >
                <el-select v-model="formData.politicalOutlook" placeholder="请选择政治面貌(1:团员 2:党员 3:群众 0:其他)" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他):"  prop="education" >
                <el-select v-model="formData.education" placeholder="请选择学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他)" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="专业:"  prop="major" >
              <el-input v-model="formData.major" :clearable="true"  placeholder="请输入专业" />
            </el-form-item>
            <el-form-item label="毕业院校:"  prop="school" >
              <el-input v-model="formData.school" :clearable="true"  placeholder="请输入毕业院校" />
            </el-form-item>
            <el-form-item label="毕业时间:"  prop="graduationDate" >
              <el-date-picker v-model="formData.graduationDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="职称技能证书:"  prop="certificate" >
                <SelectImage
                 multiple
                 v-model="formData.certificate"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="身份证号码:"  prop="idNumber" >
              <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入身份证号码" />
            </el-form-item>
            <el-form-item label="身份证地址:"  prop="idAddress" >
              <el-input v-model="formData.idAddress" :clearable="true"  placeholder="请输入身份证地址" />
            </el-form-item>
            <el-form-item label="银行账号:"  prop="cardNumber" >
              <el-input v-model="formData.cardNumber" :clearable="true"  placeholder="请输入银行账号" />
            </el-form-item>
            <el-form-item label="开户支行信息:"  prop="bank" >
              <el-input v-model="formData.bank" :clearable="true"  placeholder="请输入开户支行信息" />
            </el-form-item>
            <el-form-item label="本人联系微信:"  prop="contantWechat" >
              <el-input v-model="formData.contantWechat" :clearable="true"  placeholder="请输入本人联系微信" />
            </el-form-item>
            <el-form-item label="手机:"  prop="mobile" >
              <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入手机" />
            </el-form-item>
            <el-form-item label="常住地址:"  prop="homeAddress" >
              <el-input v-model="formData.homeAddress" :clearable="true"  placeholder="请输入常住地址" />
            </el-form-item>
            <el-form-item label="证件资料附件上传:"  prop="licenseAttment" >
                <SelectImage
                 multiple
                 v-model="formData.licenseAttment"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="联系人关系(1:父母 2:配偶 3:子女 0:其他):"  prop="relationShip" >
              <el-switch v-model="formData.relationShip" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="紧急联系人姓名:"  prop="relationName" >
              <el-input v-model="formData.relationName" :clearable="true"  placeholder="请输入紧急联系人姓名" />
            </el-form-item>
            <el-form-item label="紧急联系人电话:"  prop="relationMobile" >
              <el-input v-model="formData.relationMobile" :clearable="true"  placeholder="请输入紧急联系人电话" />
            </el-form-item>
            <el-form-item label="联系人常住地址:"  prop="relationAddress" >
              <el-input v-model="formData.relationAddress" :clearable="true"  placeholder="请输入联系人常住地址" />
            </el-form-item>
            <el-form-item label="是否经过首席执行官:"  prop="isCeopass" >
              <el-switch v-model="formData.isCeopass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="入职体检是否正常:"  prop="isBodychecknormal" >
              <el-switch v-model="formData.isBodychecknormal" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="职级:"  prop="jobLevel" >
                <el-select v-model="formData.jobLevel" placeholder="请选择职级" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他):"  prop="tryPeriod" >
                <el-select v-model="formData.tryPeriod" placeholder="请选择试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="用人部门意见:"  prop="employingOpinion" >
              <el-input v-model="formData.employingOpinion" :clearable="true"  placeholder="请输入用人部门意见" />
            </el-form-item>
            <el-form-item label="人资部意见:"  prop="humanOpinion" >
              <el-input v-model="formData.humanOpinion" :clearable="true"  placeholder="请输入人资部意见" />
            </el-form-item>
            <el-form-item label="通知紧急程度:"  prop="urgencyNotification" >
              <el-input v-model="formData.urgencyNotification" :clearable="true"  placeholder="请输入通知紧急程度" />
            </el-form-item>
            <el-form-item label="入职意见:"  prop="onboardingOpinion" >
              <el-input v-model="formData.onboardingOpinion" :clearable="true"  placeholder="请输入入职意见" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
                <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true" >
                   <el-option v-for="item in []" :key="item" :label="item" :value="item" />
                </el-select>
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
                <el-descriptions-item label="入职标题">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="姓名">
                        {{ formData.staffName }}
                </el-descriptions-item>
                <el-descriptions-item label="入职日期">
                      {{ formatDate(formData.employmentDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="入职部门">
                        {{ formData.employmentDepartment }}
                </el-descriptions-item>
                <el-descriptions-item label="入职职位">
                        {{ formData.jobPosition }}
                </el-descriptions-item>
                <el-descriptions-item label="社保电脑号">
                        {{ formData.socialNumber }}
                </el-descriptions-item>
                <el-descriptions-item label="公积金账号">
                        {{ formData.accountNumber }}
                </el-descriptions-item>
                <el-descriptions-item label="户籍类型">
                        {{ formData.householdType }}
                </el-descriptions-item>
                <el-descriptions-item label="社保公积金缴纳地">
                        {{ formData.paymentPlace }}
                </el-descriptions-item>
                <el-descriptions-item label="性别(0未知1男2女)">
                        {{ formData.gender }}
                </el-descriptions-item>
                <el-descriptions-item label="出生年月">
                      {{ formatDate(formData.birthday) }}
                </el-descriptions-item>
                <el-descriptions-item label="籍贯">
                        {{ formData.nativePlace }}
                </el-descriptions-item>
                <el-descriptions-item label="民族">
                        {{ formData.nationality }}
                </el-descriptions-item>
                <el-descriptions-item label="身高">
                        {{ formData.height }}
                </el-descriptions-item>
                <el-descriptions-item label="体重">
                        {{ formData.weight }}
                </el-descriptions-item>
                <el-descriptions-item label="婚否(1:已婚 2:未婚 3:其他)">
                        {{ formData.marriage }}
                </el-descriptions-item>
                <el-descriptions-item label="政治面貌(1:团员 2:党员 3:群众 0:其他)">
                        {{ formData.politicalOutlook }}
                </el-descriptions-item>
                <el-descriptions-item label="学历(1:小学 2:初中 3:高中 4:中专 5:大专 6:本科 7:硕士 8:博士 0:其他)">
                        {{ formData.education }}
                </el-descriptions-item>
                <el-descriptions-item label="专业">
                        {{ formData.major }}
                </el-descriptions-item>
                <el-descriptions-item label="毕业院校">
                        {{ formData.school }}
                </el-descriptions-item>
                <el-descriptions-item label="毕业时间">
                      {{ formatDate(formData.graduationDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="职称技能证书">
                        <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="ReturnArrImg(formData.certificate)" :initial-index="index" v-for="(item,index) in formData.certificate" :key="index" :src="getUrl(item)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="身份证号码">
                        {{ formData.idNumber }}
                </el-descriptions-item>
                <el-descriptions-item label="身份证地址">
                        {{ formData.idAddress }}
                </el-descriptions-item>
                <el-descriptions-item label="银行账号">
                        {{ formData.cardNumber }}
                </el-descriptions-item>
                <el-descriptions-item label="开户支行信息">
                        {{ formData.bank }}
                </el-descriptions-item>
                <el-descriptions-item label="本人联系微信">
                        {{ formData.contantWechat }}
                </el-descriptions-item>
                <el-descriptions-item label="手机">
                        {{ formData.mobile }}
                </el-descriptions-item>
                <el-descriptions-item label="常住地址">
                        {{ formData.homeAddress }}
                </el-descriptions-item>
                <el-descriptions-item label="证件资料附件上传">
                        <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="ReturnArrImg(formData.licenseAttment)" :initial-index="index" v-for="(item,index) in formData.licenseAttment" :key="index" :src="getUrl(item)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="联系人关系(1:父母 2:配偶 3:子女 0:其他)">
                    {{ formatBoolean(formData.relationShip) }}
                </el-descriptions-item>
                <el-descriptions-item label="紧急联系人姓名">
                        {{ formData.relationName }}
                </el-descriptions-item>
                <el-descriptions-item label="紧急联系人电话">
                        {{ formData.relationMobile }}
                </el-descriptions-item>
                <el-descriptions-item label="联系人常住地址">
                        {{ formData.relationAddress }}
                </el-descriptions-item>
                <el-descriptions-item label="是否经过首席执行官">
                    {{ formatBoolean(formData.isCeopass) }}
                </el-descriptions-item>
                <el-descriptions-item label="入职体检是否正常">
                    {{ formatBoolean(formData.isBodychecknormal) }}
                </el-descriptions-item>
                <el-descriptions-item label="职级">
                        {{ formData.jobLevel }}
                </el-descriptions-item>
                <el-descriptions-item label="试用期(1:无试用期 2:1个月 3:2个月 4:3个月 5:4个月 6:5个月 7:6个月 0:其他)">
                        {{ formData.tryPeriod }}
                </el-descriptions-item>
                <el-descriptions-item label="用人部门意见">
                        {{ formData.employingOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="人资部意见">
                        {{ formData.humanOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="通知紧急程度">
                        {{ formData.urgencyNotification }}
                </el-descriptions-item>
                <el-descriptions-item label="入职意见">
                        {{ formData.onboardingOpinion }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                        {{ formData.status }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWcStaffEmploymentApplication,
  deleteWcStaffEmploymentApplication,
  deleteWcStaffEmploymentApplicationByIds,
  updateWcStaffEmploymentApplication,
  findWcStaffEmploymentApplication,
  getWcStaffEmploymentApplicationList
} from '@/api/employee/wcStaffEmploymentApplication'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WcStaffEmploymentApplication'
})

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               staffName : [{
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
               employmentDepartment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               jobPosition : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               socialNumber : [{
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
               accountNumber : [{
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
               householdType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               paymentPlace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               birthday : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               nativePlace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               nationality : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               marriage : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               politicalOutlook : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               education : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               major : [{
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
               school : [{
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
               graduationDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               certificate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               idNumber : [{
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
               idAddress : [{
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
               cardNumber : [{
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
               bank : [{
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
               mobile : [{
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
               homeAddress : [{
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
               licenseAttment : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               jobLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               tryPeriod : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            staffName: 'staff_name',
            employmentDate: 'employment_date',
            employmentDepartment: 'employment_department',
            jobPosition: 'job_position',
            householdType: 'household_type',
            paymentPlace: 'payment_place',
            gender: 'gender',
            birthday: 'birthday',
            nativePlace: 'native_place',
            nationality: 'nationality',
            marriage: 'marriage',
            politicalOutlook: 'political_outlook',
            education: 'education',
            major: 'major',
            jobLevel: 'job_level',
            status: 'status',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

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
    if (searchInfo.value.relationShip === ""){
        searchInfo.value.relationShip=null
    }
    if (searchInfo.value.isCeopass === ""){
        searchInfo.value.isCeopass=null
    }
    if (searchInfo.value.isBodychecknormal === ""){
        searchInfo.value.isBodychecknormal=null
    }
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
  const table = await getWcStaffEmploymentApplicationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWcStaffEmploymentApplicationFunc(row)
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
      const res = await deleteWcStaffEmploymentApplicationByIds({ IDs })
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
const updateWcStaffEmploymentApplicationFunc = async(row) => {
    const res = await findWcStaffEmploymentApplication({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewcStaffEmploymentApplication
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWcStaffEmploymentApplicationFunc = async (row) => {
    const res = await deleteWcStaffEmploymentApplication({ ID: row.ID })
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
  const res = await findWcStaffEmploymentApplication({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewcStaffEmploymentApplication
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          idNumber: '',
          idAddress: '',
          cardNumber: '',
          bank: '',
          contantWechat: '',
          mobile: '',
          homeAddress: '',
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
        idNumber: '',
        idAddress: '',
        cardNumber: '',
        bank: '',
        contantWechat: '',
        mobile: '',
        homeAddress: '',
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
