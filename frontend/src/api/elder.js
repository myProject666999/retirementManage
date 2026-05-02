import request from '@/utils/request'

export function getElders(params) {
  return request({
    url: '/elders',
    method: 'get',
    params
  })
}

export function getAllElders() {
  return request({
    url: '/elders/all',
    method: 'get'
  })
}

export function getElder(id) {
  return request({
    url: `/elders/${id}`,
    method: 'get'
  })
}

export function createElder(data) {
  return request({
    url: '/elders',
    method: 'post',
    data
  })
}

export function updateElder(id, data) {
  return request({
    url: `/elders/${id}`,
    method: 'put',
    data
  })
}

export function deleteElder(id) {
  return request({
    url: `/elders/${id}`,
    method: 'delete'
  })
}
