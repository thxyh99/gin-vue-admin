<template>
	<el-form label-width="180px" label-position="right">
		<el-row>
			<el-col :span="8">
				<el-form-item label="身份证（人像面）">
					<el-upload list-type="picture-card">
						<el-icon><Plus /></el-icon>
					</el-upload>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="身份证（国徽面）">
					<el-upload list-type="picture-card">
						<el-icon><Plus /></el-icon>
					</el-upload>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="学历证书">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
		</el-row>
		<el-row>
			<el-col :span="8">
				<el-form-item label="学位证书">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="职称/技能证书">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="前公司离职证明">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
		</el-row>
		<el-row>
			<el-col :span="8">
				<el-form-item label="员工入职申请表">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="试用期管理规定">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
			<el-col :span="8">
				<el-form-item label="个人简历">
					<el-upload>
						<el-button size="small" type="primary">点击上传</el-button>
					</el-upload>
				</el-form-item>
			</el-col>
		</el-row>
	</el-form>
</template>

<script setup>
defineProps({
	rewcStaffMaterials: {
		type: Object,
		default: () => {
			return {}
		},
	},
})
</script>
