<template>
	<div>
		<div class="gva-search-box">
			<el-image
				style="width: 120px; height: 120px; border-radius: 10px"
				src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
			/>
			<span>xxx</span>
		</div>
		<div class="gva-table-box" v-loading="loading">
			<el-tabs v-model="activeName">
				<el-tab-pane label="基础信息" name="first">
					<baseInfo
						:rewcStaff="rewcStaff"
						:rewcStaffJob="rewcStaffJob"
						:rewcStaffEducation="rewcStaffEducation"
						:rewcStaffBank="rewcStaffBank"
						:rewcStaffContact="rewcStaffContact"
						@updateInfo="updateInfo"
					/>
				</el-tab-pane>
				<el-tab-pane label="合同信息" name="second">
					<contractInfo :id="id" :rewcStaffAgreement="rewcStaffAgreement" @updateInfoSuccess="getData" />
				</el-tab-pane>
				<el-tab-pane label="材料附件" name="third">
					<material :rewcStaffMaterials="rewcStaffMaterials" />
				</el-tab-pane>
			</el-tabs>
		</div>
		<wcStaffForm :id="id" :rewcStaff="rewcStaff" ref="wcStaffFormRef" @updateInfoSuccess="getData" />
		<wcStaffJobForm :id="id" :rewcStaffJob="rewcStaffJob" ref="wcStaffJobFormRef" @updateInfoSuccess="getData" />
		<wcStaffEducationForm :id="id" :rewcStaffEducation="rewcStaffEducation" ref="wcStaffEducationRef" @updateInfoSuccess="getData" />
		<wcStaffBankForm :id="id" :rewcStaffBank="rewcStaffBank" ref="wcStaffBankRef" @updateInfoSuccess="getData" />
		<wcStaffContactForm :id="id" :rewcStaffContact="rewcStaffContact" ref="wcStaffContactRef" @updateInfoSuccess="getData" />
	</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import baseInfo from './components/baseInfo.vue'
import contractInfo from './components/contractInfo.vue'
import material from './components/material.vue'

import wcStaffForm from './components/wcStaffForm.vue'
import wcStaffJobForm from './components/wcStaffJobForm.vue'
import wcStaffEducationForm from './components/wcStaffEducationForm.vue'
import wcStaffBankForm from './components/wcStaffBankForm.vue'
import wcStaffContactForm from './components/wcStaffContactForm.vue'

import { getObtainEmployeeRoster } from '@/api/weChat/wcStaff'

const route = useRoute()

const activeName = ref('first')
const id = ref(route.query.id)

const rewcStaff = ref({})
const rewcStaffJob = ref({})
const rewcStaffEducation = ref({})
const rewcStaffBank = ref({})
const rewcStaffContact = ref({})

const rewcStaffAgreement = ref([])

const rewcStaffMaterials = ref({})

const getData = async () => {
	loading.value = true
	const res = await getObtainEmployeeRoster({ id: id.value })
	if (res.code === 0) {
		loading.value = false
		rewcStaff.value = res.data.rewcStaff
		rewcStaffJob.value = res.data.rewcStaffJob
		rewcStaffEducation.value = res.data.rewcStaffEducation
		rewcStaffBank.value = res.data.rewcStaffBank
		rewcStaffContact.value = res.data.rewcStaffContact

		rewcStaffAgreement.value = res.data.rewcStaffAgreement

		rewcStaffMaterials.value = res.data.rewcStaffMaterials
	}
}

const loading = ref(false)

const wcStaffFormRef = ref()
const wcStaffJobFormRef = ref()
const wcStaffEducationRef = ref()
const wcStaffBankRef = ref()
const wcStaffContactRef = ref()

const updateInfo = (type) => {
	switch (type) {
		case 1:
			wcStaffFormRef.value.dialogFormVisible = true
			break
		case 2:
		  wcStaffJobFormRef.value.dialogFormVisible = true
			break
		case 3:
		  wcStaffEducationRef.value.dialogFormVisible = true
			break
		case 4:
		  wcStaffBankRef.value.dialogFormVisible = true
			break
		case 5:
		  wcStaffContactRef.value.dialogFormVisible = true
			break
		default:
			break
	}
}

onMounted(() => {
	getData()
})
</script>

<style></style>
