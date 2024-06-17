<template>
	<div>
		<el-descriptions title="个人信息" border>
			<template #extra>
				<el-button type="primary" size="small" @click="edit(1)">编 辑</el-button>
			</template>
			<el-descriptions-item label-align="right" label="姓名">{{ rewcStaff.name }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="员工编码">{{ rewcStaff.jobNum }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="手机号">{{ rewcStaff.mobile }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="身份证号">{{ rewcStaff.idNumber }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="身份证地址">{{ rewcStaff.idAddress }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="户籍类型">{{
				rewcStaff.householdTypeText
			}}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="出生日期">{{
				formatDate(rewcStaff.birthday)
			}}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="年龄（系统计算）">{{ getDateYearSub(formatDate(rewcStaff.birthday)) }}岁</el-descriptions-item>
			<el-descriptions-item label-align="right" label="籍贯">{{ rewcStaff.nativePlace }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="民族">{{ rewcStaff.nationText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="性别">{{ rewcStaff.genderText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="身高(cm)">{{ rewcStaff.height }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="体重(kg)">{{ rewcStaff.weight }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="婚否">{{ rewcStaff.marriageText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="政治面貌">{{ rewcStaff.politicalOutlookText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="常住地址">{{ rewcStaff.address }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="社保电脑号">{{ rewcStaff.socialNumber }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="公积金账号">{{ rewcStaff.accountNumber }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="社保公积金缴纳地">{{ rewcStaff.paymentPlaceText }}</el-descriptions-item>
		</el-descriptions>
		<el-divider></el-divider>
		<el-descriptions title="工作信息" border>
			<template #extra>
				<el-button type="primary" size="small" @click="edit(2)">编 辑</el-button>
			</template>
			<el-descriptions-item label-align="right" label="员工类型">{{ rewcStaffJob.typeText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="员工状态">{{ rewcStaffJob.statusText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="部门">{{ rewcStaffJob.department }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="岗位">{{ rewcStaffJob.position }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="层级">{{ rewcStaffJob.levelText }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="内外勤">{{ rewcStaffJob.ioTypeText }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="职级">{{ rewcStaffJob.rankText }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="等级工资">{{ rewcStaffJob.rankSalary }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="费用科目">{{ rewcStaffJob.expenseAccountText }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="入职容大日期">{{ formatDate(rewcStaffJob.employmentDate) }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="入职总部日期">{{ formatDate(rewcStaffJob.employmentHeadquarterDate) }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="试用期">{{ rewcStaffJob.tryPeriodText }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="转正日期">{{ formatDate(rewcStaffJob.formalDate) }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="拟转正日期">{{ formatDate(rewcStaffJob.presumeFormalDate) }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="离职日期">{{ formatDate(rewcStaffJob.leaveDate) }}</el-descriptions-item>
      <el-descriptions-item label-align="right" label="司龄（系统计算）">{{ getDateYearSub(formatDate(rewcStaffJob.employmentDate)) }}年</el-descriptions-item>
      <el-descriptions-item label-align="right" label="直属领导">{{ rewcStaffJob.leader }}</el-descriptions-item>
    </el-descriptions>
		<el-divider></el-divider>
		<el-descriptions title="学历信息" border>
			<template #extra>
				<el-button type="primary" size="small" @click="edit(3)">编 辑</el-button>
			</template>
			<el-descriptions-item label-align="right" label="学历">{{ rewcStaffEducation.educationText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="学历津贴">{{ rewcStaffEducation.educationPay }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="毕业院校">{{ rewcStaffEducation.school }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="毕业日期">{{ formatDate(rewcStaffEducation.date) }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="专业">{{ rewcStaffEducation.major }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="职称证书">{{ rewcStaffEducation.professionalCertificate }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="技能证书">{{ rewcStaffEducation.skillCertificate }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="职称技能津贴">{{ rewcStaffEducation.skillPay }}</el-descriptions-item>
		</el-descriptions>
		<el-divider></el-divider>
		<el-descriptions title="银行卡信息" border>
			<template #extra>
				<el-button type="primary" size="small" @click="edit(4)">编 辑</el-button>
			</template>
			<el-descriptions-item label-align="right" label="银行卡号">{{ rewcStaffBank.cardNumber }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="开户行">{{ rewcStaffBank.bank }}</el-descriptions-item>
		</el-descriptions>
		<el-divider></el-divider>
		<el-descriptions title="紧急联系人信息" border>
			<template #extra>
				<el-button type="primary" size="small" @click="edit(5)">编 辑</el-button>
			</template>
			<el-descriptions-item label-align="right" label="紧急联系人姓名">{{ rewcStaffContact.name }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="联系人关系">{{ rewcStaffContact.relationshipText }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="联系人电话">{{ rewcStaffContact.mobile }}</el-descriptions-item>
			<el-descriptions-item label-align="right" label="联系人常住地址">{{ rewcStaffContact.address }}</el-descriptions-item>
		</el-descriptions>
		
	</div>
</template>

<script setup>
import { formatDate } from '@/utils/format'

const emits = defineEmits(['updateInfo'])

defineProps({
	rewcStaff: {
		type: Object,
		default: () => {
			return {}
		},
	},
	rewcStaffJob: {
			type: Object,
			default: () => {
				return {}
			},
		},
		rewcStaffEducation: {
			type: Object,
			default: () => {
				return {}
			},
		},
		rewcStaffBank: {
			type: Object,
			default: () => {
				return {}
			},
		},
		rewcStaffContact: {
			type: Object,
			default: () => {
				return {}
			},
		}
})

const edit = (type) => {
	emits('updateInfo', type)
}

const getDateYearSub = (startDateStr) => {
	return new Date().getFullYear() - startDateStr.split('-')[0]
}
</script>
