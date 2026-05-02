import request from '@/utils/request'

export function getAccidents(params) {
  return request({
    url: '/accidents',
    method: 'get',
    params
  })
}

export function getAccident(id) {
  return request({
    url: `/accidents/${id}`,
    method: 'get'
  })
}

export function createAccident(data) {
  return request({
    url: '/accidents',
    method: 'post',
    data
  })
}

export function updateAccident(id, data) {
  return request({
    url: `/accidents/${id}`,
    method: 'put',
    data
  })
}
