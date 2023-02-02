import Ajax from "@/utils/ajax"

export const systemCode = <string>import.meta.env.VITE_SYSTEM_NAME || ''

export const SERVER_ROUTER_MENU_KEY = systemCode + ':' + 'server_router_menu'

export const ossConfig = {
  ginv: {
    bucket: 'ginv',
    domain: '//assets.ginv.com/'
  },
  doc: {
    bucket: 'doc',
    domain: '//doc.ginv.com/'
  }
}

export const ssoms = new Ajax(systemCode, 'ssoms')
export const auth = new Ajax(systemCode, 'auth')
export const oss = new Ajax(systemCode, 'oss')
