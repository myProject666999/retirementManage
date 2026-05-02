<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">预存充值</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增充值
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="充值单号">
          <el-input v-model="searchForm.recharge_no" placeholder="请输入充值单号" clearable />
        </el-form-item>
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="支付方式">
          <el-select v-model="searchForm.pay_method" placeholder="请选择支付方式" clearable>
            <el-option label="现金" :value="1" />
            <el-option label="微信" :value="2" />
            <el-option label="支付宝" :value="3" />
            <el-option label="银行卡" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待支付" :value="0" />
            <el-option label="已支付" :value="1" />
            <el-option label="已取消" :value="2" />
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
        <el-table-column prop="recharge_no" label="充值单号" width="180" />
        <el-table-column prop="elder_name" label="长者姓名" width="120" />
        <el-table-column prop="amount" label="充值金额" width="120">
          <template #default="scope">
            <span style="color: #67c23a; font-weight: bold;">+¥{{ scope.row.amount }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="gift_amount" label="赠送金额" width="120">
          <template #default="scope">
            <span style="color: #e6a23c;">+¥{{ scope.row.gift_amount }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="total_amount" label="到账金额" width="120">
          <template #default="scope">
            <span style="color: #409eff; font-weight: bold;">¥{{ scope.row.total_amount }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="pay_method" label="支付方式" width="100">
          <template #default="scope">
            {{ getPayMethodText(scope.row.pay_method) }}
          </template>
        </el-table-column>
        <el-table-column prop="pay_time" label="支付时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.pay_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button v-if="scope.row.status === 0" type="success" link @click="handlePay(scope.row)">确认支付</el-button>
            <el-button v-if="scope.row.status === 0" type="danger" link @click="handleCancel(scope.row)">取消</el-button>
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
  { id: 1, recharge_no: 'RC202405020001', elder_name: '王阿姨', amount: 5000, gift_amount: 500, total_amount: 5500, pay_method: 2, pay_time: '2024-05-02 10:30:00', status: 1 },
  { id: 2, recharge_no: 'RC202405020002', elder_name: '李叔叔', amount: 10000, gift_amount: 1500, total_amount: 11500, pay_method: 3, pay_time: null, status: 0 },
  { id: 3, recharge_no: 'RC202405010001', elder_name: '张奶奶', amount: 3000, gift_amount: 200, total_amount: 3200, pay_method: 1, pay_time: '2024-05-01 14:00:00', status: 1 }
])

const searchForm = reactive({
  recharge_no: '',
  elder_name: '',
  pay_method: null,
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

const getPayMethodText = (method) => {
  const texts = { 1: '现金', 2: '微信', 3: '支付宝', 4: '银行卡' }
  return texts[method] || '未知'
}

const getStatusText = (status) => {
  const texts = { 0: '待支付', 1: '已支付', 2: '已取消' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'success', 2: 'info' }
  return types[status] || ''
}

const handleAdd = () => {
  ElMessage.info('新增充值功能开发中')
}

const handlePay = async (row) => {
  try {
    await ElMessageBox.confirm('确认该充值订单已支付吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('已确认支付')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('确认失败:', error)
    }
  }
}

const handleCancel = async (row) => {
  try {
    await ElMessageBox.confirm('确定要取消该充值订单吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('已取消订单')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('取消失败:', error)
    }
  }
}

const handleSearch = () => {
  pagination.page = 1
}

const handleReset = () => {
  searchForm.recharge_no = ''
  searchForm.elder_name = ''
  searchForm.pay_method = null
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
