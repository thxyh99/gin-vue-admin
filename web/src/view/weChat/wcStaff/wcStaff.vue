<template>
	<div>
		<div class="flex mb-3">
			<el-card class="flex-1 mr-1">
				<div class="flex justify-around">
					<div class="flex flex-col items-center">
						<div class="mb-1">在职</div>
						<div>{{ statistics.onjobCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">全职</div>
						<div>{{ statistics.fullTimeCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">兼职</div>
						<div>{{ statistics.partTimeCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">跟岗实习</div>
						<div>{{ statistics.followUpJobCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">顶岗实习</div>
						<div>{{ statistics.replaceCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">退休返聘</div>
						<div>{{ statistics.retireCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">劳务外包</div>
						<div>{{ statistics.outsourcingCount }}</div>
					</div>
				</div>
			</el-card>
			<el-card class="flex-1 mr-1">
				<div class="flex justify-around">
					<div class="flex flex-col items-center">
						<div class="mb-1">待入职</div>
						<div>{{ statistics.toBeEmployedCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">试用</div>
						<div>{{ statistics.probationCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">正式</div>
						<div>{{ statistics.formalCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center">
						<div class="mb-1">待离职</div>
						<div>{{ statistics.toBeDepartedCount }}</div>
					</div>
					<el-divider class="h-10" direction="vertical"></el-divider>
					<div class="flex flex-col items-center" @click="handleTodo">
						<div class="mb-1 text-blue-600/100">待办事项</div>
						<div class="text-blue-600/100">{{ statistics.toDoCount }}</div>
					</div>
				</div>
			</el-card>
		</div>
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
					<el-input v-model="searchInfo.keyword" placeholder="输入名称、员工编码、手机" />
				</el-form-item>

				<el-form-item label="员工" prop="staffId">
					<SelectStaff v-model="searchInfo.staffId" :disabled="type === 'update' ? 'disabled' : false"> </SelectStaff>
				</el-form-item>

				<el-form-item label="部门" prop="departmentIds">
					<SelectDepartment v-model="searchInfo.departmentIds" style="width: 300px"> </SelectDepartment>
				</el-form-item>

				<el-form-item label="入职日期">
					<el-date-picker
						style="width: 220px"
						type="daterange"
						range-separator="至"
						start-placeholder="开始时间"
						end-placeholder="结束时间"
						v-model="searchInfo.employmentDateRange"
						value-format="YYYY-MM-DD"
						@change="handleDataChange"
					/>
				</el-form-item>

				<el-form-item label="历史日期">
					<el-date-picker v-model="searchInfo.historyDate" type="date" value-format="YYYY-MM-DD" />
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
				<!--        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">-->
				<!--          删除-->
				<!--        </el-button>-->
				<ExportTemplate template-id="roster" />
				<ExportExcel template-id="roster" :limit="9999" />
				<ImportExcel template-id="roster" @on-success="getTableData" btn-name="导入花名册" />
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
				<el-table-column align="left" label="姓名" prop="name" width="120">
					<template #default="scope">
						<el-link type="primary" @click="jumpRoute(scope.row)">{{ scope.row.name }}</el-link>
					</template>
				</el-table-column>
				<el-table-column align="left" label="员工编码" prop="jobNum" width="120" />
				<el-table-column align="left" label="手机" prop="mobile" width="140" />
				<el-table-column align="left" label="户籍类型" prop="householdTypeText" width="150" />
				<el-table-column align="left" label="性别" prop="genderText" width="60" />
				<el-table-column align="left" label="出生日期" width="100">
					<template #default="scope">{{ formatDate(scope.row.birthday) }}</template>
				</el-table-column>
				<el-table-column align="left" label="籍贯" prop="nativePlace" width="120" />
				<el-table-column align="left" label="民族" prop="nationText" width="60" />
				<el-table-column align="left" label="身高(cm)" prop="height" width="90" />
				<el-table-column align="left" label="体重(kg)" prop="weight" width="90" />
				<el-table-column align="left" label="婚姻状况" prop="marriageText" width="90" />
				<el-table-column align="left" label="政治面貌" prop="politicalOutlookText" width="90" />
				<el-table-column align="left" label="操作" fixed="right" min-width="240">
					<template #default="scope">
						<el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
							<el-icon style="margin-right: 5px">
								<InfoFilled />
							</el-icon>
							查看详情
						</el-button>
						<!--            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
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
				<el-form-item label="姓名:" prop="name">
					<el-input v-model="formData.name" :clearable="true" placeholder="请输入姓名" />
				</el-form-item>
				<el-form-item label="员工编码:" prop="jobNum">
					<el-input v-model="formData.jobNum" :clearable="true" placeholder="请输入员工编码" />
				</el-form-item>
				<el-form-item label="企微账号:" prop="userid">
					<el-input v-model="formData.userid" :clearable="true" placeholder="请输入企微账号" />
				</el-form-item>
				<el-form-item label="手机:" prop="mobile">
					<el-input v-model="formData.mobile" :clearable="true" placeholder="请输入手机" />
				</el-form-item>
				<el-form-item label="身份证号:" prop="idNumber">
					<el-input v-model="formData.idNumber" :clearable="true" placeholder="请输入身份证号" />
				</el-form-item>
				<el-form-item label="身份证地址:" prop="idAddress">
					<el-input v-model="formData.idAddress" :clearable="true" placeholder="请输入身份证地址" />
				</el-form-item>
				<el-form-item label="户籍类型:" prop="householdType">
					<el-select v-model="formData.householdType" placeholder="选择户籍类型">
						<el-option
							v-for="householdType in householdTypes"
							:key="householdType.value"
							:label="householdType.label"
							:value="householdType.value"
						></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="性别:" prop="gender">
					<el-select v-model="formData.gender" placeholder="选择性别">
						<el-option
							v-for="gender in genders"
							:key="gender.value"
							:label="gender.label"
							:value="gender.value"
						></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="出生日期:" prop="birthday">
					<el-date-picker
						v-model="formData.birthday"
						type="date"
						style="width: 100%"
						placeholder="选择日期"
						:clearable="true"
					/>
				</el-form-item>
				<el-form-item label="籍贯:" prop="nativePlace">
					<el-input v-model="formData.nativePlace" :clearable="true" placeholder="请输入籍贯" />
				</el-form-item>
				<el-form-item label="民族:" prop="nation">
					<el-select v-model="formData.nation" placeholder="选择民族" filterable>
						<el-option
							v-for="nation in nations"
							:key="nation.value"
							:label="nation.label"
							:value="nation.value"
						></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="身高(cm):" prop="height">
					<el-input-number v-model="formData.height" style="width: 100%" :precision="1" :clearable="true" />
				</el-form-item>
				<el-form-item label="体重(kg):" prop="weight">
					<el-input-number v-model="formData.weight" style="width: 100%" :precision="1" :clearable="true" />
				</el-form-item>
				<el-form-item label="婚姻状况:" prop="marriage">
					<el-select v-model="formData.marriage" placeholder="选择婚姻状况">
						<el-option
							v-for="marriage in marriages"
							:key="marriage.value"
							:label="marriage.label"
							:value="marriage.value"
						></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="政治面貌:" prop="politicalOutlook">
					<el-select v-model="formData.politicalOutlook" placeholder="选择政治面貌">
						<el-option
							v-for="politicalOutlook in politicalOutlooks"
							:key="politicalOutlook.value"
							:label="politicalOutlook.label"
							:value="politicalOutlook.value"
						></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="现居住地址:" prop="address">
					<el-input v-model="formData.address" :clearable="true" placeholder="请输入现居住地址" />
				</el-form-item>
				<el-form-item label="社保电脑号:" prop="socialNumber">
					<el-input v-model="formData.socialNumber" :clearable="true" placeholder="请输入社保电脑号" />
				</el-form-item>
				<el-form-item label="公积金账号:" prop="accountNumber">
					<el-input v-model="formData.accountNumber" :clearable="true" placeholder="请输入公积金账号" />
				</el-form-item>
				<el-form-item label="社保公积金缴纳地:" prop="paymentPlace">
					<el-input v-model="formData.paymentPlace" :clearable="true" placeholder="请输入社保公积金缴纳地" />
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
				<el-descriptions-item label="姓名">
					{{ formData.name }}
				</el-descriptions-item>
				<el-descriptions-item label="员工编码">
					{{ formData.jobNum }}
				</el-descriptions-item>
				<el-descriptions-item label="企微账号">
					{{ formData.userid }}
				</el-descriptions-item>
				<el-descriptions-item label="手机">
					{{ formData.mobile }}
				</el-descriptions-item>
				<el-descriptions-item label="身份证号">
					{{ formData.idNumber }}
				</el-descriptions-item>
				<el-descriptions-item label="身份证地址">
					{{ formData.idAddress }}
				</el-descriptions-item>
				<el-descriptions-item label="户籍类型">
					{{ formData.householdTypeText }}
				</el-descriptions-item>
				<el-descriptions-item label="性别">
					{{ formData.genderText }}
				</el-descriptions-item>
				<el-descriptions-item label="出生日期">
					{{ formatDate(formData.birthday) }}
				</el-descriptions-item>
				<el-descriptions-item label="籍贯">
					{{ formData.nativePlace }}
				</el-descriptions-item>
				<el-descriptions-item label="民族">
					{{ formData.nationText }}
				</el-descriptions-item>
				<el-descriptions-item label="身高(cm)">
					{{ formData.height }}
				</el-descriptions-item>
				<el-descriptions-item label="体重(kg)">
					{{ formData.weight }}
				</el-descriptions-item>
				<el-descriptions-item label="婚姻状况">
					{{ formData.marriageText }}
				</el-descriptions-item>
				<el-descriptions-item label="政治面貌">
					{{ formData.politicalOutlookText }}
				</el-descriptions-item>
				<el-descriptions-item label="现居住地址">
					{{ formData.address }}
				</el-descriptions-item>
				<el-descriptions-item label="社保电脑号">
					{{ formData.socialNumber }}
				</el-descriptions-item>
				<el-descriptions-item label="公积金账号">
					{{ formData.accountNumber }}
				</el-descriptions-item>
				<el-descriptions-item label="社保公积金缴纳地">
					{{ formData.paymentPlace }}
				</el-descriptions-item>
			</el-descriptions>
		</el-drawer>

		<todo ref="todoRef" />
	</div>
</template>

<script setup>
import {
	createWcStaff,
	deleteWcStaff,
	deleteWcStaffByIds,
	updateWcStaff,
	findWcStaff,
	getWcStaffList,
} from '@/api/weChat/wcStaff'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import ImportExcel from '@/components/exportExcel/wechat/importExcel.vue'
import ExportExcel from '@/components/exportExcel/wechat/exportExcel.vue'
import ExportTemplate from '@/components/exportExcel/wechat/exportTemplate.vue'
import { InfoFilled, QuestionFilled } from '@element-plus/icons-vue'
import SelectDepartment from '@/components/selectDepartment/index.vue'
import SelectStaff from '@/components/selectStaff/index.vue'
import todo from './components/tode.vue'

defineOptions({
	name: 'WcStaff',
})

const router = useRouter()

const householdTypes = ref([
	{ label: '本地城镇', value: 1 },
	{ label: '本地农村', value: 2 },
	{ label: '外地城镇（省内）', value: 3 },
	{ label: '外地农村（省内）', value: 4 },
	{ label: '外地城镇（省外）', value: 5 },
	{ label: '外地农村（省外）', value: 6 },
])

const marriages = ref([
	{ label: '已婚', value: 1 },
	{ label: '未婚', value: 2 },
	{ label: '其他', value: 3 },
])

const politicalOutlooks = ref([
	{ label: '其他', value: 0 },
	{ label: '团员', value: 1 },
	{ label: '党员', value: 2 },
	{ label: '群众', value: 3 },
])

const nations = ref([
	{ label: '汉族', value: 1 },
	{ label: '蒙古族', value: 2 },
	{ label: '回族', value: 3 },
	{ label: '藏族', value: 4 },
	{ label: '维吾尔族', value: 5 },
	{ label: '苗族', value: 6 },
	{ label: '彝族', value: 7 },
	{ label: '壮族', value: 8 },
	{ label: '布依族', value: 9 },
	{ label: '朝鲜族', value: 10 },
	{ label: '满族', value: 11 },
	{ label: '侗族', value: 12 },
	{ label: '瑶族', value: 13 },
	{ label: '白族', value: 14 },
	{ label: '土家族', value: 15 },
	{ label: '哈尼族', value: 16 },
	{ label: '哈萨克族', value: 17 },
	{ label: '傣族', value: 18 },
	{ label: '黎族', value: 19 },
	{ label: '傈僳族', value: 20 },
	{ label: '佤族', value: 21 },
	{ label: '畲族', value: 22 },
	{ label: '高山族', value: 23 },
	{ label: '拉祜族', value: 24 },
	{ label: '水族', value: 25 },
	{ label: '东乡族', value: 26 },
	{ label: '纳西族', value: 27 },
	{ label: '景颇族', value: 28 },
	{ label: '柯尔克孜族', value: 29 },
	{ label: '土族', value: 30 },
	{ label: '达斡尔族', value: 31 },
	{ label: '仫佬族', value: 32 },
	{ label: '羌族', value: 33 },
	{ label: '布朗族', value: 34 },
	{ label: '撒拉族', value: 35 },
	{ label: '毛南族', value: 36 },
	{ label: '仡佬族', value: 37 },
	{ label: '锡伯族', value: 38 },
	{ label: '阿昌族', value: 39 },
	{ label: '普米族', value: 40 },
	{ label: '塔吉克族', value: 41 },
	{ label: '怒族', value: 42 },
	{ label: '乌孜别克族', value: 43 },
	{ label: '俄罗斯族', value: 44 },
	{ label: '鄂温克族', value: 45 },
	{ label: '崩龙族', value: 46 },
	{ label: '保安族', value: 47 },
	{ label: '裕固族', value: 48 },
	{ label: '京族', value: 49 },
	{ label: '塔塔尔族', value: 50 },
	{ label: '独龙族', value: 51 },
	{ label: '鄂伦春族', value: 52 },
	{ label: '赫哲族', value: 53 },
	{ label: '门巴族', value: 54 },
	{ label: '珞巴族', value: 55 },
	{ label: '基诺族', value: 56 },
])

const genders = ref([
	{ label: '未知', value: 0 },
	{ label: '男', value: 1 },
	{ label: '女', value: 2 },
])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
	userId: 0,
	userid: '',
	jobNum: '',
	name: '',
	gender: '',
	mobile: '',
	idNumber: '',
	idAddress: '',
	householdType: '',
	birthday: new Date(),
	nativePlace: '',
	nation: '',
	height: 0,
	weight: 0,
	marriage: '',
	politicalOutlook: '',
	address: '',
	socialNumber: '',
	accountNumber: '',
	paymentPlace: '',
	genderText: '',
	householdTypeText: '',
	nationText: '',
	marriageText: '',
	politicalOutlookText: '',
})

// 验证规则
const rule = reactive({
	userId: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	userid: [
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
	jobNum: [
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
	gender: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	mobile: [
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
	idNumber: [
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
	idAddress: [
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
	householdType: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	birthday: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	nativePlace: [
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
	nation: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	height: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	weight: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	marriage: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	politicalOutlook: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	address: [
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
const searchInfo = ref({})

// 重置
const onReset = () => {
	searchInfo.value = {}
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

const handleDataChange = (e) => {
	if (e && e.length) {
		searchInfo.value.employmentDateStart = e[0]
		searchInfo.value.employmentDateEnd = e[1]
	} else {
		searchInfo.value.employmentDateStart = ''
		searchInfo.value.employmentDateEnd = ''
	}
}

const statistics = ref({})
// 查询
const getTableData = async () => {
	const table = await getWcStaffList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
	if (table.code === 0) {
		tableData.value = table.data.list
		total.value = table.data.total
		page.value = table.data.page
		pageSize.value = table.data.pageSize
		statistics.value = table.data.statistics
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
		deleteWcStaffFunc(row)
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
		const res = await deleteWcStaffByIds({ IDs })
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
const updateWcStaffFunc = async (row) => {
	const res = await findWcStaff({ ID: row.ID })
	type.value = 'update'
	if (res.code === 0) {
		// console.log(formData.value)
		formData.value = res.data.rewcStaff
		dialogFormVisible.value = true
	}
}

// 删除行
const deleteWcStaffFunc = async (row) => {
	const res = await deleteWcStaff({ ID: row.ID })
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
	const res = await findWcStaff({ ID: row.ID })
	if (res.code === 0) {
		formData.value = res.data.rewcStaff
		openDetailShow()
	}
}

// 关闭详情弹窗
const closeDetailShow = () => {
	detailShow.value = false
	formData.value = {
		userId: 0,
		userid: '',
		jobNum: '',
		name: '',
		gender: '',
		mobile: '',
		idNumber: '',
		idAddress: '',
		householdType: '',
		birthday: new Date(),
		nativePlace: '',
		nation: '',
		height: 0,
		weight: 0,
		marriage: '',
		politicalOutlook: '',
		address: '',
		socialNumber: '',
		accountNumber: '',
		paymentPlace: '',
		genderText: '',
		householdTypeText: '',
		nationText: '',
		marriageText: '',
		politicalOutlookText: '',
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
		userId: 0,
		userid: '',
		jobNum: '',
		name: '',
		gender: '',
		mobile: '',
		idNumber: '',
		idAddress: '',
		householdType: '',
		birthday: new Date(),
		nativePlace: '',
		nation: '',
		height: 0,
		weight: 0,
		marriage: '',
		politicalOutlook: '',
		address: '',
		socialNumber: '',
		accountNumber: '',
		paymentPlace: '',
		genderText: '',
		householdTypeText: '',
		nationText: '',
		marriageText: '',
		politicalOutlookText: '',
	}
}
// 弹窗确定
const enterDialog = async () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		let res
		switch (type.value) {
			case 'create':
				res = await createWcStaff(formData.value)
				break
			case 'update':
				res = await updateWcStaff(formData.value)
				break
			default:
				res = await createWcStaff(formData.value)
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

const jumpRoute = (row) => {
	router.push(`/layout/staff/roster?id=${row.ID}`)
}

const todoRef = ref(null)
const handleTodo = () => {
	todoRef.value.todoShow = true
}
</script>

<style></style>
