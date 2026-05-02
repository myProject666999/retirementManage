<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="stats-card">
          <div class="stats-header">
            <span class="stats-title">今日新增咨询</span>
            <el-icon class="stats-icon"><ChatDotRound /></el-icon>
          </div>
          <div class="stats-value">{{ statsData.todayConsultation || 0 }}</div>
          <div class="stats-trend">
            较昨日 <span class="trend-up">+12%</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stats-card">
          <div class="stats-header">
            <span class="stats-title">今日新增预定</span>
            <el-icon class="stats-icon"><Calendar /></el-icon>
          </div>
          <div class="stats-value">{{ statsData.todayReservation || 0 }}</div>
          <div class="stats-trend">
            较昨日 <span class="trend-up">+8%</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stats-card">
          <div class="stats-header">
            <span class="stats-title">今日新增合同</span>
            <el-icon class="stats-icon"><Document /></el-icon>
          </div>
          <div class="stats-value">{{ statsData.todayContract || 0 }}</div>
          <div class="stats-trend">
            较昨日 <span class="trend-down">-3%</span>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="stats-card">
          <div class="stats-header">
            <span class="stats-title">可售床位</span>
            <el-icon class="stats-icon"><Bed /></el-icon>
          </div>
          <div class="stats-value">{{ statsData.availableBeds || 0 }}</div>
          <div class="stats-trend">
            总床位: {{ statsData.totalBeds || 0 }}
          </div>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="20">
      <el-col :span="12">
        <div class="chart-card">
          <div class="chart-title">业务趋势</div>
          <div ref="trendChartRef" class="chart-container"></div>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="chart-card">
          <div class="chart-title">本月业绩排行</div>
          <div ref="rankingChartRef" class="chart-container"></div>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="20">
      <el-col :span="12">
        <div class="chart-card">
          <div class="chart-title">合同到期提醒</div>
          <el-table :data="expiringContracts" v-loading="loading" border stripe>
            <el-table-column prop="contract_no" label="合同编号" width="180" />
            <el-table-column prop="elder_name" label="长者姓名" width="120" />
            <el-table-column prop="end_date" label="到期日期" width="150">
              <template #default="scope">
                {{ formatDate(scope.row.end_date) }}
              </template>
            </el-table-column>
            <el-table-column prop="days_left" label="剩余天数" width="120">
              <template #default="scope">
                <el-tag :type="scope.row.days_left <= 7 ? 'danger' : 'warning'">
                  {{ scope.row.days_left }} 天
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="chart-card">
          <div class="chart-title">最近咨询记录</div>
          <el-table :data="recentConsultations" v-loading="loading" border stripe>
            <el-table-column prop="name" label="姓名" width="120" />
            <el-table-column prop="phone" label="电话" width="130" />
            <el-table-column prop="channel_name" label="来源渠道" width="120" />
            <el-table-column prop="consult_time" label="咨询时间" width="180">
              <template #default="scope">
                {{ formatDate(scope.row.consult_time) }}
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, nextTick } from 'vue'
import * as echarts from 'echarts'

const loading = ref(false)
const trendChartRef = ref(null)
const rankingChartRef = ref(null)

const statsData = reactive({
  todayConsultation: 15,
  todayReservation: 8,
  todayContract: 3,
  availableBeds: 45,
  totalBeds: 200
})

const expiringContracts = ref([
  { contract_no: 'HT202401001', elder_name: '张三', end_date: '2024-05-10', days_left: 8 },
  { contract_no: 'HT202401002', elder_name: '李四', end_date: '2024-05-05', days_left: 3 },
  { contract_no: 'HT202401003', elder_name: '王五', end_date: '2024-05-08', days_left: 6 },
  { contract_no: 'HT202401004', elder_name: '赵六', end_date: '2024-05-03', days_left: 1 }
])

const recentConsultations = ref([
  { name: '陈阿姨', phone: '13800138001', channel_name: '网络推广', consult_time: '2024-05-01 10:30:00' },
  { name: '刘叔叔', phone: '13800138002', channel_name: '口碑介绍', consult_time: '2024-05-01 14:20:00' },
  { name: '周阿姨', phone: '13800138003', channel_name: '社区活动', consult_time: '2024-05-02 09:15:00' },
  { name: '吴叔叔', phone: '13800138004', channel_name: '电话咨询', consult_time: '2024-05-02 11:00:00' }
])

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const initTrendChart = () => {
  if (!trendChartRef.value) return
  
  const chart = echarts.init(trendChartRef.value)
  const option = {
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['咨询量', '预定量', '签约量']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: ['1月', '2月', '3月', '4月', '5月', '6月', '7月']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '咨询量',
        type: 'line',
        smooth: true,
        data: [120, 132, 101, 134, 90, 230, 210]
      },
      {
        name: '预定量',
        type: 'line',
        smooth: true,
        data: [22, 18, 25, 33, 28, 45, 52]
      },
      {
        name: '签约量',
        type: 'line',
        smooth: true,
        data: [15, 12, 18, 25, 20, 35, 40]
      }
    ]
  }
  chart.setOption(option)
  
  window.addEventListener('resize', () => {
    chart.resize()
  })
}

const initRankingChart = () => {
  if (!rankingChartRef.value) return
  
  const chart = echarts.init(rankingChartRef.value)
  const option = {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'value'
    },
    yAxis: {
      type: 'category',
      data: ['周经理', '李主管', '王销售', '张顾问', '陈专员']
    },
    series: [
      {
        name: '业绩金额',
        type: 'bar',
        data: [125000, 98000, 85000, 72000, 65000],
        itemStyle: {
          color: function(params) {
            const colors = ['#67c23a', '#409eff', '#e6a23c', '#909399', '#f56c6c']
            return colors[params.dataIndex]
          }
        }
      }
    ]
  }
  chart.setOption(option)
  
  window.addEventListener('resize', () => {
    chart.resize()
  })
}

onMounted(() => {
  nextTick(() => {
    initTrendChart()
    initRankingChart()
  })
})
</script>
