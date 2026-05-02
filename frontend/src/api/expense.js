import request from '@/utils/request'

export function getExpenses(params) {
  return request({
    url: '/expenses',
    method: 'get',
    params
  })
}

export function getExpense(id) {
  return request({
    url: `/expenses/${id}`,
    method: 'get'
  })
}
