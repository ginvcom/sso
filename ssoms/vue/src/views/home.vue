<template>
<div>
  <a-row :gutter="32">
    <a-col :span="16">
      <div ref="barChartContainer"></div>
    </a-col>
    <a-col :span="8">

    </a-col>
  </a-row>
</div>
</template>
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Column } from '@antv/g2plot'

const barChartContainer = ref()

onMounted(() => {
  initChart()
})

const initChart = () => {
  fetch('https://gw.alipayobjects.com/os/antfincdn/PC3daFYjNw/column-data.json')
  .then((data) => data.json())
  .then((data) => {
    const column = new Column(barChartContainer.value, {
      data,
      xField: 'city',
      yField: 'value',
      seriesField: 'type',
      isGroup: true,
      columnStyle: {
        radius: [20, 20, 0, 0],
      },
    })
    column.render()
  })
}
</script>
