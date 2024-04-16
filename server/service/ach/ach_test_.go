package ach

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ach"
    achReq "github.com/flipped-aurora/gin-vue-admin/server/model/ach/request"
)

type AchTestService struct {
}

// CreateAchTest 创建绩效测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (achTestService *AchTestService) CreateAchTest(achTest *ach.AchTest) (err error) {
	err = global.GVA_DB.Create(achTest).Error
	return err
}

// DeleteAchTest 删除绩效测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (achTestService *AchTestService)DeleteAchTest(ID string) (err error) {
	err = global.GVA_DB.Delete(&ach.AchTest{},"id = ?",ID).Error
	return err
}

// DeleteAchTestByIds 批量删除绩效测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (achTestService *AchTestService)DeleteAchTestByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ach.AchTest{},"id in ?",IDs).Error
	return err
}

// UpdateAchTest 更新绩效测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (achTestService *AchTestService)UpdateAchTest(achTest ach.AchTest) (err error) {
	err = global.GVA_DB.Save(&achTest).Error
	return err
}

// GetAchTest 根据ID获取绩效测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (achTestService *AchTestService)GetAchTest(ID string) (achTest ach.AchTest, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&achTest).Error
	return
}

// GetAchTestInfoList 分页获取绩效测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (achTestService *AchTestService)GetAchTestInfoList(info achReq.AchTestSearch) (list []ach.AchTest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ach.AchTest{})
    var achTests []ach.AchTest
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&achTests).Error
	return  achTests, total, err
}
