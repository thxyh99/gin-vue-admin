package weChat

import "github.com/flipped-aurora/gin-vue-admin/server/model/weChat"

// WcDepartmentResponse 部门信息 结构体
type WcDepartmentResponse struct {
	weChat.WcDepartment
	FullName string `json:"fullName"` //员工名称
}

type FullDepartment struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
