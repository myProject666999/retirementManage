<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">入住签约</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增签约
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="合同编号">
          <el-input v-model="searchForm.contract_no" placeholder="请输入合同编号" clearable />
        </el-form-item>
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待生效" :value="0" />
            <el-option label="生效中" :value="1" />
            <el-option label="已到期" :value="2" />
            <el-option label="已终止" :value="3" />
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
        <el-table-column prop="contract_no" label="合同编号" width="180" />
        <el-table-column prop="elder_name" label="长者姓名" width="120" />
        <el-table-column prop="room_info" label="入住房间" width="150" />
        <el-table-column prop="start_date" label="开始日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.start_date) }}
          </template>
        </el-table-column>
        <el-table-column prop="end_date" label="结束日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.end_date) }}
          </template>
        </el-table-column>
        <el-table-column prop="amount" label="合同金额" width="120">
          <template #default="scope">
            ¥{{ scope.row.amount }}
          </template>
        </el-table-column>
        <el-table-column prop="deposit" label="押金" width="100">
          <template #default="scope">
            ¥{{ scope.row.deposit }}
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
            <el-button type="primary" link @click="handleView(scope.row)">查看</el-button>
            <el-button type="success" link @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="warning" link @click="handleTerminate(scope.row)">终止</el-button>
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
  { id: 1, contract_no: 'HT202405001', elder_name: '王阿姨', room_info: '1号楼101-1', start_date: '2024-05-01', end_date: '2025-04-30', amount: 36000, deposit: 5000, status: 1 },
  { id: 2, contract_no: 'HT202405002', elder_name: '李叔叔', room_info: '1号楼102-1', start_date: '2024-04-15', end_date: '2025-04-14', amount: 48000, deposit: 8000, status: 1 },
  { id: 3, contract_no: 'HT202405003', elder_name: '张阿姨', room_info: '2号楼201-2', start_date: '2023-12-01', end_date: '2024-05-31', amount: 24000, deposit: 3000, status: 0 }
])

const searchForm = reactive({
  contract_no: '',
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
  return new Date(date).toLocaleDateString()
}

const getStatusText = (status) => {
  const texts = { 0: '待生效', 1: '生效中', 2: '已到期', 3: '已终止' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'success', 2: 'info', 3: 'danger' }
  return types[status] || ''
}

const handleAdd = () => {
  ElMessage.info('新增签约功能开发中')
}

const handleView = (row) => {
  ElMessage.info('查看合同详情功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑合同功能开发中')
}

const handleTerminate = async (row) => {
  try {
    await ElMessageBox.confirm('确定要终止该合同吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('合同终止成功')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('终止合同失败:', error)
    }
  }
}

const handleSearch = () => {
  pagination.page = 1
}

const handleReset = () => {
  searchForm.contract_no = ''
  searchForm.elder_name = ''
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
