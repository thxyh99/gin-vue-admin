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

type WcStaffEducationService struct {
}

// CreateWcStaffEducation 创建学历信息记录
func (wcStaffEducationService *WcStaffEducationService) CreateWcStaffEducation(wcStaffEducation *weChat.WcStaffEducation) (err error) {
	err = global.GVA_DB.Create(wcStaffEducation).Error
	return err
}

// DeleteWcStaffEducation 删除学历信息记录
func (wcStaffEducationService *WcStaffEducationService) DeleteWcStaffEducation(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffEducation{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffEducationByIds 批量删除学历信息记录
func (wcStaffEducationService *WcStaffEducationService) DeleteWcStaffEducationByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffEducation{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffEducation 更新学历信息记录
func (wcStaffEducationService *WcStaffEducationService) UpdateWcStaffEducation(wcStaffEducation weChat.WcStaffEducation) (err error) {
	err = global.GVA_DB.Save(&wcStaffEducation).Error
	return err
}

// GetWcStaffEducation 根据员工ID获取学历信息记录
func (wcStaffEducationService *WcStaffEducationService) GetWcStaffEducation(ID string) (newStaffEducation weChat2.WcStaffEducationResponse, err error) {
	var staffEducation weChat.WcStaffEducation
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&staffEducation).Error
	if err != nil {
		return
	}

	newStaffEducation, err = wcStaffEducationService.AssembleStaffEducation(staffEducation)
	return
}

// GetWcStaffEducationInfoList 分页获取学历信息记录
func (wcStaffEducationService *WcStaffEducationService) GetWcStaffEducationInfoList(info weChatReq.WcStaffEducationSearch) (list []weChat2.WcStaffEducationResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffEducation{})
	var wcStaffEducations []weChat.WcStaffEducation
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

	err = db.Find(&wcStaffEducations).Error
	if err != nil {
		return
	}

	list, err = wcStaffEducationService.AssembleStaffEducationList(wcStaffEducations)
	return
}

func (wcStaffEducationService *WcStaffEducationService) AssembleStaffEducationList(staffEducations []weChat.WcStaffEducation) (newStaffEducations []weChat2.WcStaffEducationResponse, err error) {
	var newStaffEducation weChat2.WcStaffEducationResponse
	configInfo := config.GetConfigInfo()

	for _, staffEducation := range staffEducations {
		newStaffEducation.WcStaffEducation = staffEducation
		educationText, _ := utils.Find(configInfo.Education, *staffEducation.Education)
		newStaffEducation.EducationText = educationText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffEducation.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffEducationList Err:", err)
			return
		}
		newStaffEducation.StaffName = staff.Name
		newStaffEducation.JobNum = staff.JobNum

		newStaffEducations = append(newStaffEducations, newStaffEducation)
	}
	return
}

func (wcStaffEducationService *WcStaffEducationService) AssembleStaffEducation(staffEducation weChat.WcStaffEducation) (newStaffEducation weChat2.WcStaffEducationResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffEducation.WcStaffEducation = staffEducation
	educationText, _ := utils.Find(configInfo.Education, *staffEducation.Education)
	newStaffEducation.EducationText = educationText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffEducation.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffEducation Err:", err)
		return
	}

	newStaffEducation.StaffName = staff.Name
	newStaffEducation.JobNum = staff.JobNum

	return
}
