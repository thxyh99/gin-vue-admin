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
				<el-form-item label="关键词">
					<el-input style="width: 190px" v-model="searchInfo.keyword" placeholder="请输入证件号码、个人账号" />
				</el-form-item>
				<el-form-item label="选择员工:" prop="staffId">
					<SelectStaff v-model="searchInfo.staffId" :disabled="type === 'update' ? 'disabled' : false"> </SelectStaff>
				</el-form-item>
        <el-form-item label="缴费年月">
					<el-date-picker
            v-model="searchInfo.period"
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
<!--				<el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
        <ImportExcel template-id="staffSocialSzGjj" type="2" @on-success="getTableData" btnName="导入深圳公积金" />
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
        <el-table-column align="left" label="证件号码" prop="credentialNumber" width="200" />
        <el-table-column align="left" label="个人账号" prop="account" width="180" />
        <el-table-column align="left" label="缴存时段" prop="periodStart" width="120" />
				<el-table-column align="left" label="缴存基数（元）" prop="housingBase" width="150" />
        <el-table-column align="left" label="单位缴存比例" prop="housingRatioUnit" width="120" />
        <el-table-column align="left" label="个人缴存比例" prop="housingRatioSelf" width="120" />
        <el-table-column align="left" label="金额合计（元）" prop="totalHousing" width="150" />
				<el-table-column align="left" label="操作" fixed="right" min-width="200">
					<template #default="scope">
						<el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
							<el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
							查看详情
						</el-button>
<!--						<el-button type="primary" link icon="edit" class="table-button" @click="updateWcStaffSocialFunc(scope.row)"-->
<!--							>变更</el-button-->
<!--						>-->
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
					<span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
					<div>
						<el-button type="primary" @click="enterDialog">确 定</el-button>
						<el-button @click="closeDialog">取 消</el-button>
					</div>
				</div>
			</template>

			<el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
				<el-form-item label="选择员工:" prop="staffId">
					<SelectStaff v-model="formData.staffId" :disabled="type === 'update' ? 'disabled' : false"> </SelectStaff>
				</el-form-item>
				<el-form-item label="类型:" prop="type">
					<el-select v-model="formData.type" placeholder="选择类型" :disabled="type === 'update' ? 'disabled' : false">
						<el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="证件类型:" prop="type">
					<el-select
						v-model="formData.credentialType"
						placeholder="选择类型"
						:disabled="type === 'update' ? 'disabled' : false"
					>
						<el-option
							v-for="credentialType in credentialTypes"
							:key="credentialType.value"
							:label="credentialType.label"
							:value="credentialType.value"
						></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="证件号码:" prop="credentialNumber">
					<el-input
						v-model="formData.credentialNumber"
						:clearable="true"
						placeholder="请输入证件号码"
						:disabled="type === 'update' ? 'disabled' : false"
					/>
				</el-form-item>
				<el-form-item label="姓名:" prop="name">
					<el-input
						v-model="formData.name"
						:clearable="true"
						placeholder="请输入姓名"
						:disabled="type === 'update' ? 'disabled' : false"
					/>
				</el-form-item>
				<el-form-item label="账号:" prop="account">
					<el-input
						v-model="formData.account"
						:clearable="true"
						placeholder="请输入账号"
						:disabled="type === 'update' ? 'disabled' : false"
					/>
				</el-form-item>
				<el-form-item label="社保缴费合计:" prop="totalSocial">
					<el-input-number v-model="formData.totalSocial" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="社保缴费个人合计:" prop="totalSocialSelf">
					<el-input-number v-model="formData.totalSocialSelf" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="社保缴费单位合计:" prop="totalSocialUnit">
					<el-input-number v-model="formData.totalSocialUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="养老缴纳基数:" prop="pensionBase">
					<el-input-number v-model="formData.pensionBase" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="养老个人缴费:" prop="pensionSelf">
					<el-input-number v-model="formData.pensionSelf" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="养老单位缴费:" prop="pensionUnit">
					<el-input-number v-model="formData.pensionUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="医疗缴纳基数:" prop="medicalBase">
					<el-input-number v-model="formData.medicalBase" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="医疗个人缴费:" prop="medicalSelf">
					<el-input-number v-model="formData.medicalSelf" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="医疗单位缴费:" prop="medicalUnit">
					<el-input-number v-model="formData.medicalUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="失业缴纳基数:" prop="unemployedBase">
					<el-input-number v-model="formData.unemployedBase" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="失业个人缴费:" prop="unemployedSelf">
					<el-input-number v-model="formData.unemployedSelf" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="失业单位缴费:" prop="unemployedUnit">
					<el-input-number v-model="formData.unemployedUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="工商保险缴费基数:" prop="injuryInsuranceBase">
					<el-input-number
						v-model="formData.injuryInsuranceBase"
						style="width: 100%"
						:precision="2"
						:clearable="true"
					/>
				</el-form-item>
				<el-form-item label="工商保险单位缴费:" prop="injuryInsuranceUnit">
					<el-input-number
						v-model="formData.injuryInsuranceUnit"
						style="width: 100%"
						:precision="2"
						:clearable="true"
					/>
				</el-form-item>
				<el-form-item label="生育医疗缴费基数:" prop="birthBase">
					<el-input-number v-model="formData.birthBase" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="生育医疗单位缴费:" prop="birthUnit">
					<el-input-number v-model="formData.birthUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="公积金缴费合计:" prop="totalHousing">
					<el-input-number v-model="formData.totalHousing" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="公积金个人缴费合计:" prop="totalHousingSelf">
					<el-input-number v-model="formData.totalHousingSelf" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="公积金公司缴费合计:" prop="totalHousingUnit">
					<el-input-number v-model="formData.totalHousingUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="缴存基数:" prop="housingBase">
					<el-input-number v-model="formData.housingBase" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="个人缴存比例:" prop="housingRatioSelf">
					<el-input-number v-model="formData.housingRatioSelf" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="单位缴存比例:" prop="housingRatioUnit">
					<el-input-number v-model="formData.housingRatioUnit" style="width: 100%" :precision="2" :clearable="true" />
				</el-form-item>
				<el-form-item label="费款所属期起:" prop="periodStart">
					<el-input v-model="formData.periodStart" :clearable="true" placeholder="请输入费款所属期起" />
				</el-form-item>
				<el-form-item label="费款所属期止:" prop="periodEnd">
					<el-input v-model="formData.periodEnd" :clearable="true" placeholder="请输入费款所属期止" />
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
				<el-descriptions-item label="成员名称">
					{{ formData.staffName }}
				</el-descriptions-item>
				<el-descriptions-item label="员工工号">
					{{ formData.jobNum }}
        </el-descriptions-item>
        <el-descriptions-item label="证件号码">
          {{ formData.credentialNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="个人账号">
					{{ formData.account }}
        </el-descriptions-item>
        <el-descriptions-item label="缴存时段">
          {{ formData.periodStart }}
        </el-descriptions-item>
        <el-descriptions-item label="缴存基数（元）">
          {{ formData.housingBase }}
        </el-descriptions-item>
        <el-descriptions-item label="单位缴存比例">
          {{ formData.housingRatioUnit }}
        </el-descriptions-item>
        <el-descriptions-item label="个人缴存比例">
          {{ formData.housingRatioSelf }}
        </el-descriptions-item>
				<el-descriptions-item label="金额合计（元）">
					{{ formData.totalHousing }}
				</el-descriptions-item>
			</el-descriptions>
		</el-drawer>
	</div>
</template>

<script setup>
import {
	createWcStaffSocial,
	deleteWcStaffSocial,
	deleteWcStaffSocialByIds,
	updateWcStaffSocial,
	findWcStaffSocial,
	getWcStaffSocialList,
} from '@/api/weChat/wcStaffSocial'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { InfoFilled, QuestionFilled } from '@element-plus/icons-vue'
import SelectStaff from '@/components/selectStaff/index.vue'
import ImportExcel from '@/components/exportExcel/wechat/importSocialExcel.vue'

defineOptions({
	name: 'WcStaffSocial',
})

const types = ref([
	{ label: '深圳社保', value: 1 },
	{ label: '深圳公积金', value: 2 },
	{ label: '东莞社保', value: 3 },
	{ label: '东莞公积金', value: 4 },
])

const credentialTypes = ref([
	{ label: '其他', value: 0 },
	{ label: '居民身份证', value: 1 },
	{ label: '港澳通行证', value: 2 },
])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
	account: '',
	type: '',
	staffId: '',
	name: '',
	credentialType: '',
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
	staffName: '',
	jobNum: '',
	typeText: '',
	credentialTypeText: '',
})

// 验证规则
const rule = reactive({
	account: [
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
	type: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
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
	credentialType: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	credentialNumber: [
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
	totalSocial: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	totalSocialSelf: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	totalSocialUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	pensionBase: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	pensionSelf: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	pensionUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	medicalBase: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	medicalSelf: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	medicalUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	unemployedBase: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	unemployedSelf: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	unemployedUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	injuryInsuranceBase: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	injuryInsuranceUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	birthBase: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	birthUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	totalHousing: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	totalHousingSelf: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	totalHousingUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	housingBase: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	housingRatioSelf: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	housingRatioUnit: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	periodStart: [
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
	periodEnd: [
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
const searchInfo = ref({type:2})

// 重置
const onReset = () => {
	searchInfo.value = {type:2}
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
	const table = await getWcStaffSocialList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
		deleteWcStaffSocialFunc(row)
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
		const res = await deleteWcStaffSocialByIds({ IDs })
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
const updateWcStaffSocialFunc = async (row) => {
	const res = await findWcStaffSocial({ ID: row.ID })
	type.value = 'update'
	if (res.code === 0) {
		formData.value = res.data.rewcStaffSocial
		dialogFormVisible.value = true
	}
}

// 删除行
const deleteWcStaffSocialFunc = async (row) => {
	const res = await deleteWcStaffSocial({ ID: row.ID })
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
	const res = await findWcStaffSocial({ ID: row.ID })
	if (res.code === 0) {
		formData.value = res.data.rewcStaffSocial
		openDetailShow()
	}
}

// 关闭详情弹窗
const closeDetailShow = () => {
	detailShow.value = false
	formData.value = {
		account: '',
		type: '',
		staffId: '',
		name: '',
		credentialType: '',
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
		jobNum: '',
		typeText: '',
		credentialTypeText: '',
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
		account: '',
		type: '',
		staffId: '',
		name: '',
		credentialType: '',
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
		jobNum: '',
		typeText: '',
		credentialTypeText: '',
	}
}
// 弹窗确定
const enterDialog = async () => {
	elFormRef.value?.validate(async (valid) => {
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
				message: '创建/更改成功',
			})
			closeDialog()
			getTableData()
		}
	})
}
</script>

<style></style>
