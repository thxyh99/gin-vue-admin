<template>
	<div class="staffNumber-line-box">
		<div ref="echart" class="staffNumber-line" />
	</div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import { useWindowResize } from '@/hooks/use-windows-resize'

var dataAxis = []
for (var i = 1; i < 13; i++) {
	dataAxis.push(`${i}月`)
}
var data = [220, 182, 191, 234, 290, 330, 310, 123, 442, 321, 90, 149]
var yMax = 500
var dataShadow = []

for (let i = 0; i < data.length; i++) {
	dataShadow.push(yMax)
}

let chart = null
const echart = ref(null)

useWindowResize(() => {
	if (!chart) {
		return
	}
	chart.resize()
})

const initChart = () => {
	if (chart) {
		chart = null
	}
	chart = echarts.init(echart.value)
	setOptions()
}
const setOptions = () => {
	chart.setOption({
    title: {
      text: '离职人数',
    },
    tooltip: {
      show: true,
      trigger: 'axis'
    },
		xAxis: {
			type: 'category',
			data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
		},
		yAxis: {
			type: 'value',
		},
		series: [
			{
				data: [120, 200, 150, 80, 70, 110, 130],
				type: 'bar',
			},
		],
	})
}

onMounted(async () => {
	await nextTick()
	initChart()
})

onUnmounted(() => {
	if (!chart) {
		return
	}
	chart.dispose()
	chart = null
})
</script>
<style lang="scss" scoped>
.staffNumber-line-box {
	.staffNumber-line {
		background-color: #fff;
		height: 360px;
		width: 100%;
	}
	.staffNumber-line-title {
		font-weight: 600;
		margin-bottom: 12px;
	}
}
</style>
