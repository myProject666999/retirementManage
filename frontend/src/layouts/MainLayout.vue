<template>
  <div class="main-container">
    <aside class="sidebar-container" :class="{ collapsed: isCollapsed }">
      <div class="sidebar-logo">
        <h1 v-if="!isCollapsed">养老院管理系统</h1>
        <h1 v-else>RM</h1>
      </div>
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapsed"
        :collapse-transition="false"
        router
      >
        <el-menu-item v-if="userStore.isSuperAdmin" index="/dashboard">
          <el-icon><Odometer /></el-icon>
          <template #title>首页统计</template>
        </el-menu-item>

        <el-sub-menu v-if="userStore.isSales" index="marketing">
          <template #title>
            <el-icon><Promotion /></el-icon>
            <span>营销管理</span>
          </template>
          <el-menu-item index="/marketing/consultation">咨询管理</el-menu-item>
          <el-menu-item index="/marketing/reservation">预定管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu v-if="userStore.isSuperAdmin" index="checkin">
          <template #title>
            <el-icon><OfficeBuilding /></el-icon>
            <span>入住管理</span>
          </template>
          <el-menu-item index="/checkin/bed-overview">床位全景</el-menu-item>
          <el-menu-item index="/checkin/contract">入住签约</el-menu-item>
          <el-menu-item index="/checkin/outing">外出登记</el-menu-item>
          <el-menu-item index="/checkin/visit">来访登记</el-menu-item>
          <el-menu-item index="/checkin/accident">事故登记</el-menu-item>
          <el-menu-item index="/checkin/checkout">退住申请</el-menu-item>
        </el-sub-menu>

        <el-sub-menu v-if="userStore.isHR" index="person">
          <template #title>
            <el-icon><User /></el-icon>
            <span>人员管理</span>
          </template>
          <el-menu-item index="/person/elder">长者档案</el-menu-item>
          <el-menu-item index="/person/employee">员工管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu v-if="userStore.isService" index="service">
          <template #title>
            <el-icon><Service /></el-icon>
            <span>服务管理</span>
          </template>
          <el-menu-item index="/service/item">服务项目</el-menu-item>
          <el-menu-item index="/service/care-level">护理等级</el-menu-item>
          <el-menu-item index="/service/reservation">服务预定</el-menu-item>
        </el-sub-menu>

        <el-sub-menu v-if="userStore.isCatering" index="catering">
          <template #title>
            <el-icon><KnifeFork /></el-icon>
            <span>餐饮管理</span>
          </template>
          <el-menu-item index="/catering/dish">菜品管理</el-menu-item>
          <el-menu-item index="/catering/package">餐饮套餐</el-menu-item>
          <el-menu-item index="/catering/order">点餐管理</el-menu-item>
        </el-sub-menu>

        <el-sub-menu v-if="userStore.isFinance" index="finance">
          <template #title>
            <el-icon><Wallet /></el-icon>
            <span>费用管理</span>
          </template>
          <el-menu-item index="/finance/recharge">预存充值</el-menu-item>
          <el-menu-item index="/finance/expense">消费记录</el-menu-item>
          <el-menu-item index="/finance/checkout-audit">退住审核</el-menu-item>
        </el-sub-menu>

        <el-sub-menu v-if="userStore.isSuperAdmin" index="config">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>基础配置</span>
          </template>
          <el-menu-item index="/config/channel">客户来源渠道</el-menu-item>
          <el-menu-item index="/config/room-type">房间类型</el-menu-item>
          <el-menu-item index="/config/building">楼栋管理</el-menu-item>
          <el-menu-item index="/config/user">用户管理</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </aside>

    <div class="app-container">
      <header class="header-container">
        <div class="header-left">
          <span class="toggle-btn" @click="toggleSidebar">
            <el-icon v-if="!isCollapsed"><Fold /></el-icon>
            <el-icon v-else><Expand /></el-icon>
          </span>
          <el-breadcrumb class="breadcrumb-container" separator="/">
            <el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item v-for="item in breadcrumbs" :key="item.path">
              {{ item.title }}
            </el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <div class="user-info">
              <div class="user-avatar">{{ userStore.userInfo.real_name?.charAt(0) || 'U' }}</div>
              <span class="user-name">{{ userStore.userInfo.real_name || userStore.userInfo.username }}</span>
              <el-icon><ArrowDown /></el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item divided command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </header>

      <main class="content-container">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessageBox, ElMessage } from 'element-plus'
import { useUserStore } from '@/store/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const isCollapsed = ref(false)

const activeMenu = computed(() => route.path)

const breadcrumbs = computed(() => {
  const matched = route.matched.filter(item => item.meta && item.meta.title)
  return matched.map(item => ({
    path: item.path,
    title: item.meta.title
  }))
})

const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value
}

const handleCommand = async (command) => {
  if (command === 'logout') {
    try {
      await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      await userStore.logout()
      ElMessage.success('退出登录成功')
      router.push('/login')
    } catch (error) {
      if (error !== 'cancel') {
        console.error('退出登录失败:', error)
      }
    }
  }
}
</script>
