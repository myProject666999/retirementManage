<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">护理等级</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增等级
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="等级名称">
          <el-input v-model="searchForm.name" placeholder="请输入等级名称" clearable />
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
        <el-table-column prop="name" label="等级名称" width="150" />
        <el-table-column prop="code" label="等级编码" width="150" />
        <el-table-column prop="level" label="等级级别" width="120">
          <template #default="scope">
            <el-tag type="primary">{{ scope.row.level }}级</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="护理费用" width="150">
          <template #default="scope">
            ¥{{ scope.row.price }} / 月
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" width="300" show-overflow-tooltip />
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
  { id: 1, name: '一级护理', code: 'CARE01', level: 1, price: 2000, status: 1, description: '适合自理能力较强的长者，提供基础生活照料服务' },
  { id: 2, name: '二级护理', code: 'CARE02', level: 2, price: 3500, status: 1, description: '适合半自理长者，提供协助生活照料和基础医疗护理' },
  { id: 3, name: '三级护理', code: 'CARE03', level: 3, price: 5000, status: 1, description: '适合失能长者，提供24小时全天候护理服务' },
  { id: 4, name: '特级护理', code: 'CARE04', level: 4, price: 8000, status: 1, description: '适合重症、临终关怀等特殊需求长者，提供专业医疗级护理' }
])

const searchForm = reactive({
  name: '',
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 4
})

const handleAdd = () => {
  ElMessage.info('新增等级功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该等级吗？', '提示', {
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
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
