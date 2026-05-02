import request from '@/utils/request'

export function getConsultations(params) {
  return request({
    url: '/consultations',
    method: 'get',
    params
  })
}

export function getConsultation(id) {
  return request({
    url: `/consultations/${id}`,
    method: 'get'
  })
}

export function createConsultation(data) {
  return request({
    url: '/consultations',
    method: 'post',
    data
  })
}

export function updateConsultation(id, data) {
  return request({
    url: `/consultations/${id}`,
    method: 'put',
    data
  })
}

export function deleteConsultation(id) {
  return request({
    url: `/consultations/${id}`,
    method: 'delete'
  })
}
