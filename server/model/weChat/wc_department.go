// 自动生成模板WcDepartment
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// wcDepartment表 结构体  WcDepartment
type WcDepartment struct {
 global.GVA_MODEL 
      DepartmentId  *int `json:"departmentId" form:"departmentId" gorm:"column:department_id;comment:部门ID;size:6;" binding:"required"`  //部门ID 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:200;" binding:"required"`  //部门名称 
      NameEn  string `json:"nameEn" form:"nameEn" gorm:"column:name_en;comment:英文名称;size:200;"`  //英文名称 
      DepartmentLeader  string `json:"departmentLeader" form:"departmentLeader" gorm:"column:department_leader;comment:部门负责人;size:200;"`  //部门负责人 
      Parentid  *int `json:"parentid" form:"parentid" gorm:"column:parentid;comment:父部门ID;size:6;" binding:"required"`  //父部门ID 
      Order  *int `json:"order" form:"order" gorm:"column:order;comment:排序;size:6;" binding:"required"`  //排序 
}


// TableName wcDepartment表 WcDepartment自定义表名 wc_department
func (WcDepartment) TableName() string {
  return "wc_department"
}

