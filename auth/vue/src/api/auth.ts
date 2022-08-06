import { auth } from "../config"
import * as components from "./authComponents"
export * from "./authComponents"

/**
 * @description "登录账号"
 * @param req
 */
export function signIn(req: components.SignInReq) {
	return auth.post<components.SignInReply>("/sign-in", req)
}

/**
 * @description "登录账号"
 */
export function verify() {
	return auth.get<components.VerifyTokenReply>("/verify-token")
}

/**
 * @description "退出登录"
 */
export function signOut() {
	return auth.post<components.SignOutReply>("/sign-out")
}
