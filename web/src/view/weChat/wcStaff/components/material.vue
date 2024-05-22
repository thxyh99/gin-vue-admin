<template>
	<el-form label-width="180px" label-position="right">
		<el-row>
			<el-col :span="8">
				<el-form-item label="身份证（人像面）">
					<upload
						v-model:file-list="idCardPortraitList"
						:staffId="props.id"
						:type="2"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="身份证（国徽面）">
					<upload
						v-model:file-list="idCardNational"
						:staffId="props.id"
						:type="3"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="学历证书">
					<upload
						v-model:file-list="educationCertificate"
						:staffId="props.id"
						:type="4"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
		</el-row>
		<el-row>
			<el-col :span="8">
				<el-form-item label="学位证书">
					<upload
						v-model:file-list="degreeCertificate"
						:staffId="props.id"
						:type="5"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="职称/技能证书">
					<upload
						v-model:file-list="skillCertificate"
						:staffId="props.id"
						:type="10"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="前公司离职证明">
					<upload
						v-model:file-list="resignationCertificate"
						:staffId="props.id"
						:type="6"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
		</el-row>
		<el-row>
			<el-col :span="8">
				<el-form-item label="员工入职申请表">
					<upload
						v-model:file-list="onboardingForm"
						:staffId="props.id"
						:type="7"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="试用期管理规定">
					<upload
						v-model:file-list="trialProvide"
						:staffId="props.id"
						:type="8"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="个人简历">
					<upload
						v-model:file-list="personalResume"
						:staffId="props.id"
						:type="9"
						:limit="1"
						@updateInfoSuccess="updateInfoSuccess"
					/>
				</el-form-item>
			</el-col>
		</el-row>
	</el-form>
</template>

<script setup>
import { ref, watchEffect } from 'vue'

import upload from './upload.vue'
const props = defineProps({
	id: {
		type: String,
		default: '',
	},
	rewcStaffMaterials: {
		type: Object,
		default: () => {
			return {}
		},
	},
})

const emits = defineEmits(['updateInfoSuccess'])

const updateInfoSuccess = () => {
	emits('updateInfoSuccess')
}

const idCardPortraitList = ref([])
const idCardNational = ref([])
const educationCertificate = ref([])
const degreeCertificate = ref([])
const skillCertificate = ref([])
const resignationCertificate = ref([])
const onboardingForm = ref([])
const trialProvide = ref([])
const personalResume = ref([])

watchEffect(() => {
	idCardPortraitList.value = props.rewcStaffMaterials.idCardPortrait || []
	idCardNational.value = props.rewcStaffMaterials.idCardNational || []
	educationCertificate.value = props.rewcStaffMaterials.educationCertificate || []
	degreeCertificate.value = props.rewcStaffMaterials.degreeCertificate || []
	skillCertificate.value = props.rewcStaffMaterials.skillCertificate || []
	resignationCertificate.value = props.rewcStaffMaterials.resignationCertificate || []
	onboardingForm.value = props.rewcStaffMaterials.onboardingForm || []
	trialProvide.value = props.rewcStaffMaterials.trialProvide || []
	personalResume.value = props.rewcStaffMaterials.personalResume || []
})

</script>

<style scoped>
.img-box {
	position: relative;
	transition: all 0.3s ease-in-out;
	.tip {
		width: 24px;
		height: 24px;
		position: absolute;
		top: 0;
		right: -12px;
		background-color: #ccc;
		border-radius: 50%;
		cursor: pointer;
	}
}
</style>
