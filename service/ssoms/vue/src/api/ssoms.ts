import { ssoms } from "@/config"
import * as components from "./ssomsComponents"
export * from "./ssomsComponents"

/**
 * @description "首页统计数据"
 */
export function homeStatistic() {
	return ssoms.get<components.StatisticReply>(`/home/statistic`)
}

/**
 * @description "用户列表"
 * @param params
 */
export function userList(params: components.UserListReqParams) {
	return ssoms.get<components.UserListReply>(`/user`, params)
}

/**
 * @description "新增用户"
 * @param req
 */
export function addUser(req: components.UserForm) {
	return ssoms.post<components.AddUserReply>(`/user`, req)
}

/**
 * @description "用户详情"
 * @param params
 */
export function userDetail(params: components.UserDetailReqParams, uuid: string) {
	return ssoms.get<components.UserForm>(`/user/${uuid}`, params)
}

/**
 * @description "修改用户"
 * @param params
 * @param req
 */
export function updateUser(params: components.UpdateUserReqParams, req: components.UpdateUserReq, uuid: string) {
	return ssoms.put<components.UpdateUserReply>(`/user/${uuid}`, params, req)
}

/**
 * @description "删除用户"
 * @param params
 */
export function deleteUser(params: components.DeleteUserReqParams, uuid: string) {
	return ssoms.delete<components.DeleteUserReply>(`/user/${uuid}`, params)
}

/**
 * @description "用户搜索下拉选项"
 * @param params
 */
export function userFilterOptions(params: components.UserFilterOptionsReqParams) {
	return ssoms.get<components.UserFilterOptionsReply>(`/filterOptions`, params)
}

/**
 * @description "用户已分配的角色"
 * @param params
 */
export function assignedRoles(params: components.AssignedRolesReqParams) {
	return ssoms.get<components.AssignedRolesReply>(`/assignedRoles`, params)
}

/**
 * @description "给用户分配角色"
 * @param req
 */
export function assignRole(req: components.AssignRoleReq) {
	return ssoms.post<components.AssignRoleReply>(`/assignRole`, req)
}

/**
 * @description "用户通过其分配的角色获得的权限"
 * @param req
 */
export function userPermissions(req: components.UserPermissionsReq) {
	return ssoms.get<components.UserPermissionsReply>(`/permissions`, req)
}

/**
 * @description "个人中心用户信息"
 */
export function profile() {
	return ssoms.get<components.UserForm>(`/user/profile`)
}

/**
 * @description "个人中心头像上传修改"
 * @param req
 */
export function avatarUpload(req: components.AvatarUploadReq) {
	return ssoms.post<components.AvatarUploadReply>(`/user/avatarUpload`, req)
}

/**
 * @description "个人中心基础信息修改"
 * @param req
 */
export function infoEdit(req: components.InfoEditReq) {
	return ssoms.post<components.InfoEditReply>(`/user/infoEdit`, req)
}

/**
 * @description "个人中心密码重置"
 * @param req
 */
export function passwordReset(req: components.PasswordResetReq) {
	return ssoms.post<components.PasswordResetReply>(`/user/passwordReset`, req)
}

/**
 * @description "角色列表"
 * @param params
 */
export function roleList(params: components.RoleListReqParams) {
	return ssoms.get<components.RoleListReply>(`/role`, params)
}

/**
 * @description "新增角色"
 * @param req
 */
export function addRole(req: components.RoleForm) {
	return ssoms.post<components.AddRoleReply>(`/role`, req)
}

/**
 * @description "角色详情"
 * @param params
 */
export function roleDetail(params: components.RoleDetailReqParams, roleUUID: string) {
	return ssoms.get<components.RoleForm>(`/role/${roleUUID}`, params)
}

/**
 * @description "修改角色"
 * @param params
 * @param req
 */
export function updateRole(params: components.UpdateRoleReqParams, req: components.UpdateRoleReq, roleUUID: string) {
	return ssoms.put<components.UpdateRoleReply>(`/role/${roleUUID}`, params, req)
}

/**
 * @description "删除角色"
 * @param params
 */
export function deleteRole(params: components.DeleteRoleReqParams, roleUUID: string) {
	return ssoms.delete<components.DeleteRoleReply>(`/role/${roleUUID}`, params)
}

/**
 * @description "角色已拥有的用户"
 */
export function roleOptions() {
	return ssoms.get<components.OptionsReply>(`/role/options`)
}

/**
 * @description "角色已拥有的用户"
 * @param params
 */
export function assignedUsers(params: components.AssignedUsersReqParams) {
	return ssoms.get<components.AssignedUsersReply>(`/assignedUsers`, params)
}

/**
 * @description "给角色添加用户"
 * @param params
 * @param req
 */
export function assignUser(params: components.AssignUserReqParams, req: components.AssignUserReq, roleUUID: string) {
	return ssoms.patch<components.AssignUserReply>(`/role/${roleUUID}/assignUser`, params, req)
}

/**
 * @description "给角色移出用户"
 * @param params
 * @param req
 */
export function deassignUser(params: components.DeassignUserReqParams, req: components.DeassignUserReq, roleUUID: string) {
	return ssoms.patch<components.DeassignUserReply>(`/role/${roleUUID}/deassignUser`, params, req)
}

/**
 * @description "角色已继承的角色"
 * @param params
 */
export function inheritances(params: components.InheritancesReqParams, roleUUID: string) {
	return ssoms.get<components.InheritancesReply>(`/role/${roleUUID}/inheritances`, params)
}

/**
 * @description "角色继承另一个角色"
 * @param params
 * @param req
 */
export function extendRole(params: components.AddInheritanceReqParams, req: components.AddInheritanceReq, roleUUID: string) {
	return ssoms.patch<components.AddInheritanceReply>(`/role/${roleUUID}/addInheritance`, params, req)
}

/**
 * @description "系统列表"
 * @param params
 */
export function systemList(params: components.SystemListReqParams) {
	return ssoms.get<components.SystemListReply>(`/system`, params)
}

/**
 * @description "新增系统"
 * @param req
 */
export function addSystem(req: components.SystemForm) {
	return ssoms.post<components.AddSystemReply>(`/system`, req)
}

/**
 * @description "系统详情"
 * @param params
 */
export function systemDetail(params: components.SystemDetailReqParams, uuid: string) {
	return ssoms.get<components.SystemForm>(`/system/${uuid}`, params)
}

/**
 * @description "修改系统"
 * @param params
 * @param req
 */
export function updateSystem(params: components.UpdateSystemReqParams, req: components.UpdateSystemReq, uuid: string) {
	return ssoms.put<components.UpdateSystemReply>(`/system/${uuid}`, params, req)
}

/**
 * @description "删除系统"
 * @param params
 */
export function deleteSystem(params: components.DeleteSystemReqParams, uuid: string) {
	return ssoms.delete<components.DeleteSystemReply>(`/system/${uuid}`, params)
}

/**
 * @description "操作对象列表"
 * @param params
 */
export function objectList(params: components.ObjectListReqParams) {
	return ssoms.get<components.ObjectListReply>(`/object`, params)
}

/**
 * @description "新增操作对象"
 * @param req
 */
export function addObject(req: components.ObjectForm) {
	return ssoms.post<components.AddObjectReply>(`/object`, req)
}

/**
 * @description "操作对象详情"
 * @param params
 */
export function objectDetail(params: components.ObjectDetailReqParams, uuid: string) {
	return ssoms.get<components.ObjectForm>(`/object/${uuid}`, params)
}

/**
 * @description "操作对象导入"
 * @param req
 */
export function importObject(req: components.ImportObjectReq) {
	return ssoms.post<components.ImportObjectReply>(`/object/import`, req)
}

/**
 * @description "修改操作对象"
 * @param params
 * @param req
 */
export function updateObject(params: components.UpdateObjectReqParams, req: components.UpdateObjectReq, uuid: string) {
	return ssoms.put<components.UpdateObjectReply>(`/object/${uuid}`, params, req)
}

/**
 * @description "删除操作对象"
 * @param params
 */
export function deleteObject(params: components.DeleteObjectReqParams, uuid: string) {
	return ssoms.delete<components.DeleteObjectReply>(`/object/${uuid}`, params)
}

/**
 * @description "菜单选项，用于添加子菜单和操作时选择父级菜单"
 * @param params
 */
export function menuOptions(params: components.MenuOptionsReqParams) {
	return ssoms.get<components.MenuOptionsReply>(`/object/menuOptions`, params)
}

/**
 * @description "角色可以分配权限的对象"
 * @param params
 */
export function roleOperations(params: components.RoleOperationsReqParams) {
	return ssoms.get<components.RoleOperationsReply>(`/object/roleOperations`, params)
}

/**
 * @description "角色已分配的权限"
 * @param params
 */
export function rolePermissions(params: components.RolePermissionsReqParams, roleUUID: string) {
	return ssoms.get<components.RolePermissionsReply>(`/permission/role/${roleUUID}/Permissions`, params)
}

/**
 * @description "将对象执行操作的权限授予角色"
 * @param params
 * @param req
 */
export function grant(params: components.GrantReqParams, req: components.GrantReq, roleUUID: string) {
	return ssoms.put<components.GrantReply>(`/permission/role/${roleUUID}/grant`, params, req)
}
