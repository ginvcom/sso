// Code generated by goctl. DO NOT EDIT.

export interface Option {
	label: string
	value: string
	extra?: string
}

export interface UserListReq {
}

export interface UserListReqParams {
	name: string
	mobile: string
	page: number
	pageSize: number
}

export interface User {
	uuid: string
	name: string
	mobile: string
	avatar: string
	gender: number
	status: number
	roles: Array<Option>
}

export interface UserListReply {
	total: number
	list: Array<User>
}

export interface UserForm {
	uuid?: string
	name: string
	mobile: string
	password?: string
	avatar: string
	gender: number
	birth: string
	status: number
}

export interface AddUserReply {
	uuid: string
}

export interface UserDetailReq {
}

export interface UserDetailReqParams {
	uuid: string
}

export interface UpdateUserReq {
	name: string
	mobile: string
	password?: string
	avatar: string
	gender: number
	birth: string
	status: number
}

export interface UpdateUserReqParams {
	uuid: string
}

export interface UpdateUserReply {
	success: boolean
}

export interface DeleteUserReq {
}

export interface DeleteUserReqParams {
	uuid: string
}

export interface DeleteUserReply {
	success: boolean
}

export interface UserFilterOptionsReq {
	name?: string
}

export interface UserFilterOptionsReply {
	options: Array<Option>
}

export interface AssignedRolesReq {
	uuid: string
}

export interface AssignedRolesReply {
	roleUUIDArray: Array<string>
}

export interface AssignRoleReq {
	uuid: string
	roleUUIDArray: Array<string>
}

export interface AssignRoleReply {
	success: boolean
}

export interface UserPermissionsReq {
	uuid: string
	type: number // 类型: 角色
}

export interface UserPermissionsReply {
	system: Array<Option>
}

export interface RoleListReq {
}

export interface RoleListReqParams {
	roleName: string
	page: number
	pageSize: number
}

export interface Role {
	roleUUID: string // 角色uuid
	roleName: string // 角色名称
	summary: string // 角色概述
	inheritances: Array<Option> // 继承的角色
	usersAmount: number // 拥有的用户数量
}

export interface RoleListReply {
	total: number
	list: Array<Role>
}

export interface RoleForm {
	roleUUID?: string
	roleName: string
	summary: string
}

export interface AddRoleReply {
	roleUUID: string
}

export interface RoleDetailReq {
}

export interface RoleDetailReqParams {
	roleUUID: string
}

export interface UpdateRoleReq {
	roleName: string
	summary: string
}

export interface UpdateRoleReqParams {
	roleUUID: string
}

export interface UpdateRoleReply {
	success: boolean
}

export interface DeleteRoleReq {
}

export interface DeleteRoleReqParams {
	roleUUID: string
}

export interface DeleteRoleReply {
	success: boolean
}

export interface OptionsReply {
	options: Array<Option>
}

export interface AssignedUsersReq {
	roleUUID: string
}

export interface AssignedUsersReply {
	users: Array<Option>
}

export interface AssignUserReq {
	userUUID: string
}

export interface AssignUserReqParams {
	roleUUID: string
}

export interface AssignUserReply {
	success: boolean
}

export interface DeassignUserReq {
	userUUID: string
}

export interface DeassignUserReqParams {
	roleUUID: string
}

export interface DeassignUserReply {
	success: boolean
}

export interface InheritancesReq {
}

export interface InheritancesReqParams {
	roleUUID: string
}

export interface InheritancesReply {
	roles: Array<Option>
}

export interface AddInheritanceReq {
	extendedRoleUUIDArray: Array<string>
}

export interface AddInheritanceReqParams {
	roleUUID: string
}

export interface AddInheritanceReply {
	success: boolean
}

export interface ObjectListReq {
}

export interface ObjectListReqParams {
	topKey: string // 传systemCode
	objectName: string
	key: string
}

export interface Object {
	uuid: string
	objectName: string
	domain: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	icon: string // 图标
	status: number
	pUUID?: string
	children?: Array<Object>
}

export interface ObjectListReply {
	list: Array<Object>
}

export interface ObjectForm {
	uuid?: string
	objectName: string
	domain: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	icon: string // 图标
	status: number
	pUUID?: string
	topKey?: string // 传systemCode, 更新的时候不传(更新时无法修改该值)
}

export interface AddObjectReply {
	uuid: string
}

export interface ObjectDetailReq {
}

export interface ObjectDetailReqParams {
	uuid: string
}

export interface UpdateObjectReq {
	objectName: string
	domain: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	icon: string // 图标
	status: number
	pUUID?: string
}

export interface UpdateObjectReqParams {
	uuid: string
}

export interface UpdateObjectReply {
	success: boolean
}

export interface DeleteObjectReq {
}

export interface DeleteObjectReqParams {
	uuid: string
}

export interface DeleteObjectReply {
	success: boolean
}

export interface RoleOperationsReq {
}

export interface RoleOperationsReqParams {
	topKey: string // 传systemCode
	pUUID: string
}

export interface RoleOperationsReply {
	list?: Array<Object>
}

export interface RolePermissionsReq {
}

export interface RolePermissionsReqParams {
	roleUUID: string
	typeName: string // system, others
	pUUID: string
}

export interface ObjectOption {
	label: string
	value: string
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	extra: string
	children: Array<ObjectOption>
}

export interface RolePermissionsReply {
	list: Array<ObjectOption>
}

export interface GrantReq {
	typeName: string // system, others
	pUUID?: string
}

export interface GrantReply {
	uuidArray: Array<string>
}

