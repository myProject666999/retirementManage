<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">楼栋管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增楼栋
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="楼栋名称">
          <el-input v-model="searchForm.name" placeholder="请输入楼栋名称" clearable />
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
        <el-table-column prop="name" label="楼栋名称" width="150" />
        <el-table-column prop="floor_count" label="楼层数" width="100" />
        <el-table-column prop="description" label="描述" width="300" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="info" link @click="handleViewRooms(scope.row)">查看房间</el-button>
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

    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      :close-on-click-modal="false"
    >
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="楼栋名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入楼栋名称" />
        </el-form-item>
        <el-form-item label="楼层数" prop="floor_count">
          <el-input-number v-model="formData.floor_count" :min="1" :max="20" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入描述"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" :loading="submitLoading" @click="handleSubmit">
            确定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref(null)

const tableData = ref([
  { id: 1, name: '1号楼', floor_count: 6, description: '主要为自理能力较强的长者居住区域', status: 1, created_at: '2024-01-01 00:00:00' },
  { id: 2, name: '2号楼', floor_count: 6, description: '主要为半自理长者居住区域，配备护理站', status: 1, created_at: '2024-01-01 00:00:00' },
  { id: 3, name: '3号楼', floor_count: 4, description: '特护楼，主要为失能、重症长者居住区域', status: 1, created_at: '2024-01-01 00:00:00' },
  { id: 4, name: '综合楼', floor_count: 3, description: '包含餐厅、活动室、医疗室等公共设施', status: 1, created_at: '2024-01-01 00:00:00' }
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

const formData = reactive({
  id: null,
  name: '',
  floor_count: 1,
  description: '',
  status: 1
})

const formRules = {
  name: [
    { required: true, message: '请输入楼栋名称', trigger: 'blur' }
  ]
}

const dialogTitle = computed(() => isEdit.value ? '编辑楼栋' : '新增楼栋')

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const resetForm = () => {
  formData.id = null
  formData.name = ''
  formData.floor_count = 1
  formData.description = ''
  formData.status = 1
}

const handleAdd = () => {
  isEdit.value = false
  resetForm()
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  formData.id = row.id
  formData.name = row.name
  formData.floor_count = row.floor_count
  formData.description = row.description
  formData.status = row.status
  dialogVisible.value = true
}

const handleViewRooms = (row) => {
  ElMessage.info(`查看 ${row.name} 的房间列表功能开发中`)
}

const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      submitLoading.value = true
      try {
        if (isEdit.value) {
          ElMessage.success('更新成功')
        } else {
          ElMessage.success('创建成功')
        }
        dialogVisible.value = false
      } catch (error) {
        console.error('提交失败:', error)
      } finally {
        submitLoading.value = false
      }
    }
  })
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该楼栋吗？', '提示', {
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
