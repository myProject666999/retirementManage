<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">来访登记</span>
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
        <el-form-item label="访客姓名">
          <el-input v-model="searchForm.visitor_name" placeholder="请输入访客姓名" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="访问中" :value="0" />
            <el-option label="已离开" :value="1" />
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
        <el-table-column prop="visitor_name" label="访客姓名" width="120" />
        <el-table-column prop="relation" label="关系" width="100" />
        <el-table-column prop="phone" label="联系电话" width="150" />
        <el-table-column prop="id_card" label="身份证号" width="180" />
        <el-table-column prop="visit_time" label="访问时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.visit_time) }}
          </template>
        </el-table-column>
        <el-table-column prop="leave_time" label="离开时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.leave_time) || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 0 ? 'warning' : 'success'">
              {{ scope.row.status === 0 ? '访问中' : '已离开' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button v-if="scope.row.status === 0" type="success" link @click="handleLeave(scope.row)">登记离开</el-button>
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
  { id: 1, elder_name: '王阿姨', visitor_name: '王小明', relation: '儿子', phone: '13800138001', id_card: '110101198001011234', visit_time: '2024-05-02 10:00:00', leave_time: null, status: 0 },
  { id: 2, elder_name: '李叔叔', visitor_name: '李小红', relation: '女儿', phone: '13800138002', id_card: '110101198501015678', visit_time: '2024-05-01 14:00:00', leave_time: '2024-05-01 17:00:00', status: 1 }
])

const searchForm = reactive({
  elder_name: '',
  visitor_name: '',
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

const handleLeave = (row) => {
  ElMessage.success('已登记离开')
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
  searchForm.visitor_name = ''
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
