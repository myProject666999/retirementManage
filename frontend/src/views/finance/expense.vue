<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">消费记录</span>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="记录单号">
          <el-input v-model="searchForm.record_no" placeholder="请输入记录单号" clearable />
        </el-form-item>
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="消费类型">
          <el-select v-model="searchForm.type" placeholder="请选择类型" clearable>
            <el-option label="护理费" :value="1" />
            <el-option label="伙食费" :value="2" />
            <el-option label="服务费" :value="3" />
            <el-option label="其他费用" :value="4" />
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
        <el-table-column prop="record_no" label="记录单号" width="180" />
        <el-table-column prop="elder_name" label="长者姓名" width="120" />
        <el-table-column prop="type" label="消费类型" width="100">
          <template #default="scope">
            <el-tag>{{ getTypeText(scope.row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="消费标题" width="200" />
        <el-table-column prop="amount" label="消费金额" width="120">
          <template #default="scope">
            <span style="color: #f56c6c; font-weight: bold;">-¥{{ scope.row.amount }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="balance" label="账户余额" width="120">
          <template #default="scope">
            ¥{{ scope.row.balance }}
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" width="250" show-overflow-tooltip />
        <el-table-column prop="created_at" label="消费时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
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
import { ElMessage } from 'element-plus'

const loading = ref(false)

const tableData = ref([
  { id: 1, record_no: 'EXP202405020001', elder_name: '王阿姨', type: 1, title: '5月份护理费', amount: 3500, balance: 2000, description: '2024年5月份护理费用扣除', created_at: '2024-05-01 00:00:00' },
  { id: 2, record_no: 'EXP202405020002', elder_name: '李叔叔', type: 2, title: '午餐点餐', amount: 45, balance: 8500, description: '午餐点餐消费', created_at: '2024-05-02 12:30:00' },
  { id: 3, record_no: 'EXP202405020003', elder_name: '张奶奶', type: 3, title: '康复训练服务', amount: 160, balance: 2500, description: '康复训练服务费用', created_at: '2024-05-02 15:00:00' }
])

const searchForm = reactive({
  record_no: '',
  elder_name: '',
  type: null
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

const getTypeText = (type) => {
  const texts = { 1: '护理费', 2: '伙食费', 3: '服务费', 4: '其他费用' }
  return texts[type] || '未知'
}

const handleSearch = () => {
  pagination.page = 1
}

const handleReset = () => {
  searchForm.record_no = ''
  searchForm.elder_name = ''
  searchForm.type = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
