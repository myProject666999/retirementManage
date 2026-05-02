import request from '@/utils/request'

export function getServiceReservations(params) {
  return request({
    url: '/service-reservations',
    method: 'get',
    params
  })
}

export function getServiceReservation(id) {
  return request({
    url: `/service-reservations/${id}`,
    method: 'get'
  })
}

export function createServiceReservation(data) {
  return request({
    url: '/service-reservations',
    method: 'post',
    data
  })
}

export function updateServiceReservation(id, data) {
  return request({
    url: `/service-reservations/${id}`,
    method: 'put',
    data
  })
}

export function deleteServiceReservation(id) {
  return request({
    url: `/service-reservations/${id}`,
    method: 'delete'
  })
}
