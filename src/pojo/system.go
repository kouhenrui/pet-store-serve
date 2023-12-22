package pojo

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
}

type Class struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
}

// 菜单
type Menu struct {
	gorm.Model
	// 组件路径
	Component string `json:"component" gorm:"component"`
	// 菜单图标
	Icon string `json:"icon,omitempty" gorm:"icon,omitempty"`
	// 是否缓存（0存 1不存）
	IsCache bool `json:"is_cache" gorm:"is_cache"`
	// 是否为外链（0是 1否）
	IsFrame bool `json:"is_frame" gorm:"is_frame"`
	// 菜单名称
	MenuName string `json:"menu_name" gorm:"menu_name"`
	// 类型（M目录 C菜单 F按钮）
	MenuType Menutype `json:"menu_type" gorm:"menu_type"`
	// 显示顺序
	OrderNum int64 `json:"order_num" gorm:"order_num"`
	// 父菜单ID
	ParentId uint `json:"parent_id,omitempty" gorm:"parent_id,omitempty"`
	// 父菜单名称
	ParentName string `json:"parent_name,omitempty" gorm:"parent_name,omitempty"`
	// 路由地址
	Path string `json:"path" gorm:"path"`
	// 权限字符串
	Perms string `json:"perms,omitempty" gorm:"perms,omitempty"`
	// 菜单状态（0正常 1停用）
	Status bool `json:"status" gorm:"status"`
	// 显示状态（0显示，1隐藏）
	Visible bool `json:"visible" gorm:"visible"`

	Remark string `json:"remark" gorm:"remark"`
}

// 类型（M目录 C菜单 F按钮）
type Menutype string
type Routers struct {
	gorm.Model

	Name string `json:"name" gorm:"name"`
	// 组件路径
	Component string `json:"component,omitempty" gorm:"component,omitempty"`
	// 菜单图标
	Icon string `json:"icon,omitempty" gorm:"icon,omitempty"`
	// 类型（M目录 C菜单 F按钮）
	MenuType Menutype `json:"type" gorm:"type"`
	// 显示顺序
	OrderNum int64 `json:"order_num" gorm:"order_num"`

	// 父菜单ID
	ParentId uint `json:"parent_id,omitempty" gorm:"parent_id,omitempty"`

	// 路由地址
	Path string `json:"path,omitempty" gorm:"path,omitempty"`
	// 权限字符串
	Permission string `json:"permission,omitempty" gorm:"permission,omitempty"`
	// 菜单状态（0正常 1停用）
	Status bool `json:"status" gorm:"status"`
	// 显示状态（0显示，1隐藏）
	Visible bool `json:"visible" gorm:"visible"`

	Remark string `json:"remark,omitempty" gorm:"remark,omitempty"`
}
