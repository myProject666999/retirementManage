<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">外出登记</span>
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
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="外出中" :value="0" />
            <el-option label="已返回" :value="1" />
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
        <el-table-column prop="out_time" label="外出时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.out_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="return_time" label="返回时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.return_time) || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="companion" label="陪同人" width="120" />
        <el-table-column prop="relation" label="关系" width="100" />
        <el-table-column prop="phone" label="联系电话" width="150" />
        <el-table-column prop="purpose" label="外出目的" width="200" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 0 ? 'warning' : 'success'">
              {{ scope.row.status === 0 ? '外出中' : '已返回' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button v-if="scope.row.status === 0" type="success" link @click="handleReturn(scope.row)">登记返回</el-button>
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
  { id: 1, elder_name: '王阿姨', out_time: '2024-05-02 09:00:00', return_time: null, companion: '王小明', relation: '儿子', phone: '13800138001', purpose: '去医院检查', status: 0 },
  { id: 2, elder_name: '李叔叔', out_time: '2024-05-01 14:30:00', return_time: '2024-05-01 18:00:00', companion: '李小红', relation: '女儿', phone: '13800138002', purpose: '回家探亲', status: 1 }
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
  return new Date(date).toLocaleString()
}

const handleAdd = () => {
  ElMessage.info('新增登记功能开发中')
}

const handleReturn = (row) => {
  ElMessage.success('已登记返回')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
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
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
