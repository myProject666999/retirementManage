<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">服务项目</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增项目
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="项目名称">
          <el-input v-model="searchForm.name" placeholder="请输入项目名称" clearable />
        </el-form-item>
        <el-form-item label="项目类型">
          <el-select v-model="searchForm.type" placeholder="请选择类型" clearable>
            <el-option label="生活照料" :value="1" />
            <el-option label="医疗护理" :value="2" />
            <el-option label="康复训练" :value="3" />
            <el-option label="娱乐活动" :value="4" />
            <el-option label="其他服务" :value="5" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
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
        <el-table-column prop="name" label="项目名称" width="200" />
        <el-table-column prop="code" label="项目编码" width="150" />
        <el-table-column prop="type" label="类型" width="120">
          <template #default="scope">
            <el-tag>{{ getTypeText(scope.row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="服务价格" width="120">
          <template #default="scope">
            ¥{{ scope.row.price }}
          </template>
        </el-table-column>
        <el-table-column prop="duration" label="服务时长" width="120">
          <template #default="scope">
            {{ scope.row.duration }} 分钟
          </template>
        </el-table-column>
        <el-table-column prop="unit" label="单位" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
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
  { id: 1, name: '日常护理', code: 'SVC001', type: 2, price: 50, duration: 30, unit: '次', status: 1, description: '包括洗漱、穿衣、喂食等日常照料服务' },
  { id: 2, name: '康复训练', code: 'SVC002', type: 3, price: 80, duration: 60, unit: '次', status: 1, description: '专业康复师进行肢体康复训练' },
  { id: 3, name: '健康检查', code: 'SVC003', type: 2, price: 100, duration: 45, unit: '次', status: 1, description: '定期健康检查，包括血压、血糖等' },
  { id: 4, name: '娱乐活动', code: 'SVC004', type: 4, price: 30, duration: 120, unit: '次', status: 1, description: '棋牌、书法、绘画等娱乐活动' }
])

const searchForm = reactive({
  name: '',
  type: null,
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 4
})

const getTypeText = (type) => {
  const texts = { 1: '生活照料', 2: '医疗护理', 3: '康复训练', 4: '娱乐活动', 5: '其他服务' }
  return texts[type] || '未知'
}

const handleAdd = () => {
  ElMessage.info('新增项目功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该项目吗？', '提示', {
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
