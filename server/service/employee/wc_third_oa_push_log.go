package employee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	employeeReq "github.com/flipped-aurora/gin-vue-admin/server/model/employee/request"
)

type WcThirdOaPushService struct {
}

// CreateWcThirdOaPush 创建OA回调日志记录
func (wcThirdOaPushService *WcThirdOaPushService) CreateWcThirdOaPush(wcThirdOaPush *employee.WcThirdOaPush) (err error) {
	err = global.GVA_DB.Create(wcThirdOaPush).Error
	return err
}

// DeleteWcThirdOaPush 删除OA回调日志记录
func (wcThirdOaPushService *WcThirdOaPushService) DeleteWcThirdOaPush(ID string) (err error) {
	err = global.GVA_DB.Delete(&employee.WcThirdOaPush{}, "id = ?", ID).Error
	return err
}

// DeleteWcThirdOaPushByIds 批量删除OA回调日志记录
func (wcThirdOaPushService *WcThirdOaPushService) DeleteWcThirdOaPushByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]employee.WcThirdOaPush{}, "id in ?", IDs).Error
	return err
}

// UpdateWcThirdOaPush 更新OA回调日志记录
func (wcThirdOaPushService *WcThirdOaPushService) UpdateWcThirdOaPush(wcThirdOaPush employee.WcThirdOaPush) (err error) {
	err = global.GVA_DB.Save(&wcThirdOaPush).Error
	return err
}

// GetWcThirdOaPush 根据ID获取OA回调日志记录
func (wcThirdOaPushService *WcThirdOaPushService) GetWcThirdOaPush(ID string) (wcThirdOaPush employee.WcThirdOaPush, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcThirdOaPush).Error
	return
}

// GetWcThirdOaPushInfoList 分页获取OA回调日志记录
func (wcThirdOaPushService *WcThirdOaPushService) GetWcThirdOaPushInfoList(info employeeReq.WcThirdOaPushSearch) (list []employee.WcThirdOaPush, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&employee.WcThirdOaPush{})
	var wcThirdOaPushs []employee.WcThirdOaPush
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

	err = db.Find(&wcThirdOaPushs).Error
	return wcThirdOaPushs, total, err
}

type OaBaseInfo struct {
	Id    string `json:"id"`
	Lunid string `json:"lunid" form:"lunid" gorm:"column:lunid;comment:;size:50;"` //lunid字段
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:100;"`   //name字段
	Type  string `json:"type" form:"type" gorm:"column:type;comment:;size:10;"`    //type字段
}

// 根据数据类型，数据名称获取oa数据id
func GetOaDataId(dataType string, dataName string) (oaUserId string, err error) {
	var baseInfo OaBaseInfo
	err = global.GVA_DB.Table("wc_staff_byoa").Where("type = ? and name = ?", dataType, dataName).Order("id desc").First(&baseInfo).Error
	if err != nil {
		return "", err
	}
	return baseInfo.Id, nil
}
