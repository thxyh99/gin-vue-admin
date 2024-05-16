<template>
	<div>
		<el-descriptions title="合同信息" border>
			<el-descriptions-item label="合同公司" label-align="right">蜀山派</el-descriptions-item>
			<el-descriptions-item label="合同类型" label-align="right">无固定期限劳动合同</el-descriptions-item>
			<el-descriptions-item label="合同起始日" label-align="right">1083-12-19</el-descriptions-item>
			<el-descriptions-item label="合同到期日" label-align="right">1183-12-19</el-descriptions-item>
			<el-descriptions-item label="续签次数" label-align="right">5</el-descriptions-item>
      <el-descriptions-item label-align="right"></el-descriptions-item>
			<el-descriptions-item label="合同附件" label-align="right" :span="3"></el-descriptions-item>
		</el-descriptions>
    <div class="btn">
      <span>+  添加合同</span>
    </div>
	</div>
</template>

<style scoped>
.btn {
  margin-top: 20px;
  padding: 10px;
  width: 100%;
  border: 1px dashed #ccc;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}
</style>