<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">房间类型</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增类型
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="类型名称">
          <el-input v-model="searchForm.name" placeholder="请输入类型名称" clearable />
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
        <el-table-column prop="name" label="类型名称" width="150" />
        <el-table-column prop="bed_count" label="床位数" width="100" />
        <el-table-column prop="price" label="价格(元/月)" width="150">
          <template #default="scope">
            <span style="color: #f56c6c; font-weight: bold;">¥{{ scope.row.price }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" width="250" show-overflow-tooltip />
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

    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      :close-on-click-modal="false"
    >
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="类型名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入类型名称" />
        </el-form-item>
        <el-form-item label="床位数" prop="bed_count">
          <el-input-number v-model="formData.bed_count" :min="1" :max="10" />
        </el-form-item>
        <el-form-item label="价格(元/月)" prop="price">
          <el-input-number v-model="formData.price" :min="0" :precision="2" />
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
  { id: 1, name: '单人间', bed_count: 1, price: 3000, description: '独立单人间，配备独立卫生间', status: 1, created_at: '2024-01-01 00:00:00' },
  { id: 2, name: '双人间', bed_count: 2, price: 2000, description: '双人间，适合结伴入住的长者', status: 1, created_at: '2024-01-01 00:00:00' },
  { id: 3, name: '豪华间', bed_count: 1, price: 5000, description: '豪华单人间，配备独立阳台和客厅', status: 1, created_at: '2024-01-01 00:00:00' },
  { id: 4, name: '特护间', bed_count: 1, price: 8000, description: '特护病房，配备专业医疗设备和24小时护理', status: 1, created_at: '2024-01-01 00:00:00' }
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
  bed_count: 1,
  price: 0,
  description: '',
  status: 1
})

const formRules = {
  name: [
    { required: true, message: '请输入类型名称', trigger: 'blur' }
  ]
}

const dialogTitle = computed(() => isEdit.value ? '编辑类型' : '新增类型')

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const resetForm = () => {
  formData.id = null
  formData.name = ''
  formData.bed_count = 1
  formData.price = 0
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
  formData.bed_count = row.bed_count
  formData.price = row.price
  formData.description = row.description
  formData.status = row.status
  dialogVisible.value = true
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
    await ElMessageBox.confirm('确定要删除该类型吗？', '提示', {
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
