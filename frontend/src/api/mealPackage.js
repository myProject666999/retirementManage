import request from '@/utils/request'

export function getMealPackages(params) {
  return request({
    url: '/meal-packages',
    method: 'get',
    params
  })
}

export function getAllMealPackages() {
  return request({
    url: '/meal-packages/all',
    method: 'get'
  })
}

export function getMealPackage(id) {
  return request({
    url: `/meal-packages/${id}`,
    method: 'get'
  })
}

export function createMealPackage(data) {
  return request({
    url: '/meal-packages',
    method: 'post',
    data
  })
}

export function updateMealPackage(id, data) {
  return request({
    url: `/meal-packages/${id}`,
    method: 'put',
    data
  })
}

export function deleteMealPackage(id) {
  return request({
    url: `/meal-packages/${id}`,
    method: 'delete'
  })
}
