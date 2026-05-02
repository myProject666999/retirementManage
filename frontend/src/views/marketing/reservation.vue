<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">预定管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增预定
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
            <el-option label="待确认" :value="0" />
            <el-option label="已确认" :value="1" />
            <el-option label="已签约" :value="2" />
            <el-option label="已取消" :value="3" />
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
        <el-table-column prop="room_type_name" label="房间类型" width="120" />
        <el-table-column prop="deposit_amount" label="预定金" width="120">
          <template #default="scope">
            ¥{{ scope.row.deposit_amount }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="reservation_date" label="预定日期" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.reservation_date) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="success" link @click="handleConfirm(scope.row)">确认</el-button>
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
        <el-form-item label="身份证号" prop="id_card">
          <el-input v-model="formData.id_card" placeholder="请输入身份证号" />
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
        <el-form-item label="房间类型" prop="room_type_id">
          <el-select v-model="formData.room_type_id" placeholder="请选择房间类型" style="width: 100%">
            <el-option
              v-for="type in roomTypeList"
              :key="type.id"
              :label="type.name"
              :value="type.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="预定日期" prop="reservation_date">
          <el-date-picker
            v-model="formData.reservation_date"
            type="date"
            placeholder="选择预定日期"
            value-format="YYYY-MM-DD"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="预定金" prop="deposit_amount">
          <el-input-number v-model="formData.deposit_amount" :min="0" :precision="2" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :value="0">待确认</el-radio>
            <el-radio :value="1">已确认</el-radio>
            <el-radio :value="2">已签约</el-radio>
            <el-radio :value="3">已取消</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remarks">
          <el-input
            v-model="formData.remarks"
            type="textarea"
            :rows="3"
            placeholder="请输入备注"
          />
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
  { id: 1, name: '王阿姨', phone: '13900139001', gender: 2, age: 72, room_type_name: '单人间', deposit_amount: 5000, status: 1, reservation_date: '2024-05-10' },
  { id: 2, name: '李叔叔', phone: '13900139002', gender: 1, age: 75, room_type_name: '双人间', deposit_amount: 3000, status: 0, reservation_date: '2024-05-15' },
  { id: 3, name: '张阿姨', phone: '13900139003', gender: 2, age: 80, room_type_name: '豪华间', deposit_amount: 8000, status: 2, reservation_date: '2024-04-20' }
])

const roomTypeList = ref([
  { id: 1, name: '单人间' },
  { id: 2, name: '双人间' },
  { id: 3, name: '豪华间' },
  { id: 4, name: '特护间' }
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
  id_card: '',
  gender: 1,
  age: 0,
  room_type_id: null,
  reservation_date: '',
  deposit_amount: 0,
  status: 0,
  remarks: ''
})

const formRules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入电话', trigger: 'blur' }
  ]
}

const dialogTitle = computed(() => isEdit.value ? '编辑预定' : '新增预定')

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

const getStatusText = (status) => {
  const texts = { 0: '待确认', 1: '已确认', 2: '已签约', 3: '已取消' }
  return texts[status] || '未知'
}

const getStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'primary', 2: 'success', 3: 'info' }
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
  formData.id_card = ''
  formData.gender = 1
  formData.age = 0
  formData.room_type_id = null
  formData.reservation_date = ''
  formData.deposit_amount = 0
  formData.status = 0
  formData.remarks = ''
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
  formData.deposit_amount = row.deposit_amount
  formData.status = row.status
  dialogVisible.value = true
}

const handleConfirm = (row) => {
  ElMessage.success('确认成功')
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
