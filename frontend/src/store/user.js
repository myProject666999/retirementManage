import { defineStore } from 'pinia'
import { login, getUserInfo, logout } from '@/api/auth'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userInfo: {
      id: null,
      username: '',
      real_name: '',
      phone: '',
      email: '',
      role_id: null,
      role_name: '',
      role_code: ''
    }
  }),

  getters: {
    isLoggedIn: (state) => !!state.token,
    isSuperAdmin: (state) => state.userInfo.role_code === 'super_admin',
    isSales: (state) => state.userInfo.role_code === 'sales' || state.userInfo.role_code === 'super_admin',
    isHR: (state) => state.userInfo.role_code === 'hr' || state.userInfo.role_code === 'super_admin',
    isService: (state) => state.userInfo.role_code === 'service' || state.userInfo.role_code === 'super_admin',
    isCatering: (state) => state.userInfo.role_code === 'catering' || state.userInfo.role_code === 'super_admin',
    isFinance: (state) => state.userInfo.role_code === 'finance' || state.userInfo.role_code === 'super_admin'
  },

  actions: {
    async login(loginForm) {
      try {
        const res = await login(loginForm)
        this.token = res.data.token
        localStorage.setItem('token', res.data.token)
        this.userInfo = res.data.user
        return res
      } catch (error) {
        throw error
      }
    },

    async getUserInfo() {
      try {
        const res = await getUserInfo()
        this.userInfo = res.data
        return res
      } catch (error) {
        throw error
      }
    },

    async logout() {
      try {
        await logout()
      } finally {
        this.token = ''
        this.userInfo = {
          id: null,
          username: '',
          real_name: '',
          phone: '',
          email: '',
          role_id: null,
          role_name: '',
          role_code: ''
        }
        localStorage.removeItem('token')
      }
    }
  }
})
