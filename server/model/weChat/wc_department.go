// 自动生成模板WcDepartment
package weChat

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// wcDepartment表 结构体  WcDepartment
type WcDepartment struct {
	global.GVA_MODEL
	Name             string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:200;" binding:"required"`                    //部门名称
	DepartmentLeader string `json:"departmentLeader" form:"departmentLeader" gorm:"column:department_leader;comment:部门负责人;size:200;"` //部门负责人
	Parentid         *int   `json:"parentid" form:"parentid" gorm:"column:parentid;comment:父部门ID;size:6;" binding:"required"`         //父部门ID
	Order            *int   `json:"order" form:"order" gorm:"column:order;comment:排序;size:6;" binding:"required"`                     //排序
}

// TableName wcDepartment表 WcDepartment自定义表名 wc_department
func (WcDepartment) TableName() string {
	return "wc_department"
}

func GetFullDepartmentById(ID int) string {
	cacheKey := fmt.Sprintf("GetFullDepartmentById:%d", ID)
	cacheValue, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Result()
	if err != nil {
		//global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		fmt.Println("RedisStoreGetError:", err)
	}
	if cacheValue != "" {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("ID:", ID)
		fmt.Println("cacheValue:", cacheValue)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		return cacheValue
	} else {
		var wd WcDepartment
		err = global.GVA_DB.Where("id = ?", ID).First(&wd).Error
		if err != nil {
			return ""
		}
		parentId := *wd.Parentid
		if parentId == 0 {
			err = global.GVA_REDIS.Set(context.Background(), cacheKey, wd.Name, time.Hour*24).Err()
			if err != nil {
				//global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
				fmt.Println("RedisStoreSetError:", err)
				return ""
			}
			return wd.Name
		}
		parentName := GetFullDepartmentById(parentId)

		fullName := parentName + "/" + wd.Name

		err = global.GVA_REDIS.Set(context.Background(), cacheKey, fullName, time.Hour*24).Err()
		if err != nil {
			//global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
			fmt.Println("RedisStoreSetError:", err)
			return ""
		}
		return fullName
	}
}
