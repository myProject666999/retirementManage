<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">菜品管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增菜品
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="菜品名称">
          <el-input v-model="searchForm.name" placeholder="请输入菜品名称" clearable />
        </el-form-item>
        <el-form-item label="菜品类型">
          <el-select v-model="searchForm.type" placeholder="请选择类型" clearable>
            <el-option label="早餐" :value="1" />
            <el-option label="午餐" :value="2" />
            <el-option label="晚餐" :value="3" />
            <el-option label="点心" :value="4" />
            <el-option label="其他" :value="5" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="上架" :value="1" />
            <el-option label="下架" :value="0" />
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
        <el-table-column prop="name" label="菜品名称" width="150" />
        <el-table-column prop="code" label="菜品编码" width="150" />
        <el-table-column prop="type" label="类型" width="100">
          <template #default="scope">
            <el-tag>{{ getTypeText(scope.row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="价格" width="120">
          <template #default="scope">
            ¥{{ scope.row.price }}
          </template>
        </el-table-column>
        <el-table-column prop="unit" label="单位" width="80" />
        <el-table-column prop="is_hot" label="热销" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.is_hot === 1" type="danger">热销</el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '上架' : '下架' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" width="200" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleEdit(scope.row)">编辑</el-button>
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
  { id: 1, name: '小米粥', code: 'DISH001', type: 1, price: 5, unit: '碗', is_hot: 0, status: 1, description: '营养早餐小米粥' },
  { id: 2, name: '鸡蛋', code: 'DISH002', type: 1, price: 2, unit: '个', is_hot: 1, status: 1, description: '新鲜煮鸡蛋' },
  { id: 3, name: '红烧肉', code: 'DISH003', type: 2, price: 25, unit: '份', is_hot: 1, status: 1, description: '精选五花肉红烧' },
  { id: 4, name: '清蒸鱼', code: 'DISH004', type: 3, price: 35, unit: '条', is_hot: 0, status: 1, description: '新鲜活鱼清蒸' },
  { id: 5, name: '水果拼盘', code: 'DISH005', type: 4, price: 15, unit: '份', is_hot: 0, status: 1, description: '时令新鲜水果' }
])

const searchForm = reactive({
  name: '',
  type: null,
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 5
})

const getTypeText = (type) => {
  const texts = { 1: '早餐', 2: '午餐', 3: '晚餐', 4: '点心', 5: '其他' }
  return texts[type] || '未知'
}

const handleAdd = () => {
  ElMessage.info('新增菜品功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该菜品吗？', '提示', {
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
  searchForm.name = ''
  searchForm.type = null
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
