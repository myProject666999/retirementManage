import request from '@/utils/request'

export function getContracts(params) {
  return request({
    url: '/contracts',
    method: 'get',
    params
  })
}

export function getContract(id) {
  return request({
    url: `/contracts/${id}`,
    method: 'get'
  })
}

export function createContract(data) {
  return request({
    url: '/contracts',
    method: 'post',
    data
  })
}

export function updateContract(id, data) {
  return request({
    url: `/contracts/${id}`,
    method: 'put',
    data
  })
}
