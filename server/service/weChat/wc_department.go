package weChat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"time"
)

type WcDepartmentService struct {
}

// CreateWcDepartment 创建wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) CreateWcDepartment(wcDepartment *weChat.WcDepartment) (err error) {
	err = global.GVA_DB.Create(wcDepartment).Error
	return err
}

// DeleteWcDepartment 删除wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) DeleteWcDepartment(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcDepartment{}, "id = ?", ID).Error
	return err
}

// DeleteWcDepartmentByIds 批量删除wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) DeleteWcDepartmentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcDepartment{}, "id in ?", IDs).Error
	return err
}

// UpdateWcDepartment 更新wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) UpdateWcDepartment(wcDepartment weChat.WcDepartment) (err error) {
	err = global.GVA_DB.Save(&wcDepartment).Error
	return err
}

// GetWcDepartment 根据ID获取wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) GetWcDepartment(ID string) (newDepartment weChat2.WcDepartmentResponse, err error) {
	var wcDepartment weChat.WcDepartment
	err = global.GVA_DB.Where("id = ?", ID).First(&wcDepartment).Error
	newDepartment.WcDepartment = wcDepartment
	newDepartment.FullName = weChat.GetFullDepartmentById(int(wcDepartment.ID))
	return
}

// GetWcDepartmentInfoList 分页获取wcDepartment表记录
func (wcDepartmentService *WcDepartmentService) GetWcDepartmentInfoList(info weChatReq.WcDepartmentSearch) (list []weChat2.WcDepartmentResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcDepartment{})
	var wcDepartments []weChat.WcDepartment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcDepartments).Error
	if err != nil {
		return
	}
	newWcDepartments := wcDepartmentService.AssembleDepartmentList(wcDepartments)
	return newWcDepartments, total, err
}

// GetAllFullDepartmentList 获取部门列表
func (wcDepartmentService *WcDepartmentService) GetAllFullDepartmentList() (list []weChat2.FullDepartment, total int64, err error) {
	list, err = wcDepartmentService.GetAllFullDepartments()

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("list", list)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	total = int64(len(list))

	return
}

// GetAllFullDepartments 获取全层级部门名称列表
func (wcDepartmentService *WcDepartmentService) GetAllFullDepartments() (list []weChat2.FullDepartment, err error) {
	cacheKey := fmt.Sprintf("GetAllFullDepartments")
	cacheValue, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Result()
	if err != nil {
		//global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		fmt.Println("RedisStoreGetError:", err)
	}
	if cacheValue != "" {
		err = json.Unmarshal([]byte(cacheValue), &list)
	} else {
		var wds []weChat.WcDepartment
		err = global.GVA_DB.Find(&wds).Error
		if err != nil {
			fmt.Println("GetAllFullDepartments Error:", err)
			return list, err
		}
		for _, wd := range wds {
			id := int(wd.ID)
			fd := weChat2.FullDepartment{
				Id:   id,
				Name: weChat.GetFullDepartmentById(id),
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

func (wcDepartmentService *WcDepartmentService) AssembleDepartmentList(departments []weChat.WcDepartment) (newDepartments []weChat2.WcDepartmentResponse) {
	var newDepartment weChat2.WcDepartmentResponse

	for _, department := range departments {
		newDepartment.WcDepartment = department
		newDepartment.FullName = weChat.GetFullDepartmentById(int(department.ID))
		newDepartments = append(newDepartments, newDepartment)
	}
	return
}
