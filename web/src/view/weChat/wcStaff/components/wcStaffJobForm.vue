<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false" @open="handleOpen">
		<template #title>
			<div class="flex justify-between items-center">
				<span class="text-lg">修改</span>
				<div>
					<el-button type="primary" @click="enterDialog">确 定</el-button>
					<el-button @click="closeDialog">取 消</el-button>
				</div>
			</div>
		</template>

		<el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
			<el-form-item label="员工类型:" prop="type">
				<el-select v-model="formData.type" placeholder="选择员工类型">
					<el-option v-for="type in types" :key="type.value" :label="type.label" :value="type.value"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="员工状态:" prop="status">
				<el-select v-model="formData.status" placeholder="选择员工状态">
					<el-option
						v-for="status in statusList"
						:key="status.value"
						:label="status.label"
						:value="status.value"
					></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="部门信息:" prop="departmentIds">
				<SelectDepartment v-model="formData.departmentIds"> </SelectDepartment>
			</el-form-item>
			<el-form-item label="岗位信息:" prop="positionIds">
				<SelectPosition v-model="formData.positionIds"> </SelectPosition>
			</el-form-item>
      <el-form-item label="层级:" prop="level">
        <el-select v-model="formData.level" placeholder="选择层级">
          <el-option
              v-for="level in levels"
              :key="level.value"
              :label="level.label"
              :value="level.value"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="内外勤:" prop="ioType">
        <el-select v-model="formData.ioType" placeholder="选择内外勤">
          <el-option
              v-for="ioType in ioTypes"
              :key="ioType.value"
              :label="ioType.label"
              :value="ioType.value"
          ></el-option>
        </el-select>
      </el-form-item>
			<el-form-item label="职级类型:" prop="rankType">
				<SelectRankType v-model="formData.rankType" placeholder="选择职级类型" @change="handleTypeChange">
				</SelectRankType>
			</el-form-item>
			<el-form-item label="职级:" prop="rank">
				<SelectRank ref="SelectRankRef" v-model="formData.rank" placeholder="选择职级"> </SelectRank>
			</el-form-item>
			<el-form-item label="等级工资:" prop="rankSalary">
				<el-input-number v-model="formData.rankSalary" style="width: 100%" :precision="0" :clearable="true" />
			</el-form-item>
			<el-form-item label="费用科目:" prop="expenseAccount">
				<el-select v-model="formData.expenseAccount" placeholder="选择费用科目">
					<el-option
						v-for="expenseAccount in expenseAccounts"
						:key="expenseAccount.value"
						:label="expenseAccount.label"
						:value="expenseAccount.value"
					></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="入职容大日期:" prop="employmentDate">
				<el-date-picker
					v-model="formData.employmentDate"
					type="date"
					style="width: 100%"
					placeholder="选择日期"
					:clearable="true"
				/>
			</el-form-item>
      <el-form-item label="入职总部日期:" prop="employmentHeadquarterDate">
				<el-date-picker
					v-model="formData.employmentHeadquarterDate"
					type="date"
					style="width: 100%"
					placeholder="选择日期"
					:clearable="true"
				/>
			</el-form-item>
			<el-form-item label="试用期:" prop="tryPeriod">
				<el-select v-model="formData.tryPeriod" placeholder="选择试用期">
					<el-option
						v-for="tryPeriod in tryPeriods"
						:key="tryPeriod.value"
						:label="tryPeriod.label"
						:value="tryPeriod.value"
					></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="转正日期:" prop="formalDate">
				<el-date-picker
					v-model="formData.formalDate"
					type="date"
					style="width: 100%"
					placeholder="选择日期"
					:clearable="true"
				/>
			</el-form-item>
      <el-form-item label="拟转正日期:" prop="presumeFormalDate">
        <el-date-picker
            v-model="formData.presumeFormalDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
        />
      </el-form-item>
      <el-form-item label="直属领导:" prop="leaderId">
        <SelectStaff v-model="formData.leaderId">
        </SelectStaff>
      </el-form-item>
		</el-form>
	</el-drawer>
</template>

<script setup>
import { ref, reactive, watchEffect, nextTick } from 'vue'

import SelectDepartment from '@/components/selectDepartment/index.vue'
import SelectPosition from '@/components/selectPosition/index.vue'
import SelectRankType from '@/components/selectRankType/index.vue'
import SelectRank from '@/components/selectRank/index.vue'
import { updateWcStaffJob } from '@/api/weChat/wcStaffJob'
import SelectStaff from "@/components/selectStaff/index.vue";

const props = defineProps({
	rewcStaffJob: {
		type: Object,
		default: () => {
			return {}
		},
	},
	id: {
		type: String,
		required: true,
	},
})

const emits = defineEmits(['updateInfoSuccess'])

const dialogFormVisible = ref(false)

const SelectRankRef = ref()
const formData = ref()
const elFormRef = ref()

const rule = reactive({
	staffId: [
		{
			required: true,
			message: '',
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
	status: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	tryPeriod: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	expenseAccount: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	employmentDate: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
  employmentHeadquarterDate: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	formalDate: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	presumeFormalDate: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	rankType: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	rank: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	departmentIds: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	positionIds: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
	rankSalary: [
		{
			required: true,
			message: '',
			trigger: ['input', 'blur'],
		},
	],
  leaderId: [
    {
      required: false,
      message: '',
      trigger: ['input', 'blur'],
    },
  ],
})

watchEffect(() => {
	formData.value = { ...props.rewcStaffJob }
})

const types = [
	{ label: '其他', value: 0 },
	{ label: '全职', value: 1 },
	{ label: '兼职', value: 2 },
	{ label: '跟岗实习', value: 3 },
	{ label: '顶岗实习', value: 4 },
	{ label: '退休返聘', value: 5 },
	{ label: '劳务外包', value: 6 },
]

const statusList = [
	{ label: '待入职', value: 0 },
	{ label: '试用', value: 1 },
	{ label: '正式', value: 2 },
	{ label: '待离职', value: 3 },
	{ label: '已离职', value: 4 },
]

const tryPeriods = [
	{ label: '其他', value: 0 },
	{ label: '无试用期', value: 1 },
	{ label: '2个月', value: 2 },
	{ label: '6个月', value: 3 },
]

const expenseAccounts = [
	{ label: '管理费用', value: 1 },
	{ label: '研发费用', value: 2 },
	{ label: '生产费用', value: 3 },
	{ label: '销售费用', value: 4 },
]

const levels = [
	{ label: '基层', value: 1 },
	{ label: '中层（副经理）', value: 2 },
	{ label: '中层', value: 3 },
	{ label: '高层（副总监级）', value: 4 },
	{ label: '高层（总监级）', value: 5 },
	{ label: '老总', value: 6 },
]

const ioTypes = [
	{ label: '内勤', value: 1 },
	{ label: '外勤', value: 2 },
]

const handleOpen = () => {
	if (formData.value.rankType) {
		SelectRankRef.value.initSelectOptions(formData.value.rankType)
	}
}

const handleTypeChange = (e) => {
	formData.value.rank = ''
	SelectRankRef.value.initSelectOptions(e)
}

const closeDialog = () => {
	formData.value = { ...props.rewcStaffJob }
	dialogFormVisible.value = false
}

const enterDialog = () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		const res = await updateWcStaffJob({ ...formData.value, staffId: Number(props.id), ID: Number(props.rewcStaffJob.ID) })
		if (res.code === 0) {
			emits('updateInfoSuccess')
			dialogFormVisible.value = false
		}
	})
}

defineExpose({
	dialogFormVisible,
})
</script>
