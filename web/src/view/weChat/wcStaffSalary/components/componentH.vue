<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <el-form-item label="选择员工:" prop="staffId">
          <SelectStaff v-model="searchInfo.staffId" :disabled="type === 'update' ? 'disabled' : false"> </SelectStaff>
        </el-form-item>

        <el-form-item label="关键词">
          <el-input style="width: 190px" v-model="searchInfo.keyword" placeholder="请输入部门名称" />
        </el-form-item>

        <el-form-item label="工资月份">
          <el-date-picker
              v-model="searchInfo.month"
              type="month"
              value-format="YYYYMM"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
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
        <el-table-column align="left" label="成员名称" prop="staffName" width="100" />
        <el-table-column align="left" label="员工工号" prop="jobNum" width="120" />
        <el-table-column align="left" label="工资月份" prop="month" width="120" />
        <el-table-column align="left" label="项目数" prop="projectNumber" width="120" />
        <el-table-column align="left" label="手工费" prop="handicraftCosts" width="120" />
        <el-table-column align="left" label="业绩" prop="performance" width="120" />
        <el-table-column align="left" label="提成" prop="commission" width="120" />
        <el-table-column align="left" label="应发合计" prop="shouldPay" width="120" />
        <el-table-column align="left" label="个税" prop="thisShouldPayTax" width="120" />
        <el-table-column align="left" label="实发合计" prop="actualPay" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
<!--            <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffSalaryFunc(scope.row)">变更</el-button>-->
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

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <template #title>
        <div class="flex justify-between items-center">
          <span class="text-lg">查看详情</span>
        </div>
      </template>
      <el-descriptions :column="1" border>
        <el-descriptions-item label="成员名称">
          {{ formData.staffName }}
        </el-descriptions-item>
        <el-descriptions-item label="员工工号">
          {{ formData.jobNum }}
        </el-descriptions-item>
        <el-descriptions-item label="工资月份">
          {{ formData.month }}
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
        <el-descriptions-item label="应发合计">
          {{ formData.shouldPay }}
        </el-descriptions-item>
        <el-descriptions-item label="个税">
          {{ formData.thisShouldPayTax }}
        </el-descriptions-item>
        <el-descriptions-item label="实发合计">
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
import {InfoFilled, QuestionFilled} from "@element-plus/icons-vue";
import SelectStaff from "@/components/selectStaff/index.vue";

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
  staffName: '',
  jobNum: '',
  typeText: '',
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
const searchInfo = ref({type:8})

// 重置
const onReset = () => {
  searchInfo.value = {type:8}
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
    staffName: '',
    jobNum: '',
    typeText: '',
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
    staffName: '',
    jobNum: '',
    typeText: '',
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