import request from '@/utils/request'

export function getVisits(params) {
  return request({
    url: '/visits',
    method: 'get',
    params
  })
}

export function getVisit(id) {
  return request({
    url: `/visits/${id}`,
    method: 'get'
  })
}

export function createVisit(data) {
  return request({
    url: '/visits',
    method: 'post',
    data
  })
}

export function updateVisit(id, data) {
  return request({
    url: `/visits/${id}`,
    method: 'put',
    data
  })
}
