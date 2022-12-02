// Code generated by goctl. DO NOT EDIT.
package types

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Extra string `json:"extra,omitempty"`
}

type OptionWithDisabled struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Extra    string `json:"extra,omitempty"`
	Disabled bool   `json:"disabled"`
}

type Statistic struct {
	Month string `json:"month"`
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type StatisticReply struct {
	RoleAmount       int64       `json:"roleAmount"`       // 角色数量
	UserAmount       int64       `json:"userAmount"`       // 用户数量
	SystemAmount     int64       `json:"systemAmount"`     // 系统数量
	MenuAmount       int64       `json:"menuAmount"`       // 菜单数量
	ActionAmount     int64       `json:"actionAmount"`     // 操作数量
	PermissionAmount int64       `json:"permissionAmount"` // 授权数量
	Statistics       []Statistic `json:"statistics"`       // 菜单&操作统计
}

type UserListReq struct {
	Name     string `form:"name,optional"`
	Mobile   string `form:"mobile,optional"`
	Page     int64  `form:"page"`
	PageSize int64  `form:"pageSize,default=20"`
}

type User struct {
	UUID   string   `json:"uuid"`
	Name   string   `json:"name"`
	Mobile string   `json:"mobile"`
	Avatar string   `json:"avatar"`
	Gender int64    `json:"gender"`
	Status int64    `json:"status"`
	Roles  []Option `json:"roles"`
}

type UserListReply struct {
	Total int64  `json:"total"`
	List  []User `json:"list"`
}

type UserForm struct {
	UUID         string `json:"uuid,optional"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	Password     string `json:"password,optional"`
	Avatar       string `json:"avatar"`
	Gender       int64  `json:"gender"`
	Birth        string `json:"birth"`
	Introduction string `json:"introduction"`
	Status       int64  `json:"status,default=1"`
}

type AddUserReply struct {
	UUID string `json:"uuid"`
}

type UserDetailReq struct {
	UUID string `path:"uuid"`
}

type UpdateUserReq struct {
	UUID         string `path:"uuid"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	Password     string `json:"password,optional"`
	Avatar       string `json:"avatar"`
	Gender       int64  `json:"gender"`
	Birth        string `json:"birth"`
	Introduction string `json:"introduction"`
	Status       int64  `json:"status,default=1"`
}

type UpdateUserReply struct {
	Success bool `json:"success"`
}

type DeleteUserReq struct {
	UUID string `path:"uuid"`
}

type DeleteUserReply struct {
	Success bool `json:"success"`
}

type UserFilterOptionsReq struct {
	Name string `form:"name,optional"`
}

type UserFilterOptionsReply struct {
	Options []Option `json:"options"`
}

type AssignedRolesReq struct {
	UUID string `form:"uuid"`
}

type AssignedRolesReply struct {
	UUID     string   `json:"uuid"`
	Name     string   `json:"name"`
	Assigned []Option `json:"assigned"`
	Options  []Option `json:"options"`
}

type AssignRoleReq struct {
	UUID          string   `json:"uuid"`
	RoleUUIDArray []string `json:"roleUUIDArray"`
}

type AssignRoleReply struct {
	Success bool `json:"success"`
}

type UserPermissionsReq struct {
	UUID string `json:"uuid"`
	Typ  int64  `json:"type"` // 类型: 角色
}

type UserPermissionsReply struct {
	Systems []Option `json:"system"`
}

type AvatarUploadReq struct {
	Avatar string `json:"avatar"`
}

type AvatarUploadReply struct {
	Success bool `json:"success"`
}

type InfoEditReq struct {
	Introduction string `json:"introduction"`
}

type InfoEditReply struct {
	Success bool `json:"success"`
}

type PasswordResetReq struct {
	OldPassword     string `json:"oldPassword"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type PasswordResetReply struct {
	Success bool `json:"success"`
}

type RoleListReq struct {
	RoleName string `form:"roleName,optional"`
	Page     int64  `form:"page"`
	PageSize int64  `form:"pageSize,default=20"`
}

type Role struct {
	RoleUUID     string   `json:"roleUUID"`     // 角色uuid
	RoleName     string   `json:"roleName"`     // 角色名称
	Summary      string   `json:"summary"`      // 角色概述
	Inheritances []Option `json:"inheritances"` // 继承的角色
	UsersAmount  int64    `json:"usersAmount"`  // 拥有的用户数量
}

type RoleListReply struct {
	Total int64  `json:"total"`
	List  []Role `json:"list"`
}

type RoleForm struct {
	RoleUUID string `json:"roleUUID,optional"`
	RoleName string `json:"roleName"`
	Summary  string `json:"summary"`
}

type AddRoleReply struct {
	RoleUUID string `json:"roleUUID"`
}

type RoleDetailReq struct {
	RoleUUID string `path:"roleUUID"`
}

type UpdateRoleReq struct {
	RoleUUID string `path:"roleUUID"`
	RoleName string `json:"roleName"`
	Summary  string `json:"summary"`
}

type UpdateRoleReply struct {
	Success bool `json:"success"`
}

type DeleteRoleReq struct {
	RoleUUID string `path:"roleUUID"`
}

type DeleteRoleReply struct {
	Success bool `json:"success"`
}

type OptionsReply struct {
	Options []Option `json:"options"`
}

type AssignedUsersReq struct {
	RoleUUID string `form:"roleUUID"`
	Page     int64  `form:"page"`
	PageSize int64  `form:"pageSize,default=20"`
}

type UserOtion struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar"`
	Gender   int64  `json:"gender"`
	Status   int64  `json:"status"`
	IsDelete int64  `json:"isDelete"`
}

type AssignedUsersReply struct {
	Total    int64       `json:"total"`
	List     []UserOtion `json:"list"`
	RoleName string      `json:"roleName"`
}

type AssignUserReq struct {
	RoleUUID string `path:"roleUUID"`
	UserUUID string `json:"userUUID"`
}

type AssignUserReply struct {
	Success bool `json:"success"`
}

type DeassignUserReq struct {
	RoleUUID string `path:"roleUUID"`
	UserUUID string `json:"userUUID"`
}

type DeassignUserReply struct {
	Success bool `json:"success"`
}

type InheritancesReq struct {
	RoleUUID string `path:"roleUUID"`
}

type InheritancesReply struct {
	Roles []Option `json:"roles"`
}

type AddInheritanceReq struct {
	RoleUUID                string   `path:"roleUUID"`
	AddInheritanceUUIDArray []string `json:"extendedRoleUUIDArray"`
}

type AddInheritanceReply struct {
	Success bool `json:"success"`
}

type SystemListReq struct {
	SystemCode string `form:"systemCode,optional"`
	SystemName string `form:"systemName,optional"`
}

type System struct {
	UUID       string `json:"uuid"`
	SystemCode string `json:"systemCode"`
	SystemName string `json:"systemName"`
	Domain     string `json:"domain"`
	Sort       int64  `json:"sort"`
	SubType    int64  `json:"subType"`        // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Extra      string `json:"extra,optional"` // 扩展字段
	Icon       string `json:"icon"`           // 图标
	Status     int64  `json:"status"`
}

type SystemListReply struct {
	List []*System `json:"list"`
}

type SystemForm struct {
	UUID       string `json:"uuid,optional"`
	SystemCode string `json:"systemCode"`
	SystemName string `json:"systemName"`
	Domain     string `json:"domain"`
	Sort       int64  `json:"sort"`
	SubType    int64  `json:"subType"` // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Icon       string `json:"icon"`    // 图标
	Status     int64  `json:"status"`
}

type AddSystemReply struct {
	UUID string `json:"uuid"`
}

type SystemDetailReq struct {
	UUID string `path:"uuid"`
}

type UpdateSystemReq struct {
	UUID       string `path:"uuid"`
	SystemCode string `json:"systemCode"`
	SystemName string `json:"systemName"`
	Domain     string `json:"domain"`
	Sort       int64  `json:"sort"`
	SubType    int64  `json:"subType"` // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Icon       string `json:"icon"`    // 图标
	Status     int64  `json:"status"`
}

type UpdateSystemReply struct {
	Success bool `json:"success"`
}

type DeleteSystemReq struct {
	UUID string `path:"uuid"`
}

type DeleteSystemReply struct {
	Success bool `json:"success"`
}

type ObjectListReq struct {
	TopKey     string `form:"topKey,optional"` // 传systemCode
	ObjectName string `form:"objectName,optional"`
	Key        string `form:"key,optional"`
}

type Object struct {
	UUID       string    `json:"uuid"`
	ObjectName string    `json:"objectName"`
	Identifier string    `json:"identifier"`
	Key        string    `json:"key"` // 操作对象的systemCode, 菜单的path, 操作的uri
	Sort       int64     `json:"sort"`
	Typ        int64     `json:"type"`           // 类型, 1操作对象, 2菜单，3操作(接口)
	SubType    int64     `json:"subType"`        // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Extra      string    `json:"extra,optional"` // 扩展字段
	Icon       string    `json:"icon"`           // 图标
	Status     int64     `json:"status"`
	PUUID      string    `json:"pUUID,optional"`
	Children   []*Object `json:"children,optional"`
}

type ObjectOption struct {
	Value    string          `json:"value"`
	Label    string          `json:"label"`
	PUUID    string          `json:"pUUID,optional"`
	Typ      int64           `json:"type"`              // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	SubType  int64           `json:"subType"`           // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Children []*ObjectOption `json:"children,optional"` // 子菜单
	Apis     []*ObjectOption `json:"apis"`              // 操作
}

type MenuOptionsReq struct {
	ExcludeHide bool `form:"excludeHide"` // 是否排查隐藏菜单
}

type MenuOption struct {
	Value    string        `json:"value"`
	Label    string        `json:"label"`
	PUUID    string        `json:"pUUID,optional"`
	Children []*MenuOption `json:"children,optional"` // 子菜单
}

type MenuOptionsReply struct {
	List []*MenuOption `json:"list"`
}

type ObjectListReply struct {
	List []*Object `json:"list"`
}

type ObjectForm struct {
	UUID       string `json:"uuid,optional"`
	ObjectName string `json:"objectName"`
	Identifier string `json:"identifier"`
	Key        string `json:"key"` // 操作对象的systemCode, 菜单的path, 操作的uri
	Sort       int64  `json:"sort"`
	Typ        int64  `json:"type"`           // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	SubType    int64  `json:"subType"`        // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Extra      string `json:"extra,optional"` // 扩展字段
	Icon       string `json:"icon"`           // 图标
	Status     int64  `json:"status"`
	PUUID      string `json:"pUUID,optional"`
	TopKey     string `json:"topKey,optional"` // 传systemCode
}

type AddObjectReply struct {
	UUID string `json:"uuid"`
}

type ImportObjectReq struct {
	UUID       string `json:"uuid"`
	ObjectName string `json:"objectName"`
	Identifier string `json:"identifier,optional"`
	Key        string `json:"key"` // 操作对象的systemCode, 菜单的path, 操作的uri
	Sort       int64  `json:"sort"`
	Typ        int64  `json:"type"`           // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	SubType    int64  `json:"subType"`        // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Extra      string `json:"extra,optional"` // 扩展字段
	Icon       string `json:"icon,optional"`  // 图标
	Status     int64  `json:"status"`
	PUUID      string `json:"pUUID,optional"`
	TopKey     string `json:"topKey"` // 传systemCode
}

type ImportObjectReply struct {
	Status string `json:"status"` // 导入结果状态
	Msg    string `json:"msg"`    // 导入结果信息
}

type ObjectDetailReq struct {
	UUID string `path:"uuid"`
}

type UpdateObjectReq struct {
	UUID       string `path:"uuid"`
	ObjectName string `json:"objectName"`
	Identifier string `json:"identifier"`
	Key        string `json:"key"` // 操作对象的systemCode, 菜单的path, 操作的uri
	Sort       int64  `json:"sort"`
	Typ        int64  `json:"type"`           // 类型, 1系统, 2菜单, 3操作(接口)
	SubType    int64  `json:"subType"`        // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	Extra      string `json:"extra,optional"` // 扩展字段
	Icon       string `json:"icon"`           // 图标
	Status     int64  `json:"status"`
	PUUID      string `json:"pUUID,optional"`
	TopKey     string `json:"topKey,optional"` // 更新时无法修改该值, 但是更新前需要使用该值校验对象是否已存在
}

type UpdateObjectReply struct {
	Success bool `json:"success"`
}

type DeleteObjectReq struct {
	UUID string `path:"uuid"`
}

type DeleteObjectReply struct {
	Success bool `json:"success"`
}

type RoleOperationsReq struct {
	TopKey string `form:"topKey,optional"` // 传systemCode
}

type RoleOperationsReply struct {
	List []*ObjectOption `json:"list,optional"`
}

type RolePermissionsReq struct {
	RoleUUID string `path:"roleUUID"`
	TopKey   string `form:"topKey"`
}

type RolePermissionsReply struct {
	RoleName        string   `json:"roleName"`
	MenuUUIDArray   []string `json:"menuUUIDArray"`
	ActionUUIDArray []string `json:"actionUUIDArray"`
}

type GrantReq struct {
	RoleUUID        string   `path:"roleUUID"`
	TopKey          string   `json:"topKey"`
	MenuUUIDArray   []string `json:"menuUUIDArray"`
	ActionUUIDArray []string `json:"actionUUIDArray"`
}

type GrantReply struct {
	Success bool `json:"success"`
}
