<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">事故登记</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增登记
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="事故类型">
          <el-select v-model="searchForm.type" placeholder="请选择事故类型" clearable>
            <el-option label="跌倒" :value="1" />
            <el-option label="摔伤" :value="2" />
            <el-option label="身体不适" :value="3" />
            <el-option label="其他" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="处理中" :value="0" />
            <el-option label="已处理" :value="1" />
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
        <el-table-column prop="accident_time" label="事故时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.accident_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="location" label="发生地点" width="150" />
        <el-table-column prop="type" label="事故类型" width="100">
          <template #default="scope">
            {{ getTypeText(scope.row.type) }}
          </template>
        </el-table-column>
        <el-table-column prop="title" label="事故标题" width="200" show-overflow-tooltip />
        <el-table-column prop="handler" label="处理人" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 0 ? 'warning' : 'success'">
              {{ scope.row.status === 0 ? '处理中' : '已处理' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleView(scope.row)">详情</el-button>
            <el-button type="success" link @click="handleEdit(scope.row)">编辑</el-button>
            <el-button v-if="scope.row.status === 0" type="warning" link @click="handleResolve(scope.row)">处理</el-button>
            <el-button type="danger" link @click="handleDelete(scope.row)">删除</el-button>
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
  { id: 1, elder_name: '王阿姨', accident_time: '2024-05-02 08:30:00', location: '走廊', type: 1, title: '在走廊行走时不慎跌倒', handler: '张护士', status: 0 },
  { id: 2, elder_name: '李叔叔', accident_time: '2024-05-01 15:00:00', location: '餐厅', type: 3, title: '用餐时突发头晕', handler: '王医生', status: 1 }
])

const searchForm = reactive({
  elder_name: '',
  type: null,
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 2
})

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const getTypeText = (type) => {
  const texts = { 1: '跌倒', 2: '摔伤', 3: '身体不适', 4: '其他' }
  return texts[type] || '未知'
}

const handleAdd = () => {
  ElMessage.info('新增登记功能开发中')
}

const handleView = (row) => {
  ElMessage.info('查看详情功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleResolve = (row) => {
  ElMessage.success('已标记为已处理')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该记录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('删除成功')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
    }
  }
}

const handleSearch = () => {
  pagination.page = 1
}

const handleReset = () => {
  searchForm.elder_name = ''
  searchForm.type = null
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
