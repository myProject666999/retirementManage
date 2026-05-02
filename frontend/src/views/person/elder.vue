<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">长者档案</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增档案
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="姓名">
          <el-input v-model="searchForm.name" placeholder="请输入姓名" clearable />
        </el-form-item>
        <el-form-item label="电话">
          <el-input v-model="searchForm.phone" placeholder="请输入电话" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="待入住" :value="0" />
            <el-option label="已入住" :value="1" />
            <el-option label="已退住" :value="2" />
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
        <el-table-column prop="name" label="姓名" width="120" />
        <el-table-column prop="gender" label="性别" width="80">
          <template #default="scope">
            {{ scope.row.gender === 1 ? '男' : '女' }}
          </template>
        </el-table-column>
        <el-table-column prop="age" label="年龄" width="80" />
        <el-table-column prop="id_card" label="身份证号" width="180" />
        <el-table-column prop="phone" label="电话" width="150" />
        <el-table-column prop="care_level_name" label="护理等级" width="120" />
        <el-table-column prop="health_status" label="健康状态" width="120">
          <template #default="scope">
            <el-tag :type="getHealthTagType(scope.row.health_status)">
              {{ getHealthStatusText(scope.row.health_status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="入住状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleView(scope.row)">详情</el-button>
            <el-button type="success" link @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="warning" link @click="handleContact(scope.row)">联系人</el-button>
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
  { id: 1, name: '王阿姨', gender: 2, age: 75, id_card: '110101194901011234', phone: '13800138001', care_level_name: '二级护理', health_status: 2, status: 1 },
  { id: 2, name: '李叔叔', gender: 1, age: 78, id_card: '110101194605015678', phone: '13800138002', care_level_name: '三级护理', health_status: 1, status: 1 },
  { id: 3, name: '张奶奶', gender: 2, age: 82, id_card: '110101194208019012', phone: '13800138003', care_level_name: '特级护理', health_status: 3, status: 1 }
])

const searchForm = reactive({
  name: '',
  phone: '',
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 3
})

const getHealthStatusText = (status) => {
  const texts = { 1: '良好', 2: '一般', 3: '较差' }
  return texts[status] || '未知'
}

const getHealthTagType = (status) => {
  const types = { 1: 'success', 2: 'warning', 3: 'danger' }
  return types[status] || ''
}

const getStatusText = (status) => {
  const texts = { 0: '待入住', 1: '已入住', 2: '已退住' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'info', 1: 'success', 2: 'danger' }
  return types[status] || ''
}

const handleAdd = () => {
  ElMessage.info('新增档案功能开发中')
}

const handleView = (row) => {
  ElMessage.info('查看详情功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleContact = (row) => {
  ElMessage.info('联系人管理功能开发中')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该档案吗？', '提示', {
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
  searchForm.phone = ''
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
