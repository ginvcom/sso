import { oss } from "@/config"
import * as components from "./ossComponents"
export * from "./ossComponents"

/**
 * @description "操作对象列表"
 * @param params
 */
export function sts(params: components.StsReqParams) {
	return oss.get<components.StsReply>(`/aliyun/sts`, params)
}
