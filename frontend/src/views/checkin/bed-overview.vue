<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">床位全景</span>
    </div>

    <el-card class="filter-card">
      <el-form :inline="true" :model="filterForm">
        <el-form-item label="楼栋">
          <el-select v-model="filterForm.building_id" placeholder="请选择楼栋" clearable>
            <el-option
              v-for="building in buildingList"
              :key="building.id"
              :label="building.name"
              :value="building.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="楼层">
          <el-select v-model="filterForm.floor" placeholder="请选择楼层" clearable>
            <el-option
              v-for="floor in floorList"
              :key="floor"
              :label="`${floor}楼`"
              :value="floor"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="filterForm.status" placeholder="请选择状态" clearable>
            <el-option label="空闲" :value="0" />
            <el-option label="已入住" :value="1" />
            <el-option label="维修中" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>
            查询
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="stats-card">
      <el-row :gutter="20">
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-value" style="color: #409eff">{{ stats.total }}</div>
            <div class="stat-label">总床位数</div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-value" style="color: #67c23a">{{ stats.occupied }}</div>
            <div class="stat-label">已入住</div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-value" style="color: #e6a23c">{{ stats.available }}</div>
            <div class="stat-label">可售床位</div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-value" style="color: #909399">{{ stats.maintenance }}</div>
            <div class="stat-label">维修中</div>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <el-card>
      <template #header>
        <span>床位分布</span>
      </template>
      
      <div class="bed-grid">
        <div
          v-for="bed in bedList"
          :key="bed.id"
          class="bed-item"
          :class="getBedClass(bed.status)"
          @click="handleBedClick(bed)"
        >
          <div class="bed-number">{{ bed.room_number }}-{{ bed.bed_number }}</div>
          <div class="bed-status">{{ getBedStatusText(bed.status) }}</div>
          <div class="bed-elder" v-if="bed.status === 1 && bed.elder_name">
            {{ bed.elder_name }}
          </div>
        </div>
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="房间号">{{ currentBed.room_number }}</el-descriptions-item>
        <el-descriptions-item label="床位号">{{ currentBed.bed_number }}</el-descriptions-item>
        <el-descriptions-item label="楼栋">{{ currentBed.building_name }}</el-descriptions-item>
        <el-descriptions-item label="楼层">{{ currentBed.floor }}楼</el-descriptions-item>
        <el-descriptions-item label="房间类型">{{ currentBed.room_type_name }}</el-descriptions-item>
        <el-descriptions-item label="床位状态">
          <el-tag :type="getBedStatusTagType(currentBed.status)">
            {{ getBedStatusText(currentBed.status) }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item v-if="currentBed.status === 1" label="入住老人" :span="2">
          {{ currentBed.elder_name }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'

const dialogVisible = ref(false)
const dialogTitle = ref('床位详情')

const filterForm = reactive({
  building_id: null,
  floor: null,
  status: null
})

const stats = reactive({
  total: 200,
  occupied: 155,
  available: 40,
  maintenance: 5
})

const buildingList = ref([
  { id: 1, name: '1号楼' },
  { id: 2, name: '2号楼' },
  { id: 3, name: '3号楼' }
])

const floorList = ref([1, 2, 3, 4, 5, 6])

const bedList = ref([])
const currentBed = ref({})

const generateMockData = () => {
  const beds = []
  const roomNumbers = ['101', '102', '103', '201', '202', '203', '301', '302', '303']
  
  for (let i = 0; i < 27; i++) {
    const roomIdx = Math.floor(i / 3)
    const bedIdx = (i % 3) + 1
    beds.push({
      id: i + 1,
      room_number: roomNumbers[roomIdx],
      bed_number: bedIdx,
      building_name: '1号楼',
      floor: parseInt(roomNumbers[roomIdx].charAt(0)),
      room_type_name: roomIdx % 3 === 0 ? '单人间' : '双人间',
      status: Math.floor(Math.random() * 3),
      elder_name: Math.random() > 0.5 ? `张${['阿姨', '叔叔', '奶奶', '爷爷'][Math.floor(Math.random() * 4)]}` : null
    })
  }
  bedList.value = beds
}

const getBedClass = (status) => {
  const classes = {
    0: 'bed-available',
    1: 'bed-occupied',
    2: 'bed-maintenance'
  }
  return classes[status] || ''
}

const getBedStatusText = (status) => {
  const texts = { 0: '空闲', 1: '已入住', 2: '维修中' }
  return texts[status] || '未知'
}

const getBedStatusTagType = (status) => {
  const types = { 0: 'warning', 1: 'success', 2: 'info' }
  return types[status] || ''
}

const handleBedClick = (bed) => {
  currentBed.value = bed
  dialogTitle.value = `${bed.room_number}-${bed.bed_number} 详情`
  dialogVisible.value = true
}

const handleSearch = () => {
  generateMockData()
}

const handleReset = () => {
  filterForm.building_id = null
  filterForm.floor = null
  filterForm.status = null
  generateMockData()
}

onMounted(() => {
  generateMockData()
})
</script>

<style scoped>
.filter-card {
  margin-bottom: 20px;
}

.stats-card {
  margin-bottom: 20px;
}

.stat-item {
  text-align: center;
  padding: 20px;
}

.stat-value {
  font-size: 32px;
  font-weight: bold;
  margin-bottom: 10px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.bed-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 15px;
}

.bed-item {
  padding: 15px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  text-align: center;
  border: 2px solid transparent;
}

.bed-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.bed-available {
  background-color: #fdf6ec;
  border-color: #e6a23c;
}

.bed-occupied {
  background-color: #f0f9eb;
  border-color: #67c23a;
}

.bed-maintenance {
  background-color: #f4f4f5;
  border-color: #909399;
}

.bed-number {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}

.bed-status {
  font-size: 12px;
  margin-bottom: 5px;
}

.bed-elder {
  font-size: 12px;
  color: #606266;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
