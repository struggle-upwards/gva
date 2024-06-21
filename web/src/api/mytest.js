import service from '@/utils/request'

export const getMyTest = () => {
  return service({
    url: '/mytest/testapi',
    method: 'get'
  })
}