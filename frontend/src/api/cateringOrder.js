import request from '@/utils/request'

export function getCateringOrders(params) {
  return request({
    url: '/catering-orders',
    method: 'get',
    params
  })
}

export function getCateringOrder(id) {
  return request({
    url: `/catering-orders/${id}`,
    method: 'get'
  })
}

export function updateCateringOrderStatus(id, status) {
  return request({
    url: `/catering-orders/${id}/status`,
    method: 'put',
    data: { status }
  })
}
