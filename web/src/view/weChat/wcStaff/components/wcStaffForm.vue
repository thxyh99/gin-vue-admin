<template>
	<el-drawer size="800" v-model="dialogFormVisible" :show-close="false">
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
			<el-form-item label="姓名:" prop="name">
				<el-input v-model="formData.name" :clearable="true" placeholder="请输入成员名称" />
			</el-form-item>
			<el-form-item label="员工编码:" prop="jobNum">
				<el-input v-model="formData.jobNum" :clearable="true" placeholder="请输入员工工号" />
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
					<el-option label="未知" :value="0"></el-option>
					<el-option label="男" :value="1"></el-option>
					<el-option label="女" :value="2"></el-option>
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
				<el-input-number v-model="formData.height" style="width: 100%" :precision="2" :clearable="true" />
			</el-form-item>
			<el-form-item label="体重(kg):" prop="weight">
				<el-input-number v-model="formData.weight" style="width: 100%" :precision="2" :clearable="true" />
			</el-form-item>
			<el-form-item label="婚否:" prop="marriage">
				<el-select v-model="formData.marriage" placeholder="选择婚否">
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
			<el-form-item label="常住地址:" prop="address">
				<el-input v-model="formData.address" :clearable="true" placeholder="请输入常住地址" />
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
</template>

<script setup>
import { ref, reactive, watchEffect } from 'vue'

import { updateWcStaff } from '@/api/weChat/wcStaff'

const props = defineProps({
	rewcStaff: {
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

const formData = ref()

watchEffect(() => {
	formData.value = { ...props.rewcStaff }
})

const marriages = [
	{ label: '已婚', value: 1 },
	{ label: '未婚', value: 2 },
	{ label: '其他', value: 3 },
]

const householdTypes = [
	{ label: '未知', value: 0 },
	{ label: '本地城镇', value: 1 },
	{ label: '本地农村', value: 2 },
	{ label: '外地城镇（省内）', value: 3 },
	{ label: '外地农村（省内）', value: 4 },
	{ label: '外地城镇（省外）', value: 5 },
	{ label: '外地农村（省外）', value: 6 },
]
const politicalOutlooks = [
	{ label: '其他', value: 0 },
	{ label: '团员', value: 1 },
	{ label: '党员', value: 2 },
	{ label: '群众', value: 3 },
]

const nations = [
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
]

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

const dialogFormVisible = ref(false)

const elFormRef = ref()

const closeDialog = () => {
	formData.value = { ...props.rewcStaff }
	dialogFormVisible.value = false
}

const enterDialog = async () => {
	elFormRef.value?.validate(async (valid) => {
		if (!valid) return
		const res = await updateWcStaff({ ...formData.value, ID: Number(props.id) })
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
