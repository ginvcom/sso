export const setCookie = (name: string, value: string, days?: number) => {
  if (!days) {
    days = 30 // 默认cookie将被保存30天
  }
  const exp  = new Date()    //new Date("December 31, 9998");
  exp.setTime(exp.getTime() + days * 24 * 60 * 60 * 1000)
  document.cookie = name + '=' + decodeURI(value) + ';expires=' + exp.toUTCString() + 'domain=localhost'
}

export const getCookie = (name: string): string | null => {
  let arr = document.cookie.match(new RegExp('(^| )' + name + '=([^;]*)(;|$)'))
  if (arr !== null) {
    return decodeURIComponent(arr[2])
  }

  return null
}

export const removeCookie = (name: string) => {
  document.cookie= name + '=0;expires=' + new Date(0).toUTCString()
}