<template>
	<div class="assistant-container">
		<el-card class="assistant-baseInfo">
			<div class="baseInfo-title">
				<span class="name">李逍遥</span>
				<span class="time">已入职1000天</span>
			</div>
			<div class="baseInfo-form">
				<div class="form-item">
					<span class="label">类型</span>
					<span class="value">全职</span>
				</div>
				<div class="form-item">
					<span class="label">状态</span>
					<span class="value">正式</span>
				</div>
				<div class="form-item">
					<span class="label">部门</span>
					<span class="value">蜀山派</span>
				</div>
				<div class="form-item">
					<span class="label">上级</span>
					<span class="value">酒剑仙</span>
				</div>
				<div class="form-item">
					<span class="label">职位</span>
					<span class="value">掌门人</span>
				</div>
			</div>
			<div class="baseInfo-btn">
				<el-button round>查看员工档案 ></el-button>
			</div>
		</el-card>
		<el-card>
			<div class="assistant-salary">
				<div class="salary-item">
					<el-icon size="20px">
						<Edit />
					</el-icon>
					<span class="value">工资条</span>
				</div>
				<div class="salary-item">
					<el-icon size="20px">
						<Edit />
					</el-icon>
					<span class="value">假期</span>
				</div>
			</div>
		</el-card>
		<el-card>
			<div class="assistant-salary">
				<div class="salary-item">
					<el-icon size="20px">
						<Edit />
					</el-icon>
					<span class="value">人资类证明用印申请</span>
				</div>
				<div class="salary-item">
					<el-icon size="20px">
						<Edit />
					</el-icon>
					<span class="value">调动申请</span>
				</div>
        <div class="salary-item">
					<el-icon size="20px">
						<Edit />
					</el-icon>
					<span class="value">劳动关系解除审批流程</span>
				</div>
        <div class="salary-item">
					<el-icon size="20px">
						<Edit />
					</el-icon>
					<span class="value">申请记录</span>
				</div>
			</div>
		</el-card>
	</div>
</template>

<script setup></script>

<style lang="scss" scoped>
.assistant-container {
	height: 100%;
	background-color: #fff;
	padding: 20px;
	position: relative;
	.assistant-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		height: 140px;
		background-color: rgba(72, 115, 192, 1);
	}
	.assistant-baseInfo {
		.baseInfo-title {
			margin-bottom: 40px;
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			.name {
				font-size: 18px;
				font-weight: bold;
				color: #000;
			}
			.time {
				font-size: 14px;
				color: #999;
				margin-top: 10px;
			}
		}
		.baseInfo-form {
			display: flex;
			flex-wrap: wrap;
			margin-bottom: 20px;
			.form-item {
				margin-bottom: 18px;
				width: 50%;
				display: flex;
				.label {
					width: 40%;
					margin-right: 20px;
					font-size: 16px;
					text-align: right;
					color: #999;
				}
				.value {
					font-size: 16px;
					color: #000;
					font-weight: bold;
				}
			}
		}
		.baseInfo-btn {
			display: flex;
			justify-content: center;
			align-items: center;
		}
	}
	.assistant-salary {
		display: flex;
		align-items: flex-start;
    flex-wrap: wrap;
		.salary-item {
      cursor: pointer;
      margin-bottom: 20px;
			width: 33%;
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			.value {
        text-align: center;
				margin-top: 10px;
				font-size: 14px;
				color: #000;
			}
		}
	}
}
</style>
