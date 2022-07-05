import * as components from "./ssomsComponents"
import ajax from "@/utils/ajax"
export * from "./ssomsComponents"

/**
 * @description "用户列表"
 * @param req
 */
export function userList(req: components.UserListReq) {
	return ajax<components.UserListReply>("/user/list", req)
}

/**
 * @description "新增用户"
 * @param req
 */
export function addUser(req: components.AddUserReq) {
	return ajax<components.AddUserReply>("/user/add", req)
}

/**
 * @description "用户详情"
 * @param req
 */
export function userDetail(req: components.UserDetailReq) {
	return ajax<components.UserDetailReply>("/user/detail", req)
}

/**
 * @description "修改用户"
 * @param req
 */
export function updateUser(req: components.UpdateUserReq) {
	return ajax<components.UpdateUserReply>("/user/update", req)
}

/**
 * @description "删除用户"
 * @param req
 */
export function deleteUser(req: components.DeleteUserReq) {
	return ajax<components.DeleteUserReply>("/user/delete", req)
}

/**
 * @description "用户搜索下拉选项"
 * @param req
 */
export function userFilterOptions(req: components.UserFilterOptionsReq) {
	return ajax<components.UserFilterOptionsReply>("/user/filterOptions", req)
}

/**
 * @description "用户已分配的角色"
 * @param req
 */
export function assignedRoles(req: components.AssignedRolesReq) {
	return ajax<components.AssignedRolesReply>("/user/assignedRoles", req)
}

/**
 * @description "给用户分配角色"
 * @param req
 */
export function assignRole(req: components.AssignRoleReq) {
	return ajax<components.AssignRoleReply>("/user/assignRole", req)
}

/**
 * @description "用户通过其分配的角色获得的权限"
 * @param req
 */
export function userPermissions(req: components.UserPermissionsReq) {
	return ajax<components.UserPermissionsReply>("/user/permissions", req)
}

/**
 * @description "角色列表"
 * @param req
 */
export function roleList(req: components.RoleListReq) {
	return ajax<components.RoleListReply>("/role/list", req)
}

/**
 * @description "新增角色"
 * @param req
 */
export function addRole(req: components.AddRoleReq) {
	return ajax<components.AddRoleReply>("/role/add", req)
}

/**
 * @description "角色详情"
 * @param req
 */
export function roleDetail(req: components.RoleDetailReq) {
	return ajax<components.RoleDetailReply>("/role/detail", req)
}

/**
 * @description "修改角色"
 * @param req
 */
export function updateRole(req: components.UpdateRoleReq) {
	return ajax<components.UpdateRoleReply>("/role/update", req)
}

/**
 * @description "删除角色"
 * @param req
 */
export function deleteRole(req: components.DeleteRoleReq) {
	return ajax<components.DeleteRoleReply>("/role/delete", req)
}

/**
 * @description "角色已拥有的用户"
 */
export function roleOptions() {
	return ajax<components.OptionsReply>("/role/options")
}

/**
 * @description "角色已拥有的用户"
 * @param req
 */
export function assignedUsers(req: components.AssignedUsersReq) {
	return ajax<components.AssignedUsersReply>("/role/assignedUsers", req)
}

/**
 * @description "给角色添加用户"
 * @param req
 */
export function assignUser(req: components.AssignUserReq) {
	return ajax<components.AssignUserReply>("/role/assignUser", req)
}

/**
 * @description "给角色移出用户"
 * @param req
 */
export function deassignUser(req: components.DeassignUserReq) {
	return ajax<components.DeassignUserReply>("/role/deassignUser", req)
}

/**
 * @description "角色已继承的角色"
 * @param req
 */
export function inheritances(req: components.InheritancesReq) {
	return ajax<components.InheritancesReply>("/role/inheritances", req)
}

/**
 * @description "角色继承另一个角色"
 * @param req
 */
export function extendRole(req: components.AddInheritanceReq) {
	return ajax<components.AddInheritanceReply>("/role/addInheritance", req)
}

/**
 * @description "操作对象列表"
 * @param req
 */
export function objectList(req: components.ObjectListReq) {
	return ajax<components.ObjectListReply>("/object/list", req)
}

/**
 * @description "新增操作对象"
 * @param req
 */
export function addObject(req: components.AddObjectReq) {
	return ajax<components.AddObjectReply>("/object/add", req)
}

/**
 * @description "操作对象详情"
 * @param req
 */
export function objectDetail(req: components.ObjectDetailReq) {
	return ajax<components.ObjectDetailReply>("/object/detail", req)
}

/**
 * @description "修改操作对象"
 * @param req
 */
export function updateObject(req: components.UpdateObjectReq) {
	return ajax<components.UpdateObjectReply>("/object/update", req)
}

/**
 * @description "删除操作对象"
 * @param req
 */
export function deleteObject(req: components.DeleteObjectReq) {
	return ajax<components.DeleteObjectReply>("/object/delete", req)
}

/**
 * @description "角色可以分配权限的对象"
 * @param req
 */
export function roleOperations(req: components.RoleOperationsReq) {
	return ajax<components.RoleOperationsReply>("/object/roleOperations", req)
}

/**
 * @description "角色已分配的权限"
 * @param req
 */
export function rolePermissions(req: components.RolePermissionsReq) {
	return ajax<components.RolePermissionsReply>("/permission/RolePermissions", req)
}

/**
 * @description "将对象执行操作的权限授予角色"
 * @param req
 */
export function grant(req: components.GrantReq) {
	return ajax<components.GrantReply>("/permission/grant", req)
}
