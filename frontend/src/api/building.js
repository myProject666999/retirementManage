import request from '@/utils/request'

export function getBuildings(params) {
  return request({
    url: '/buildings',
    method: 'get',
    params
  })
}

export function getAllBuildings() {
  return request({
    url: '/buildings/all',
    method: 'get'
  })
}

export function getBuilding(id) {
  return request({
    url: `/buildings/${id}`,
    method: 'get'
  })
}

export function createBuilding(data) {
  return request({
    url: '/buildings',
    method: 'post',
    data
  })
}

export function updateBuilding(id, data) {
  return request({
    url: `/buildings/${id}`,
    method: 'put',
    data
  })
}

export function deleteBuilding(id) {
  return request({
    url: `/buildings/${id}`,
    method: 'delete'
  })
}
