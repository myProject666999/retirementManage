import request from '@/utils/request'

export function getRecharges(params) {
  return request({
    url: '/recharges',
    method: 'get',
    params
  })
}

export function getRecharge(id) {
  return request({
    url: `/recharges/${id}`,
    method: 'get'
  })
}

export function createRecharge(data) {
  return request({
    url: '/recharges',
    method: 'post',
    data
  })
}

export function confirmRecharge(id) {
  return request({
    url: `/recharges/${id}/confirm`,
    method: 'post'
  })
}

export function cancelRecharge(id) {
  return request({
    url: `/recharges/${id}/cancel`,
    method: 'post'
  })
}
