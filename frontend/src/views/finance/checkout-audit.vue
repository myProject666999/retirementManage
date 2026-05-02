<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">退住审核</span>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待审核" :value="0" />
            <el-option label="审核通过" :value="1" />
            <el-option label="审核拒绝" :value="2" />
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
        <el-table-column prop="elder_name" label="长者姓名" width="120" />
        <el-table-column prop="contract_no" label="合同编号" width="180" />
        <el-table-column prop="checkout_date" label="申请退住日期" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.checkout_date) }}
          </template>
        </el-table-column>
        <el-table-column prop="reason" label="退住原因" width="250" show-overflow-tooltip />
        <el-table-column prop="refund_amount" label="退款金额" width="120">
          <template #default="scope">
            ¥{{ scope.row.refund_amount }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="audit_time" label="审核时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.audit_time) || '-' }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleView(scope.row)">详情</el-button>
            <el-button v-if="scope.row.status === 0" type="success" link @click="handleApprove(scope.row)">通过</el-button>
            <el-button v-if="scope.row.status === 0" type="danger" link @click="handleReject(scope.row)">拒绝</el-button>
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
import { ElMessage, ElMessageBox, ElInput } from 'element-plus'

const loading = ref(false)

const tableData = ref([
  { id: 1, elder_name: '张阿姨', contract_no: 'HT202312003', checkout_date: '2024-05-10', reason: '子女接回家照顾', refund_amount: 5000, status: 0, audit_time: null },
  { id: 2, elder_name: '刘叔叔', contract_no: 'HT202401005', checkout_date: '2024-04-20', reason: '身体原因需要专业医院治疗', refund_amount: 8000, status: 1, audit_time: '2024-04-21 10:30:00' }
])

const searchForm = reactive({
  elder_name: '',
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 2
})

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

const getStatusText = (status) => {
  const texts = { 0: '待审核', 1: '审核通过', 2: '审核拒绝' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'success', 2: 'danger' }
  return types[status] || ''
}

const handleView = (row) => {
  ElMessage.info('查看详情功能开发中')
}

const handleApprove = async (row) => {
  try {
    await ElMessageBox.confirm('确定要通过该退住申请吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('已通过审核')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('审核失败:', error)
    }
  }
}

const handleReject = async (row) => {
  try {
    await ElMessageBox.prompt('请输入拒绝原因：', '拒绝申请', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /.+/,
      inputErrorMessage: '请输入拒绝原因'
    })
    ElMessage.success('已拒绝申请')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('拒绝失败:', error)
    }
  }
}

const handleSearch = () => {
  pagination.page = 1
}

const handleReset = () => {
  searchForm.elder_name = ''
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
