import { ssoms } from "../config"

interface StsParams {
    bucket: string
}

interface StsReply {
    accessKeyId: string
    accessKeySecret: string
    stsToken: string
}

/**
 * @description "获取临时token"
 * @param params
 */
 export function sts(params: StsParams) {
	return ssoms.get<StsReply>("/aliyun/sts", params)
}