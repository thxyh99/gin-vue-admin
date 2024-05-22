package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
)

// WcFileResponse  文件结构体
type WcFileResponse struct {
	weChat.WcFile
	StaffName        string `json:"staffName"`        //员工名称
	JobNum           string `json:"jobNum"`           //员工工号
	TypeText         string `json:"typeText"`         //文件分类
	PrivateAccessURL string `json:"privateAccessURL"` //私有空间文件下载的 URL
}
