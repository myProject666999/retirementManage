<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">餐饮套餐</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增套餐
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="套餐名称">
          <el-input v-model="searchForm.name" placeholder="请输入套餐名称" clearable />
        </el-form-item>
        <el-form-item label="套餐类型">
          <el-select v-model="searchForm.type" placeholder="请选择类型" clearable>
            <el-option label="单人餐" :value="1" />
            <el-option label="双人餐" :value="2" />
            <el-option label="家庭餐" :value="3" />
            <el-option label="其他" :value="4" />
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
        <el-table-column prop="name" label="套餐名称" width="150" />
        <el-table-column prop="code" label="套餐编码" width="150" />
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
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '上架' : '下架' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" width="250" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button type="info" link @click="handleView(scope.row)">查看菜品</el-button>
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
  { id: 1, name: '营养早餐套餐', code: 'PKG001', type: 1, price: 15, status: 1, description: '包含小米粥、鸡蛋、包子' },
  { id: 2, name: '商务午餐套餐', code: 'PKG002', type: 1, price: 45, status: 1, description: '包含主菜、配菜、汤品' },
  { id: 3, name: '家庭晚餐套餐', code: 'PKG003', type: 3, price: 88, status: 1, description: '适合3-4人用餐的家庭套餐' }
])

const searchForm = reactive({
  name: '',
  type: null,
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 3
})

const getTypeText = (type) => {
  const texts = { 1: '单人餐', 2: '双人餐', 3: '家庭餐', 4: '其他' }
  return texts[type] || '未知'
}

const handleAdd = () => {
  ElMessage.info('新增套餐功能开发中')
}

const handleView = (row) => {
  ElMessage.info('查看套餐菜品功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该套餐吗？', '提示', {
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
