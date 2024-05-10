// 自动生成模板WcDepartment
package weChat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
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

// GetAllFullDepartments 获取全层级部门名称列表
func GetAllFullDepartments() (list []FullDepartment, err error) {
	cacheKey := fmt.Sprintf("GetAllFullDepartments")
	cacheValue, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Result()
	if err != nil {
		//global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		fmt.Println("RedisStoreGetError:", err)
	}
	if cacheValue != "" {
		err = json.Unmarshal([]byte(cacheValue), &list)
	} else {
		var wds []WcDepartment
		err = global.GVA_DB.Find(&wds).Error
		if err != nil {
			fmt.Println("GetAllFullDepartments Error:", err)
			return list, err
		}
		for _, wd := range wds {
			id := int(wd.ID)
			fd := FullDepartment{
				Id:   id,
				Name: GetFullDepartmentById(id),
			}
			list = append(list, fd)
		}

		cacheByte, err := json.Marshal(&list)
		cacheValue = string(cacheByte)
		err = global.GVA_REDIS.Set(context.Background(), cacheKey, cacheValue, time.Hour*24).Err()
		if err != nil {
			//global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
			fmt.Println("RedisStoreSetError:", err)
			return list, err
		}
	}
	return
}
