<template>
	<div>
		<div class="gva-search-box">
			<el-form ref="elSearchFormRef" :inline="true" :model="searchInfo">
				<el-form-item label="工资年份">
					<el-date-picker style="width: 250px" type="year" />
				</el-form-item>
				<el-form-item label="员工">
					<el-select style="width: 250px"></el-select>
				</el-form-item>
				<el-form-item label="部门">
					<el-select style="width: 250px"></el-select>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" icon="search">查 询</el-button>
					<el-button icon="delete" style="margin-left: 10px">重 置</el-button>
					<el-button icon="upload" style="margin-left: 10px">导 出</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="gva-table-box">
			<el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID">
				<el-table-column type="selection" width="55" />
				<el-table-column align="left" label="部门" prop="" />
				<el-table-column align="left" label="岗位" prop="" />
				<el-table-column align="left" label="姓名" prop="" />
				<el-table-column align="left" label="入职日期" width="100" prop="" />
				<el-table-column align="left" label="职级" prop="" />
				<el-table-column align="left" label="等级工资" width="100" prop="" />
				<el-table-column align="center" label="年度应发汇总">
					<el-table-column align="center" label="应发固定工资">
						<el-table-column label="1月"></el-table-column>
						<el-table-column label="2月"></el-table-column>
						<el-table-column label="3月"></el-table-column>
						<el-table-column label="4月"></el-table-column>
						<el-table-column label="5月"></el-table-column>
						<el-table-column label="6月"></el-table-column>
						<el-table-column label="7月"></el-table-column>
						<el-table-column label="8月"></el-table-column>
						<el-table-column label="9月"></el-table-column>
						<el-table-column label="10月"></el-table-column>
						<el-table-column label="11月"></el-table-column>
						<el-table-column label="12月"></el-table-column>
						<el-table-column label="合计"></el-table-column>
					</el-table-column>
					<el-table-column align="center" label="应发提成奖金">
						<el-table-column label="1月"></el-table-column>
						<el-table-column label="2月"></el-table-column>
						<el-table-column label="3月"></el-table-column>
						<el-table-column label="4月"></el-table-column>
						<el-table-column label="5月"></el-table-column>
						<el-table-column label="6月"></el-table-column>
						<el-table-column label="7月"></el-table-column>
						<el-table-column label="8月"></el-table-column>
						<el-table-column label="9月"></el-table-column>
						<el-table-column label="10月"></el-table-column>
						<el-table-column label="11月"></el-table-column>
						<el-table-column label="12月"></el-table-column>
						<el-table-column label="合计"></el-table-column>
					</el-table-column>
					<el-table-column label="月度提成补发" width="150"></el-table-column>
					<el-table-column label="年度提成" width="100"></el-table-column>
					<el-table-column label="子品牌经营绩效奖励">
						<el-table-column label="半年奖"></el-table-column>
						<el-table-column label="年度增量奖" width="100"></el-table-column>
						<el-table-column label="年度利润奖" width="100"></el-table-column>
					</el-table-column>
					<el-table-column label="年收入2024年应发" width="200"></el-table-column>
					<el-table-column label="年收入2023年应发" width="200"></el-table-column>
					<el-table-column label="同比增长" width="100"></el-table-column>
					<el-table-column align="center" label="10号实际发放">
						<el-table-column label="1月"></el-table-column>
						<el-table-column label="2月"></el-table-column>
						<el-table-column label="3月"></el-table-column>
						<el-table-column label="4月"></el-table-column>
						<el-table-column label="5月"></el-table-column>
						<el-table-column label="6月"></el-table-column>
						<el-table-column label="7月"></el-table-column>
						<el-table-column label="8月"></el-table-column>
						<el-table-column label="9月"></el-table-column>
						<el-table-column label="10月"></el-table-column>
						<el-table-column label="11月"></el-table-column>
						<el-table-column label="12月"></el-table-column>
						<el-table-column label="合计"></el-table-column>
					</el-table-column>
					<el-table-column align="center" label="20-25号实际发放">
						<el-table-column label="1月"></el-table-column>
						<el-table-column label="2月"></el-table-column>
						<el-table-column label="3月"></el-table-column>
						<el-table-column label="4月"></el-table-column>
						<el-table-column label="5月"></el-table-column>
						<el-table-column label="6月"></el-table-column>
						<el-table-column label="7月"></el-table-column>
						<el-table-column label="8月"></el-table-column>
						<el-table-column label="9月"></el-table-column>
						<el-table-column label="10月"></el-table-column>
						<el-table-column label="11月"></el-table-column>
						<el-table-column label="12月"></el-table-column>
						<el-table-column label="合计"></el-table-column>
					</el-table-column>
          <el-table-column label="半年奖实发" width="150"></el-table-column>
          <el-table-column label="年度提成及年终奖实发" width="200"></el-table-column>
          <el-table-column align="center" label="节日金汇总（单独发放）">
            <el-table-column label="元旦"></el-table-column>
            <el-table-column label="春节"></el-table-column>
            <el-table-column label="开工利是" width="100"></el-table-column>
            <el-table-column label="端午"></el-table-column>
            <el-table-column label="中秋"></el-table-column>
            <el-table-column label="国庆"></el-table-column>
          </el-table-column>
          <el-table-column label="年度实发总计" width="150"></el-table-column>
				</el-table-column>
			</el-table>
		</div>
	</div>
</template>

<script setup>
import { ref } from 'vue'

const searchInfo = ref({})

const tableData = ref([])
</script>
