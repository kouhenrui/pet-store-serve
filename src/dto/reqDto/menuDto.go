package reqDto

type MenuList struct {
	Take       int    `json:"take,omitempty" validate:"required"`
	Skip       int    `json:"skip,omitempty" validate:"required"`
	Name       string `json:"name,omitempty"`
	ParentName string `json:"parent_name,omitempty"`
}
type MenuUpd struct {
	Id        uint   `json:"id" validate:"required"`
	Component string `json:"component" `
	// 菜单图标
	Icon string `json:"icon,omitempty" `
	// 是否缓存（0存 1不存）
	IsCache bool `json:"is_cache" `
	// 是否为外链（0是 1否）
	IsFrame bool `json:"is_frame" `
	// 菜单名称
	MenuName string `json:"menu_name" `
	// 类型（M目录 C菜单 F按钮）
	MenuType Menutype `json:"menu_type" `
	// 显示顺序
	OrderNum int64 `json:"order_num" `
	// 父菜单ID
	ParentId uint `json:"parent_id,omitempty" `
	// 父菜单名称
	ParentName string `json:"parent_name,omitempty" `
	// 路由地址
	Path string `json:"path" `
	// 权限字符串
	Perms string `json:"perms,omitempty" `
	// 菜单状态（0正常 1停用）
	Status bool `json:"status" `
	// 显示状态（0显示，1隐藏）
	Visible bool `json:"visible" `

	Remark string `json:"remark" `
}

type MenuAdd struct {
	Component string `json:"component" `
	// 菜单图标
	Icon string `json:"icon,omitempty" `
	// 是否缓存（0存 1不存）
	IsCache bool `json:"is_cache" `
	// 是否为外链（0是 1否）
	IsFrame bool `json:"is_frame" `
	// 菜单名称
	MenuName string `json:"menu_name" `
	// 类型（M目录 C菜单 F按钮）
	MenuType Menutype `json:"menu_type" `
	// 显示顺序
	OrderNum int64 `json:"order_num" `
	// 父菜单ID
	ParentId uint `json:"parent_id,omitempty" `
	// 父菜单名称
	ParentName string `json:"parent_name,omitempty" `
	// 路由地址
	Path string `json:"path" `
	// 权限字符串
	Perms string `json:"perms,omitempty" `
	// 菜单状态（0正常 1停用）
	Status bool `json:"status" `
	// 显示状态（0显示，1隐藏）
	Visible bool `json:"visible" `

	Remark string `json:"remark" `
}
