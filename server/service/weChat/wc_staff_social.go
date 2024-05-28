package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WcStaffSocialService struct {
}

// CreateWcStaffSocial 创建社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) CreateWcStaffSocial(wcStaffSocial *weChat.WcStaffSocial) (err error) {
	err = global.GVA_DB.Create(wcStaffSocial).Error
	return err
}

// DeleteWcStaffSocial 删除社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) DeleteWcStaffSocial(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffSocial{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffSocialByIds 批量删除社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) DeleteWcStaffSocialByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffSocial{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffSocial 更新社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) UpdateWcStaffSocial(wcStaffSocial weChat.WcStaffSocial) (err error) {
	err = global.GVA_DB.Save(&wcStaffSocial).Error
	return err
}

// GetWcStaffSocial 根据ID获取社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) GetWcStaffSocial(ID string) (wcStaffSocial weChat2.WcStaffSocialResponse, err error) {
	var staffSocial weChat.WcStaffSocial
	err = global.GVA_DB.Where("id = ?", ID).First(&staffSocial).Error
	if err != nil {
		return
	}
	wcStaffSocial, err = wcStaffSocialService.AssembleStaffSocial(staffSocial)
	return
}

// GetWcStaffSocialInfoList 分页获取社保公积金管理记录
func (wcStaffSocialService *WcStaffSocialService) GetWcStaffSocialInfoList(info weChatReq.WcStaffSocialSearch) (list []weChat2.WcStaffSocialResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffSocial{})
	var wcStaffSocials []weChat.WcStaffSocial
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.StaffId != nil {
		db = db.Where("staff_id = ?", info.StaffId)
	}
	if info.Keyword != "" {
		keyword := "%" + info.Keyword + "%"
		db = db.Where("(account LIKE ? OR credential_number LIKE ?)", keyword, keyword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wcStaffSocials).Error
	if err != nil {
		return
	}
	list, err = wcStaffSocialService.AssembleStaffSocialList(wcStaffSocials)
	return
}

func (wcStaffSocialService *WcStaffSocialService) AssembleStaffSocialList(staffSocials []weChat.WcStaffSocial) (newStaffSocials []weChat2.WcStaffSocialResponse, err error) {
	var newStaffSocial weChat2.WcStaffSocialResponse
	configInfo := config.GetConfigInfo()

	for _, staffSocial := range staffSocials {
		newStaffSocial.WcStaffSocial = staffSocial
		socialTypeText, _ := utils.Find(configInfo.SocialType, *staffSocial.Type)
		newStaffSocial.TypeText = socialTypeText
		credentialTypeText, _ := utils.Find(configInfo.CredentialType, *staffSocial.CredentialType)
		newStaffSocial.CredentialTypeText = credentialTypeText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffSocial.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffSocialList Err:", err)
			return
		}
		newStaffSocial.StaffName = staff.Name
		newStaffSocial.JobNum = staff.JobNum

		newStaffSocials = append(newStaffSocials, newStaffSocial)
	}
	return
}

func (wcStaffSocialService *WcStaffSocialService) AssembleStaffSocial(staffSocial weChat.WcStaffSocial) (newStaffSocial weChat2.WcStaffSocialResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffSocial.WcStaffSocial = staffSocial
	socialTypeText, _ := utils.Find(configInfo.SocialType, *staffSocial.Type)
	newStaffSocial.TypeText = socialTypeText
	credentialTypeText, _ := utils.Find(configInfo.CredentialType, *staffSocial.CredentialType)
	newStaffSocial.CredentialTypeText = credentialTypeText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffSocial.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffSocial Err:", err)
		return
	}

	newStaffSocial.StaffName = staff.Name
	newStaffSocial.JobNum = staff.JobNum

	return
}
