import { auth } from "@/config"
import * as components from "./authComponents"
export * from "./authComponents"

/**
 * @description "登录账号"
 * @param req
 */
export function signIn(req: components.SignInReq) {
	return auth.post<components.SignInReply>(`/sign-in`, req)
}

/**
 * @description "验证请求"
 * @param req
 */
export function verify(req: components.VerifyRequestReq) {
	return auth.post<components.VerifyRequestReply>(`/verify-request`, req)
}

/**
 * @description "当前作用的角色分配的菜单权限"
 */
export function menus() {
	return auth.get<components.SessionMenusReply>(`/session-menus`)
}

/**
 * @description "退出登录"
 */
export function signOut() {
	return auth.post<components.SignOutReply>(`/sign-out`)
}
