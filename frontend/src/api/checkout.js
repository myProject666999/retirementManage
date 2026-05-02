import request from '@/utils/request'

export function getCheckouts(params) {
  return request({
    url: '/checkouts',
    method: 'get',
    params
  })
}

export function getCheckout(id) {
  return request({
    url: `/checkouts/${id}`,
    method: 'get'
  })
}

export function approveCheckout(id) {
  return request({
    url: `/checkouts/${id}/approve`,
    method: 'post'
  })
}

export function rejectCheckout(id, auditRemark) {
  return request({
    url: `/checkouts/${id}/reject`,
    method: 'post',
    data: { audit_remark: auditRemark }
  })
}
