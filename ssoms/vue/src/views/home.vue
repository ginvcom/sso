<template>
<div>
  <a-row :gutter="32">
    <a-col :span="16">
      <h1 class="home__header">菜单 & 授权统计</h1>
      <div ref="barChartContainer" style="height: 240px"></div>
    </a-col>
    <a-col :span="8">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-statistic class="home__statistic" title="角色" :value="4">
            <template #suffix>
              <team-outlined class="blue" />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="12">
          <a-statistic class="home__statistic" title="用户" :value="1">
            <template #suffix>
              <user-outlined class="yellow" />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="12">
          <a-statistic class="home__statistic" title="系统" :value="5">
            <template #suffix>
              <appstore-outlined class="purple" />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="12">
          <a-statistic class="home__statistic" title="菜单" :value="30">
            <template #suffix>
              <unordered-list-outlined class="pink" />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="12">
          <a-statistic class="home__statistic" title="操作" :value="1128">
            <template #suffix>
              <calculator-outlined class="orange" />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="12">
          <a-statistic class="home__statistic" title="授权" :value="523">
            <template #suffix>
              <audit-outlined class="green" />
            </template>
          </a-statistic>
        </a-col>
      </a-row>
    </a-col>
    <a-col :span="24">
      <a-card title="已授权系统" class="home__systems">
        <div v-for="(obj, index) in state.systems" :key="obj.uuid" class="object__item" :class="`col-${index % 4 + 1}`">
          <div
            class="object__current-value object__option"
            @click="joinSystem(obj)"
          >
            <a-image
              :width="48"
              :height="48"
              :src="ossConfig.ginvdoc.domain + obj.icon"
              :preview="false"
              fallback="data:image/png;base64,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"
            />
            <p class="object__current-system-name">{{obj.objectName}}</p>
          </div>
        </div>
      </a-card>
    </a-col>
  </a-row>
</div>
</template>
<script setup lang="ts">
import { onBeforeMount, onMounted, reactive, ref } from 'vue'
import { Column } from '@antv/g2plot'
import {
  TeamOutlined,
  UserOutlined,
  AppstoreOutlined,
  UnorderedListOutlined,
  AuditOutlined,
  CalculatorOutlined
} from '@ant-design/icons-vue'
// Object 是js的关键字, 别名处理一下
import {
  roleOperations,
  userDetail,
  addUser,
  updateUser,
  objectList,
  ObjectListReply,
  RoleOperationsReqParams,
  ObjectOption,
  RoleOperationsReply,
  UserForm,
  Object as Obj
} from '../api/ssoms'
import { ossConfig } from '@/config'

interface System {
  name: string
  icon: string
}

interface State {
  loading: boolean
  systems: Array<Obj>
}

const state = reactive<State>({
  loading: false,
  systems: []
})

const barChartContainer = ref()

onBeforeMount(() => {
  getSystemOptions()
})

const getSystemOptions = () => {
  objectList({ topKey: '', key: '', objectName: '' }).then((data: ObjectListReply) => {
    state.systems = data.list
  }).finally(() => {
    state.loading = false
  })
}

const joinSystem = (obj: Obj) => {
  // 进入系统
}

onMounted(() => {
  initChart()
})

const initChart = () => {
  fetch('https://gw.alipayobjects.com/os/antfincdn/PC3daFYjNw/column-data.json')
  .then((data) => data.json())
  .then((data) => {
    const column = new Column(barChartContainer.value, {
      data,
      theme: {
        colors10: ['#1890ff', '#fa541c', '#41ca9d', '#ed8936']
      },
      maxColumnWidth: 10,
      xField: 'city',
      yField: 'value',
      seriesField: 'type',
      isGroup: true,
      columnStyle: {
        radius: [2, 2, 0, 0],
      },
    })
    column.render()
  })
}
</script>
