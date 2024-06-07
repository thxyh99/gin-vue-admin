<template>
	<div class="payslip-container">
		<el-card>
			<div class="payslip-header">工资条</div>
			<div class="payslip-content">
				<div class="content-item">
					<div class="label">等级工资</div>
					<div class="value">李逍遥</div>
				</div>
				<div class="content-item">
					<div class="label">职称技能津贴</div>
					<div class="value">12345678910</div>
				</div>
				<div class="content-item">
					<div class="label">津贴</div>
					<div class="value">440001441111111</div>
				</div>
				<div class="content-item">
					<div class="label">司龄津贴</div>
					<div class="value">浙江省余杭镇盛渔村李家客栈地字四号</div>
				</div>
				<div class="content-item">
					<div class="label">学历工资</div>
					<div class="value">本地城镇</div>
				</div>
				<div class="content-item">
					<div class="label">尽孝基金</div>
					<div class="value">男</div>
				</div>
				<div class="content-item">
					<div class="label">业绩奖</div>
					<div class="value">1083-12-19</div>
				</div>
				<div class="content-item">
					<div class="label">应发工资</div>
					<div class="value">浙江余杭</div>
				</div>
				<div class="content-item">
					<div class="label">养老保险</div>
					<div class="value">汉</div>
				</div>
				<div class="content-item">
					<div class="label">医疗保险</div>
					<div class="value">173</div>
				</div>
				<div class="content-item">
					<div class="label">实发工资</div>
					<div class="value">60公斤</div>
				</div>
			</div>
		</el-card>
	</div>
</template>

<script setup></script>

<style lang="scss" scoped>
.payslip-container {
	height: 100%;
	box-sizing: border-box;
	overflow: scroll;
	padding: 20px;
	.payslip-header {
		font-size: 18px;
		font-weight: bold;
		color: #000;
	}
	.payslip-content {
		display: flex;
		flex-direction: column;
		.content-item {
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin-top: 20px;
			.label,
			.value {
				font-size: 18px;
				color: #333;
			}
			.label {
				margin-right: 10px;
			}
			.value {
				flex: 1;
				text-align: right;
				white-space: nowrap;
				text-overflow: ellipsis;
				overflow: hidden;
				word-break: break-all;
			}
		}
	}
}
</style>
