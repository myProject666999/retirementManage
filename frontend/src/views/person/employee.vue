<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">员工管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增员工
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
        <el-form-item label="部门">
          <el-select v-model="searchForm.department_id" placeholder="请选择部门" clearable>
            <el-option
              v-for="dept in departmentList"
              :key="dept.id"
              :label="dept.name"
              :value="dept.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="在职" :value="1" />
            <el-option label="离职" :value="0" />
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
        <el-table-column prop="id_card" label="身份证号" width="180" />
        <el-table-column prop="phone" label="电话" width="150" />
        <el-table-column prop="email" label="邮箱" width="200" />
        <el-table-column prop="department_name" label="部门" width="120" />
        <el-table-column prop="position" label="职位" width="120" />
        <el-table-column prop="entry_date" label="入职日期" width="120">
          <template #default="scope">
            {{ formatDate(scope.row.entry_date) }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '在职' : '离职' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleView(scope.row)">详情</el-button>
            <el-button type="success" link @click="handleEdit(scope.row)">编辑</el-button>
            <el-button v-if="scope.row.status === 1" type="warning" link @click="handleLeave(scope.row)">离职</el-button>
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
  { id: 1, name: '张护士', gender: 2, id_card: '110101199001011234', phone: '13900139001', email: 'zhang@example.com', department_name: '护理部', position: '护士', entry_date: '2020-03-01', status: 1 },
  { id: 2, name: '李医生', gender: 1, id_card: '110101198505015678', phone: '13900139002', email: 'li@example.com', department_name: '医疗部', position: '医生', entry_date: '2019-06-15', status: 1 },
  { id: 3, name: '王阿姨', gender: 2, id_card: '110101197508019012', phone: '13900139003', email: 'wang@example.com', department_name: '后勤部', position: '保洁', entry_date: '2021-01-10', status: 1 }
])

const departmentList = ref([
  { id: 1, name: '护理部' },
  { id: 2, name: '医疗部' },
  { id: 3, name: '后勤部' },
  { id: 4, name: '行政部' },
  { id: 5, name: '财务部' }
])

const searchForm = reactive({
  name: '',
  phone: '',
  department_id: null,
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 3
})

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

const handleAdd = () => {
  ElMessage.info('新增员工功能开发中')
}

const handleView = (row) => {
  ElMessage.info('查看详情功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info('编辑功能开发中')
}

const handleLeave = (row) => {
  ElMessage.info('离职登记功能开发中')
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该员工吗？', '提示', {
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
  searchForm.department_id = null
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = () => {}
const handleCurrentChange = () => {}

onMounted(() => {})
</script>
