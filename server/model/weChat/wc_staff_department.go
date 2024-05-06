// 自动生成模板WcStaffDepartment
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 员工部门管理 结构体  WcStaffDepartment
type WcStaffDepartment struct {
 global.GVA_MODEL 
      StaffId  *int `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`  //员工ID 
      DepartmentId  *int `json:"departmentId" form:"departmentId" gorm:"column:department_id;comment:成员所属部门ID;size:6;" binding:"required"`  //成员所属部门ID 
      IsMain  *bool `json:"isMain" form:"isMain" gorm:"column:is_main;comment:是否主部门(1:是 0:否);size:1;" binding:"required"`  //是否主部门(1:是 0:否) 
}


// TableName 员工部门管理 WcStaffDepartment自定义表名 wc_staff_department
func (WcStaffDepartment) TableName() string {
  return "wc_staff_department"
}

