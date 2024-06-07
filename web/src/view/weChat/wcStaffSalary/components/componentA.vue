<template>
	<div>
		<div class="gva-search-box">
			<el-form
				ref="elSearchFormRef"
				:inline="true"
				:model="searchInfo"
				class="demo-form-inline"
				:rules="searchRule"
				@keyup.enter="onSubmit"
			>
				<el-form-item label="选择员工:" prop="staffId">
					<SelectStaff v-model="searchInfo.staffId" :disabled="type === 'update' ? 'disabled' : false"> </SelectStaff>
				</el-form-item>

				<el-form-item label="关键词">
					<el-input style="width: 190px" v-model="searchInfo.keyword" placeholder="请输入部门名称" />
				</el-form-item>

				<el-form-item label="工资月份">
					<el-date-picker v-model="searchInfo.month" type="month" value-format="YYYYMM" />
				</el-form-item>

				<el-form-item>
					<el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
					<el-button icon="refresh" @click="onReset">重置</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button icon="delete" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="onDelete"
					>删除</el-button
				>
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
				<el-table-column align="left" label="一级部门" prop="departmentFirst" width="120" />
				<el-table-column align="left" label="二级部门" prop="departmentSecond" width="120" />
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
				<el-table-column align="left" label="岗位工资/职称技能津贴" prop="skillSalary" width="170" />
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
				<el-table-column align="left" label="合计(加班及缺勤)" prop="totalFee" width="140" />
				<el-table-column align="left" label="应发工资" prop="shouldPay" width="120" />
				<el-table-column align="left" label="养老保险" prop="pensionPay" width="120" />
				<el-table-column align="left" label="医疗保险" prop="medicalPay" width="120" />
				<el-table-column align="left" label="失业保险" prop="unemploymentPay" width="120" />
				<el-table-column align="left" label="住房公积金" prop="housingPay" width="120" />
				<el-table-column align="left" label="专项抵扣" prop="specialTax" width="120" />
				<el-table-column align="left" label="本次应缴税所得额" prop="thisPayTax" width="140" />
				<el-table-column align="left" label="上次累计应缴预缴所得额" prop="lastTotalPayTax" width="180" />
				<el-table-column align="left" label="本次累计应缴预缴所得额" prop="thisTotalPayTax" width="180" />
				<el-table-column align="left" label="累计税额" prop="totalPayTax" width="120" />
				<el-table-column align="left" label="上次累计已扣个税" prop="lastTotalPaidTax" width="140" />
				<el-table-column align="left" label="本次累计已缴个税" prop="thisTotalPaidTax" width="140" />
				<el-table-column align="left" label="本次应扣个税" prop="thisShouldPayTax" width="120" />
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
						<!--        <el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffSalaryFunc(scope.row)">变更</el-button>-->
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
				<el-descriptions-item label="工资类型">
					{{ formData.typeText }}
				</el-descriptions-item>
				<el-descriptions-item label="一级部门">
					{{ formData.departmentFirst }}
				</el-descriptions-item>
				<el-descriptions-item label="二级部门">
					{{ formData.departmentSecond }}
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
	getWcStaffSalaryList,
} from '@/api/weChat/wcStaffSalary'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watchEffect, nextTick } from 'vue'
import { InfoFilled, QuestionFilled } from '@element-plus/icons-vue'
import SelectStaff from '@/components/selectStaff/index.vue'

defineOptions({
	name: 'WcStaffSalary',
})

const props = defineProps({
	month: {
		type: String,
		default: '',
	},
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
	staffId: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	name: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
		{
			whitespace: true,
			message: '不能只输入空格',
			trigger: ['input', 'blur'],
		},
	],
})

const searchRule = reactive({
	createdAt: [
		{
			validator: (rule, value, callback) => {
				if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
					callback(new Error('请填写结束日期'))
				} else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
					callback(new Error('请填写开始日期'))
				} else if (
					searchInfo.value.startCreatedAt &&
					searchInfo.value.endCreatedAt &&
					(searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() ||
						searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())
				) {
					callback(new Error('开始日期应当早于结束日期'))
				} else {
					callback()
				}
			},
			trigger: 'change',
		},
	],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({ type: 1, month: '' })

// 重置
const onReset = () => {
	searchInfo.value = { type: 1 }
	getTableData()
}

// 搜索
const onSubmit = () => {
	elSearchFormRef.value?.validate(async (valid) => {
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
const getTableData = async () => {
	const table = await getWcStaffSalaryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
	if (table.code === 0) {
		tableData.value = table.data.list
		total.value = table.data.total
		page.value = table.data.page
		pageSize.value = table.data.pageSize
	}
}

if (props.month) {
	searchInfo.value.month = props.month
	getTableData()
} else {
	getTableData()
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {}

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
		type: 'warning',
	}).then(() => {
		deleteWcStaffSalaryFunc(row)
	})
}

// 多选删除
const onDelete = async () => {
	ElMessageBox.confirm('确定要删除吗?', '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	}).then(async () => {
		const IDs = []
		if (multipleSelection.value.length === 0) {
			ElMessage({
				type: 'warning',
				message: '请选择要删除的数据',
			})
			return
		}
		multipleSelection.value &&
			multipleSelection.value.map((item) => {
				IDs.push(item.ID)
			})
		const res = await deleteWcStaffSalaryByIds({ IDs })
		if (res.code === 0) {
			ElMessage({
				type: 'success',
				message: '删除成功',
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
const updateWcStaffSalaryFunc = async (row) => {
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
			message: '删除成功',
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
	elFormRef.value?.validate(async (valid) => {
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
				message: '创建/更改成功',
			})
			closeDialog()
			getTableData()
		}
	})
}
</script>

<style></style>
