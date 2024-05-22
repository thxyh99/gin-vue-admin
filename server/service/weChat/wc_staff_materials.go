package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"strings"
)

type WcStaffMaterialsService struct {
}

// CreateWcStaffMaterials 创建证件材料记录
func (wcStaffMaterialsService *WcStaffMaterialsService) CreateWcStaffMaterials(wcStaffMaterials *weChat.WcStaffMaterials) (err error) {
	err = global.GVA_DB.Create(wcStaffMaterials).Error
	return err
}

// DeleteWcStaffMaterials 删除证件材料记录
func (wcStaffMaterialsService *WcStaffMaterialsService) DeleteWcStaffMaterials(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffMaterials{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffMaterialsByIds 批量删除证件材料记录
func (wcStaffMaterialsService *WcStaffMaterialsService) DeleteWcStaffMaterialsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffMaterials{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffMaterials 更新证件材料记录
func (wcStaffMaterialsService *WcStaffMaterialsService) UpdateWcStaffMaterials(wcStaffMaterials weChat.WcStaffMaterials) (err error) {
	err = global.GVA_DB.Save(&wcStaffMaterials).Error
	return err
}

// GetWcStaffMaterials 根据员工ID获取证件材料记录
func (wcStaffMaterialsService *WcStaffMaterialsService) GetWcStaffMaterials(ID string) (newStaffMaterials weChat2.WcStaffMaterialsResponseDiscard, err error) {
	var wcStaffMaterials weChat.WcStaffMaterials
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&wcStaffMaterials).Error
	if err != nil {
		return
	}

	newStaffMaterials, err = wcStaffMaterialsService.AssembleStaffMaterials(wcStaffMaterials)
	return
}

// GetWcStaffMaterialsInfoList 分页获取证件材料记录
func (wcStaffMaterialsService *WcStaffMaterialsService) GetWcStaffMaterialsInfoList(info weChatReq.WcStaffMaterialsSearch) (list []weChat2.WcStaffMaterialsResponseDiscard, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffMaterials{})
	var wcStaffMaterials []weChat.WcStaffMaterials
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

	err = db.Find(&wcStaffMaterials).Error
	if err != nil {
		return
	}

	list, err = wcStaffMaterialsService.AssembleStaffMaterialsList(wcStaffMaterials)
	return
}

func (wcStaffMaterialsService *WcStaffMaterialsService) AssembleStaffMaterialsList(staffMaterials []weChat.WcStaffMaterials) (newStaffMaterials []weChat2.WcStaffMaterialsResponseDiscard, err error) {
	var newStaffMater weChat2.WcStaffMaterialsResponseDiscard

	for _, staffMaterial := range staffMaterials {
		newStaffMater.WcStaffMaterials = staffMaterial

		if newStaffMater.SkillCertificate != "" {
			newStaffMater.SkillCertificateList = strings.Split(staffMaterial.SkillCertificate, "###")
		}

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffMaterial.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffMaterialsList Err:", err)
			return
		}
		newStaffMater.StaffName = staff.Name
		newStaffMater.JobNum = staff.JobNum

		newStaffMaterials = append(newStaffMaterials, newStaffMater)
	}
	return
}

func (wcStaffMaterialsService *WcStaffMaterialsService) AssembleStaffMaterials(staffMaterials weChat.WcStaffMaterials) (newStaffMaterials weChat2.WcStaffMaterialsResponseDiscard, err error) {
	newStaffMaterials.WcStaffMaterials = staffMaterials

	if staffMaterials.SkillCertificate != "" {
		newStaffMaterials.SkillCertificateList = strings.Split(staffMaterials.SkillCertificate, "###")
	}

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffMaterials.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffMaterials Err:", err)
		return
	}

	newStaffMaterials.StaffName = staff.Name
	newStaffMaterials.JobNum = staff.JobNum

	return
}
