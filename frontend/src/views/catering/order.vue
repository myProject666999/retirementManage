<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">点餐管理</span>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="订单编号">
          <el-input v-model="searchForm.order_no" placeholder="请输入订单编号" clearable />
        </el-form-item>
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待确认" :value="0" />
            <el-option label="已确认" :value="1" />
            <el-option label="制作中" :value="2" />
            <el-option label="配送中" :value="3" />
            <el-option label="已完成" :value="4" />
            <el-option label="已取消" :value="5" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="table-container">
      <el-table :data="tableData" v-loading="loading" border stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="order_no" label="订单编号" width="180" />
        <el-table-column prop="elder_name" label="长者姓名" width="120" />
        <el-table-column prop="order_time" label="下单时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.order_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="delivery_place" label="配送地点" width="150" />
        <el-table-column prop="total_amount" label="订单金额" width="120">
          <template #default="scope">
            ¥{{ scope.row.total_amount }}
          </template>
        </el-table-column>
        <el-table-column prop="paid_amount" label="实付金额" width="120">
          <template #default="scope">
            ¥{{ scope.row.paid_amount }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleView(scope.row)">详情</el-button>
            <el-button v-if="scope.row.status === 0" type="success" link @click="handleConfirm(scope.row)">确认</el-button>
            <el-button v-if="scope.row.status === 1" type="warning" link @click="handleStart(scope.row)">开始制作</el-button>
            <el-button v-if="scope.row.status === 2" type="info" link @click="handleDelivery(scope.row)">开始配送</el-button>
            <el-button v-if="scope.row.status === 3" type="success" link @click="handleComplete(scope.row)">完成</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <div class="pagination-container">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.page_size"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)

const tableData = ref([
  { id: 1, order_no: 'ORD202405020001', elder_name: '王阿姨', order_time: '2024-05-02 11:30:00', delivery_place: '1号楼101房间', total_amount: 45, paid_amount: 45, status: 0 },
  { id: 2, order_no: 'ORD202405020002', elder_name: '李叔叔', order_time: '2024-05-02 11:00:00', delivery_place: '1号楼102房间', total_amount: 30, paid_amount: 30, status: 2 },
  { id: 3, order_no: 'ORD202405020003', elder_name: '张奶奶', order_time: '2024-05-02 08:00:00', delivery_place: '2号楼201房间', total_amount: 15, paid_amount: 15, status: 4 }
])

const searchForm = reactive({
  order_no: '',
  elder_name: '',
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 3
})

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const getStatusText = (status) => {
  const texts = { 0: '待确认', 1: '已确认', 2: '制作中', 3: '配送中', 4: '已完成', 5: '已取消' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'primary', 2: 'info', 3: 'warning', 4: 'success', 5: 'info' }
  return types[status] || ''
}

const handleView = (row) => {
  ElMessage.info('查看订单详情功能开发中')
}

const handleConfirm = (row) => {
  ElMessage.success('已确认订单')
}

const handleStart = (row) => {
  ElMessage.success('已开始制作')
}

const handleDelivery = (row) => {
  ElMessage.success('已开始配送')
}

const handleComplete = (row) => {
  ElMessage.success('已完成订单')
}

const handleSearch = () => {
  pagination.page = 1
}

const handleReset = () => {
  searchForm.order_no = ''
  searchForm.elder_name = ''
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
