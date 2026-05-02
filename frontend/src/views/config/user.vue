<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">用户管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增用户
      </el-button>
    </div>

    <div class="search-form">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="用户名">
          <el-input v-model="searchForm.username" placeholder="请输入用户名" clearable />
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
        <el-table-column prop="username" label="用户名" width="150" />
        <el-table-column prop="real_name" label="姓名" width="120" />
        <el-table-column prop="phone" label="手机号" width="150" />
        <el-table-column prop="email" label="邮箱" width="200" />
        <el-table-column label="角色" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.role" :type="getRoleTagType(scope.row.role.code)">
              {{ scope.row.role.name }}
            </el-tag>
          </template>
        </el-table-column>
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
            <el-button type="primary" link @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="warning" link @click="handleResetPassword(scope.row)">重置密码</el-button>
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
      width="600px"
      :close-on-click-modal="false"
    >
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="formData.username" placeholder="请输入用户名" :disabled="isEdit" />
        </el-form-item>
        <el-form-item v-if="!isEdit" label="密码" prop="password">
          <el-input v-model="formData.password" type="password" placeholder="请输入密码" show-password />
        </el-form-item>
        <el-form-item label="姓名" prop="real_name">
          <el-input v-model="formData.real_name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="formData.phone" placeholder="请输入手机号" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="formData.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="角色" prop="role_id">
          <el-select v-model="formData.role_id" placeholder="请选择角色" style="width: 100%">
            <el-option
              v-for="role in roleList"
              :key="role.id"
              :label="role.name"
              :value="role.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-if="isEdit" label="状态" prop="status">
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
import { getUsers, createUser, updateUser, deleteUser, resetPassword } from '@/api/user'

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref(null)

const tableData = ref([])
const roleList = ref([
  { id: 1, name: '销售员', code: 'sales' },
  { id: 2, name: '人事', code: 'hr' },
  { id: 3, name: '服务员', code: 'service' },
  { id: 4, name: '餐饮员', code: 'catering' },
  { id: 5, name: '财务人员', code: 'finance' },
  { id: 6, name: '超级管理员', code: 'super_admin' }
])

const searchForm = reactive({
  username: '',
  status: null
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

const formData = reactive({
  id: null,
  username: '',
  password: '',
  real_name: '',
  phone: '',
  email: '',
  role_id: null,
  status: 1
})

const formRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
  role_id: [
    { required: true, message: '请选择角色', trigger: 'change' }
  ]
}

const dialogTitle = computed(() => isEdit.value ? '编辑用户' : '新增用户')

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const getRoleTagType = (code) => {
  const types = {
    'sales': 'primary',
    'hr': 'success',
    'service': 'warning',
    'catering': 'info',
    'finance': 'danger',
    'super_admin': 'danger'
  }
  return types[code] || ''
}

const fetchData = async () => {
  loading.value = true
  try {
    const res = await getUsers({
      page: pagination.page,
      page_size: pagination.page_size,
      ...searchForm
    })
    tableData.value = res.data.list
    pagination.total = res.data.total
  } catch (error) {
    console.error('获取数据失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  fetchData()
}

const handleReset = () => {
  searchForm.username = ''
  searchForm.status = null
  handleSearch()
}

const handleSizeChange = (size) => {
  pagination.page_size = size
  fetchData()
}

const handleCurrentChange = (page) => {
  pagination.page = page
  fetchData()
}

const resetForm = () => {
  formData.id = null
  formData.username = ''
  formData.password = ''
  formData.real_name = ''
  formData.phone = ''
  formData.email = ''
  formData.role_id = null
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
  formData.username = row.username
  formData.real_name = row.real_name
  formData.phone = row.phone
  formData.email = row.email
  formData.role_id = row.role_id
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
          await updateUser(formData.id, formData)
          ElMessage.success('更新成功')
        } else {
          await createUser(formData)
          ElMessage.success('创建成功')
        }
        dialogVisible.value = false
        fetchData()
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
    await ElMessageBox.confirm('确定要删除该用户吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteUser(row.id)
    ElMessage.success('删除成功')
    fetchData()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
    }
  }
}

const handleResetPassword = async (row) => {
  try {
    await ElMessageBox.confirm('确定要重置该用户的密码吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await resetPassword(row.id)
    ElMessage.success('密码重置成功，默认密码: 123456')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('重置密码失败:', error)
    }
  }
}

onMounted(() => {
  fetchData()
})
</script>
