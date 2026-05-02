import request from '@/utils/request'

export function getServiceItems(params) {
  return request({
    url: '/service-items',
    method: 'get',
    params
  })
}

export function getAllServiceItems() {
  return request({
    url: '/service-items/all',
    method: 'get'
  })
}

export function getServiceItem(id) {
  return request({
    url: `/service-items/${id}`,
    method: 'get'
  })
}

export function createServiceItem(data) {
  return request({
    url: '/service-items',
    method: 'post',
    data
  })
}

export function updateServiceItem(id, data) {
  return request({
    url: `/service-items/${id}`,
    method: 'put',
    data
  })
}

export function deleteServiceItem(id) {
  return request({
    url: `/service-items/${id}`,
    method: 'delete'
  })
}
