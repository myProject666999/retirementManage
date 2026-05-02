import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/store/user'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { title: '登录', requiresAuth: false }
  },
  {
    path: '/',
    component: () => import('@/layouts/MainLayout.vue'),
    redirect: '/dashboard',
    meta: { requiresAuth: true },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: { title: '首页统计', roles: ['super_admin'] }
      },
      {
        path: 'marketing',
        name: 'Marketing',
        meta: { title: '营销管理', roles: ['sales', 'super_admin'] },
        children: [
          {
            path: 'consultation',
            name: 'Consultation',
            component: () => import('@/views/marketing/consultation.vue'),
            meta: { title: '咨询管理' }
          },
          {
            path: 'reservation',
            name: 'Reservation',
            component: () => import('@/views/marketing/reservation.vue'),
            meta: { title: '预定管理' }
          }
        ]
      },
      {
        path: 'checkin',
        name: 'Checkin',
        meta: { title: '入住管理', roles: ['super_admin'] },
        children: [
          {
            path: 'bed-overview',
            name: 'BedOverview',
            component: () => import('@/views/checkin/bed-overview.vue'),
            meta: { title: '床位全景' }
          },
          {
            path: 'contract',
            name: 'Contract',
            component: () => import('@/views/checkin/contract.vue'),
            meta: { title: '入住签约' }
          },
          {
            path: 'outing',
            name: 'Outing',
            component: () => import('@/views/checkin/outing.vue'),
            meta: { title: '外出登记' }
          },
          {
            path: 'visit',
            name: 'Visit',
            component: () => import('@/views/checkin/visit.vue'),
            meta: { title: '来访登记' }
          },
          {
            path: 'accident',
            name: 'Accident',
            component: () => import('@/views/checkin/accident.vue'),
            meta: { title: '事故登记' }
          },
          {
            path: 'checkout',
            name: 'Checkout',
            component: () => import('@/views/checkin/checkout.vue'),
            meta: { title: '退住申请' }
          }
        ]
      },
      {
        path: 'person',
        name: 'Person',
        meta: { title: '人员管理', roles: ['hr', 'super_admin'] },
        children: [
          {
            path: 'elder',
            name: 'Elder',
            component: () => import('@/views/person/elder.vue'),
            meta: { title: '长者档案' }
          },
          {
            path: 'employee',
            name: 'Employee',
            component: () => import('@/views/person/employee.vue'),
            meta: { title: '员工管理' }
          }
        ]
      },
      {
        path: 'service',
        name: 'Service',
        meta: { title: '服务管理', roles: ['service', 'super_admin'] },
        children: [
          {
            path: 'item',
            name: 'ServiceItem',
            component: () => import('@/views/service/item.vue'),
            meta: { title: '服务项目' }
          },
          {
            path: 'care-level',
            name: 'CareLevel',
            component: () => import('@/views/service/care-level.vue'),
            meta: { title: '护理等级' }
          },
          {
            path: 'reservation',
            name: 'ServiceReservation',
            component: () => import('@/views/service/reservation.vue'),
            meta: { title: '服务预定' }
          }
        ]
      },
      {
        path: 'catering',
        name: 'Catering',
        meta: { title: '餐饮管理', roles: ['catering', 'super_admin'] },
        children: [
          {
            path: 'dish',
            name: 'Dish',
            component: () => import('@/views/catering/dish.vue'),
            meta: { title: '菜品管理' }
          },
          {
            path: 'package',
            name: 'MealPackage',
            component: () => import('@/views/catering/package.vue'),
            meta: { title: '餐饮套餐' }
          },
          {
            path: 'order',
            name: 'Order',
            component: () => import('@/views/catering/order.vue'),
            meta: { title: '点餐管理' }
          }
        ]
      },
      {
        path: 'finance',
        name: 'Finance',
        meta: { title: '费用管理', roles: ['finance', 'super_admin'] },
        children: [
          {
            path: 'recharge',
            name: 'Recharge',
            component: () => import('@/views/finance/recharge.vue'),
            meta: { title: '预存充值' }
          },
          {
            path: 'expense',
            name: 'Expense',
            component: () => import('@/views/finance/expense.vue'),
            meta: { title: '消费记录' }
          },
          {
            path: 'checkout-audit',
            name: 'CheckoutAudit',
            component: () => import('@/views/finance/checkout-audit.vue'),
            meta: { title: '退住审核' }
          }
        ]
      },
      {
        path: 'config',
        name: 'Config',
        meta: { title: '基础配置', roles: ['super_admin'] },
        children: [
          {
            path: 'channel',
            name: 'Channel',
            component: () => import('@/views/config/channel.vue'),
            meta: { title: '客户来源渠道' }
          },
          {
            path: 'room-type',
            name: 'RoomType',
            component: () => import('@/views/config/room-type.vue'),
            meta: { title: '房间类型' }
          },
          {
            path: 'building',
            name: 'Building',
            component: () => import('@/views/config/building.vue'),
            meta: { title: '楼栋管理' }
          },
          {
            path: 'user',
            name: 'UserManagement',
            component: () => import('@/views/config/user.vue'),
            meta: { title: '用户管理' }
          }
        ]
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} - 养老院管理系统` : '养老院管理系统'

  const userStore = useUserStore()
  const token = localStorage.getItem('token')

  if (to.meta.requiresAuth !== false) {
    if (!token) {
      next({ path: '/login', query: { redirect: to.fullPath } })
      return
    }

    if (!userStore.userInfo.id) {
      try {
        await userStore.getUserInfo()
      } catch (error) {
        localStorage.removeItem('token')
        next({ path: '/login' })
        return
      }
    }

    if (to.meta.roles && to.meta.roles.length > 0) {
      const hasRole = to.meta.roles.includes(userStore.userInfo.role_code)
      if (!hasRole) {
        next({ path: '/403' })
        return
      }
    }

    next()
  } else {
    if (token && to.path === '/login') {
      next({ path: '/dashboard' })
      return
    }
    next()
  }
})

export default router
