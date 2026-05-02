import request from '@/utils/request'

export function getCareLevels(params) {
  return request({
    url: '/care-levels',
    method: 'get',
    params
  })
}

export function getAllCareLevels() {
  return request({
    url: '/care-levels/all',
    method: 'get'
  })
}

export function getCareLevel(id) {
  return request({
    url: `/care-levels/${id}`,
    method: 'get'
  })
}

export function createCareLevel(data) {
  return request({
    url: '/care-levels',
    method: 'post',
    data
  })
}

export function updateCareLevel(id, data) {
  return request({
    url: `/care-levels/${id}`,
    method: 'put',
    data
  })
}

export function deleteCareLevel(id) {
  return request({
    url: `/care-levels/${id}`,
    method: 'delete'
  })
}
