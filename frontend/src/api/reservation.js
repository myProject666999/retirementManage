import request from '@/utils/request'

export function getReservations(params) {
  return request({
    url: '/reservations',
    method: 'get',
    params
  })
}

export function getReservation(id) {
  return request({
    url: `/reservations/${id}`,
    method: 'get'
  })
}

export function createReservation(data) {
  return request({
    url: '/reservations',
    method: 'post',
    data
  })
}

export function updateReservation(id, data) {
  return request({
    url: `/reservations/${id}`,
    method: 'put',
    data
  })
}

export function deleteReservation(id) {
  return request({
    url: `/reservations/${id}`,
    method: 'delete'
  })
}
