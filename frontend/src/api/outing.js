import request from '@/utils/request'

export function getOutings(params) {
  return request({
    url: '/outings',
    method: 'get',
    params
  })
}

export function getOuting(id) {
  return request({
    url: `/outings/${id}`,
    method: 'get'
  })
}

export function createOuting(data) {
  return request({
    url: '/outings',
    method: 'post',
    data
  })
}

export function updateOuting(id, data) {
  return request({
    url: `/outings/${id}`,
    method: 'put',
    data
  })
}
