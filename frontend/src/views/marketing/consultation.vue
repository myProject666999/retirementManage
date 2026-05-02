<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">咨询管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增咨询
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
            <el-option label="待跟进" :value="0" />
            <el-option label="已跟进" :value="1" />
            <el-option label="已转化" :value="2" />
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
        <el-table-column prop="phone" label="电话" width="150" />
        <el-table-column prop="gender" label="性别" width="80">
          <template #default="scope">
            {{ scope.row.gender === 1 ? '男' : '女' }}
          </template>
        </el-table-column>
        <el-table-column prop="age" label="年龄" width="80" />
        <el-table-column prop="channel_name" label="来源渠道" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="consult_time" label="咨询时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.consult_time) }}
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
      width="600px"
      :close-on-click-modal="false"
    >
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="姓名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="电话" prop="phone">
          <el-input v-model="formData.phone" placeholder="请输入电话" />
        </el-form-item>
        <el-form-item label="性别" prop="gender">
          <el-radio-group v-model="formData.gender">
            <el-radio :value="1">男</el-radio>
            <el-radio :value="2">女</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="年龄" prop="age">
          <el-input-number v-model="formData.age" :min="0" :max="150" />
        </el-form-item>
        <el-form-item label="来源渠道" prop="channel_id">
          <el-select v-model="formData.channel_id" placeholder="请选择来源渠道" style="width: 100%">
            <el-option
              v-for="channel in channelList"
              :key="channel.id"
              :label="channel.name"
              :value="channel.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="咨询内容" prop="content">
          <el-input
            v-model="formData.content"
            type="textarea"
            :rows="4"
            placeholder="请输入咨询内容"
          />
        </el-form-item>
        <el-form-item label="跟进情况" prop="follow_up">
          <el-input
            v-model="formData.follow_up"
            type="textarea"
            :rows="3"
            placeholder="请输入跟进情况"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="0">待跟进</el-radio>
            <el-radio :value="1">已跟进</el-radio>
            <el-radio :value="2">已转化</el-radio>
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
  { id: 1, name: '陈阿姨', phone: '13800138001', gender: 2, age: 75, channel_name: '网络推广', status: 0, consult_time: '2024-05-01 10:30:00' },
  { id: 2, name: '刘叔叔', phone: '13800138002', gender: 1, age: 78, channel_name: '口碑介绍', status: 1, consult_time: '2024-05-01 14:20:00' },
  { id: 3, name: '周阿姨', phone: '13800138003', gender: 2, age: 82, channel_name: '社区活动', status: 2, consult_time: '2024-05-02 09:15:00' }
])

const channelList = ref([
  { id: 1, name: '网络推广' },
  { id: 2, name: '口碑介绍' },
  { id: 3, name: '社区活动' },
  { id: 4, name: '电话咨询' },
  { id: 5, name: '其他渠道' }
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

const formData = reactive({
  id: null,
  name: '',
  phone: '',
  gender: 1,
  age: 0,
  channel_id: null,
  content: '',
  follow_up: '',
  status: 0
})

const formRules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入电话', trigger: 'blur' }
  ]
}

const dialogTitle = computed(() => isEdit.value ? '编辑咨询' : '新增咨询')

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

const getStatusText = (status) => {
  const texts = { 0: '待跟进', 1: '已跟进', 2: '已转化' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'primary', 2: 'success' }
  return types[status] || ''
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

const resetForm = () => {
  formData.id = null
  formData.name = ''
  formData.phone = ''
  formData.gender = 1
  formData.age = 0
  formData.channel_id = null
  formData.content = ''
  formData.follow_up = ''
  formData.status = 0
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
  formData.phone = row.phone
  formData.gender = row.gender
  formData.age = row.age
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

onMounted(() => {})
</script>
