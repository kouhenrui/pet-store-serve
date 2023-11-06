package pojo

import (
	"gorm.io/gorm"
	"time"
)

type App struct {
	gorm.Model     `json:"gorm.Model"`
	AppName        string    `json:"appName,omitempty"`
	Version        string    `json:"version,omitempty"`
	AppType        string    `json:"appType,omitempty"`
	H5Url          string    `json:"h5Url,omitempty"`
	PackageUrl     string    `json:"packageUrl,omitempty"`
	PackageName    string    `json:"packageName,omitempty"`
	AppId          string    `json:"appId,omitempty"`
	AppletId       string    `json:"appletId,omitempty"`
	Charger        string    `json:"charger,omitempty"`
	Phone          string    `json:"phone,omitempty"`
	OnLineTime     time.Time `json:"onLineTime"`
	OnLineCharger  string    `json:"onLineCharger,omitempty"`
	AppDescription string    `json:"appDescription,omitempty"`
}
