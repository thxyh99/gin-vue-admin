// 自动生成模板WcDepartment
package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wcDepartment表 结构体  WcDepartment
type WcDepartment struct {
	global.GVA_MODEL
	DepartmentId     *int   `json:"departmentId" form:"departmentId" gorm:"column:department_id;comment:部门ID;size:6;" binding:"required"` //部门ID
	Name             string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:200;" binding:"required"`                        //部门名称
	DepartmentLeader string `json:"departmentLeader" form:"departmentLeader" gorm:"column:department_leader;comment:部门负责人;size:200;"`     //部门负责人
	Parentid         *int   `json:"parentid" form:"parentid" gorm:"column:parentid;comment:父部门ID;size:6;" binding:"required"`             //父部门ID
	Order            *int   `json:"order" form:"order" gorm:"column:order;comment:排序;size:6;" binding:"required"`                         //排序
}

type FullDepartment struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// TableName wcDepartment表 WcDepartment自定义表名 wc_department
func (WcDepartment) TableName() string {
	return "wc_department"
}

// GetFullDepartmentById 通过ID获取员工具体部门信息
func GetFullDepartmentById(ID int) string {
	var wd WcDepartment
	err := global.GVA_DB.Where("id = ?", ID).First(&wd).Error
	if err != nil {
		return ""
	}
	parentId := *wd.Parentid
	//fmt.Println("parentId", parentId)
	if parentId == 0 {
		return wd.Name
	}
	parentName := GetFullDepartmentById(parentId)
	//fmt.Println("parentName", parentName)

	return parentName + "/" + wd.Name
}

// GetAllFullDepartments 获取全层级部门名称列表
func GetAllFullDepartments() (list []FullDepartment, err error) {
	var wds []WcDepartment
	err = global.GVA_DB.Find(&wds).Error
	if err != nil {
		fmt.Println("GetAllFullDepartments Error:", err)
		return
	}
	for _, wd := range wds {
		id := int(wd.ID)
		fd := FullDepartment{
			Id:   id,
			Name: GetFullDepartmentById(id),
		}
		list = append(list, fd)
	}
	return
}
