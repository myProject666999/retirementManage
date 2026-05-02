<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">服务预定</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增预定
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="长者姓名">
          <el-input v-model="searchForm.elder_name" placeholder="请输入长者姓名" clearable />
        </el-form-item>
        <el-form-item label="服务项目">
          <el-select v-model="searchForm.service_item_id" placeholder="请选择服务项目" clearable>
            <el-option
              v-for="item in serviceItemList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待服务" :value="0" />
            <el-option label="服务中" :value="1" />
            <el-option label="已完成" :value="2" />
            <el-option label="已取消" :value="3" />
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
        <el-table-column prop="service_item_name" label="服务项目" width="150" />
        <el-table-column prop="service_time" label="服务时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.service_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="service_place" label="服务地点" width="150" />
        <el-table-column prop="quantity" label="数量" width="80" />
        <el-table-column prop="amount" label="金额" width="120">
          <template #default="scope">
            ¥{{ scope.row.amount }}
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
            <el-button v-if="scope.row.status === 0" type="primary" link @click="handleStart(scope.row)">开始服务</el-button>
            <el-button v-if="scope.row.status === 1" type="success" link @click="handleComplete(scope.row)">完成服务</el-button>
            <el-button type="info" link @click="handleView(scope.row)">详情</el-button>
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
  { id: 1, elder_name: '王阿姨', service_item_name: '日常护理', service_time: '2024-05-03 09:00:00', service_place: '1号楼101房间', quantity: 1, amount: 50, status: 0 },
  { id: 2, elder_name: '李叔叔', service_item_name: '康复训练', service_time: '2024-05-02 14:00:00', service_place: '康复室', quantity: 2, amount: 160, status: 1 },
  { id: 3, elder_name: '张奶奶', service_item_name: '健康检查', service_time: '2024-05-01 10:00:00', service_place: '医疗室', quantity: 1, amount: 100, status: 2 }
])

const serviceItemList = ref([
  { id: 1, name: '日常护理' },
  { id: 2, name: '康复训练' },
  { id: 3, name: '健康检查' },
  { id: 4, name: '娱乐活动' }
])

const searchForm = reactive({
  elder_name: '',
  service_item_id: null,
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
  const texts = { 0: '待服务', 1: '服务中', 2: '已完成', 3: '已取消' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'primary', 2: 'success', 3: 'info' }
  return types[status] || ''
}

const handleAdd = () => {
  ElMessage.info('新增预定功能开发中')
}

const handleView = (row) => {
  ElMessage.info('查看详情功能开发中')
}

const handleStart = (row) => {
  ElMessage.success('已开始服务')
}

const handleComplete = (row) => {
  ElMessage.success('已完成服务')
}

const handleCancel = async (row) => {
  try {
    await ElMessageBox.confirm('确定要取消该预定吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('已取消预定')
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
  searchForm.elder_name = ''
  searchForm.service_item_id = null
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
