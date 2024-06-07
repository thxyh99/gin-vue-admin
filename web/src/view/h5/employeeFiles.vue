<!-- 员工档案 -->
<template>
	<div class="employeeFiles-container">
		<el-card class="employeeFiles-baseInfo">
			<div class="baseInfo-header">基本信息</div>
			<div class="baseInfo-content">
				<div class="content-item">
					<div class="label">姓名</div>
					<div class="value">李逍遥</div>
				</div>
				<div class="content-item">
					<div class="label">手机</div>
					<div class="value">12345678910</div>
				</div>
				<div class="content-item">
					<div class="label">身份证号</div>
					<div class="value">440001441111111</div>
				</div>
				<div class="content-item">
					<div class="label">身份证地址</div>
					<div class="value">浙江省余杭镇盛渔村李家客栈地字四号</div>
				</div>
				<div class="content-item">
					<div class="label">户籍类型</div>
					<div class="value">本地城镇</div>
				</div>
				<div class="content-item">
					<div class="label">性别</div>
					<div class="value">男</div>
				</div>
				<div class="content-item">
					<div class="label">出生日期</div>
					<div class="value">1083-12-19</div>
				</div>
				<div class="content-item">
					<div class="label">籍贯</div>
					<div class="value">浙江余杭</div>
				</div>
				<div class="content-item">
					<div class="label">民族</div>
					<div class="value">汉</div>
				</div>
				<div class="content-item">
					<div class="label">身高</div>
					<div class="value">173</div>
				</div>
				<div class="content-item">
					<div class="label">体重</div>
					<div class="value">60公斤</div>
				</div>
				<div class="content-item">
					<div class="label">婚否</div>
					<div class="value">已婚</div>
				</div>
				<div class="content-item">
					<div class="label">政治面貌</div>
					<div class="value">群众</div>
				</div>
				<div class="content-item">
					<div class="label">常住地址</div>
					<div class="value">浙江省余杭镇盛渔村李家客栈地字四号…</div>
				</div>
				<div class="content-item">
					<div class="label">社保电脑号</div>
					<div class="value">10000</div>
				</div>
				<div class="content-item">
					<div class="label">公积金账号</div>
					<div class="value">10000</div>
				</div>
				<div class="content-item">
					<div class="label">社保公积金缴纳地</div>
					<div class="value">浙江省余杭镇</div>
				</div>
			</div>
		</el-card>
		<el-card class="employeeFiles-baseInfo">
			<div class="baseInfo-header">工作信息</div>
			<div class="baseInfo-content">
				<div class="content-item">
					<div class="label">员工类型</div>
					<div class="value">全职</div>
				</div>
				<div class="content-item">
					<div class="label">员工状态</div>
					<div class="value">正式</div>
				</div>
				<div class="content-item">
					<div class="label">部门</div>
					<div class="value">蜀山七圣</div>
				</div>
				<div class="content-item">
					<div class="label">职位</div>
					<div class="value">蜀山七圣之一</div>
				</div>
				<div class="content-item">
					<div class="label">出生日期</div>
					<div class="value">1083-12-19</div>
				</div>
			</div>
		</el-card>
	</div>
</template>

<script setup></script>

<style lang="scss" scoped>
.employeeFiles-container {
  height: 100%;
  box-sizing: border-box;
  overflow: scroll;
	padding: 20px;
	.employeeFiles-baseInfo {
		.baseInfo-header {
			font-size: 18px;
			font-weight: bold;
			color: #000;
		}
		.baseInfo-content {
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
}
</style>
