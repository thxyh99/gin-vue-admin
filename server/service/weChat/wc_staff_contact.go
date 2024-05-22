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

type WcStaffContactService struct {
}

// CreateWcStaffContact 创建紧急联系人记录
func (wcStaffContactService *WcStaffContactService) CreateWcStaffContact(wcStaffContact *weChat.WcStaffContact) (err error) {
	err = global.GVA_DB.Create(wcStaffContact).Error
	return err
}

// DeleteWcStaffContact 删除紧急联系人记录
func (wcStaffContactService *WcStaffContactService) DeleteWcStaffContact(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffContact{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffContactByIds 批量删除紧急联系人记录
func (wcStaffContactService *WcStaffContactService) DeleteWcStaffContactByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffContact{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffContact 更新紧急联系人记录
func (wcStaffContactService *WcStaffContactService) UpdateWcStaffContact(wcStaffContact weChat.WcStaffContact) (err error) {
	err = global.GVA_DB.Save(&wcStaffContact).Error
	return err
}

// GetWcStaffContact 根据员工ID获取紧急联系人记录
func (wcStaffContactService *WcStaffContactService) GetWcStaffContact(ID string) (newStaffContact weChat2.WcStaffContactResponse, err error) {
	var staffContact weChat.WcStaffContact
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&staffContact).Error
	if err != nil {
		return
	}

	newStaffContact, err = wcStaffContactService.AssembleStaffContact(staffContact)
	return
}

// GetWcStaffContactInfoList 分页获取紧急联系人记录
func (wcStaffContactService *WcStaffContactService) GetWcStaffContactInfoList(info weChatReq.WcStaffContactSearch) (list []weChat2.WcStaffContactResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffContact{})
	var wcStaffContacts []weChat.WcStaffContact
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

	err = db.Find(&wcStaffContacts).Error
	if err != nil {
		return
	}

	list, err = wcStaffContactService.AssembleStaffContactList(wcStaffContacts)
	return
}

func (wcStaffContactService *WcStaffContactService) AssembleStaffContactList(staffContacts []weChat.WcStaffContact) (newStaffContacts []weChat2.WcStaffContactResponse, err error) {
	var newStaffContact weChat2.WcStaffContactResponse
	configInfo := config.GetConfigInfo()

	for _, staffContact := range staffContacts {
		newStaffContact.WcStaffContact = staffContact
		relationshipText, _ := utils.Find(configInfo.Relationship, *staffContact.Relationship)
		newStaffContact.RelationshipText = relationshipText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffContact.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffContactList Err:", err)
			return
		}
		newStaffContact.StaffName = staff.Name
		newStaffContact.JobNum = staff.JobNum

		newStaffContacts = append(newStaffContacts, newStaffContact)
	}
	return
}

func (wcStaffContactService *WcStaffContactService) AssembleStaffContact(staffContact weChat.WcStaffContact) (newStaffContact weChat2.WcStaffContactResponse, err error) {
	configInfo := config.GetConfigInfo()
	newStaffContact.WcStaffContact = staffContact
	relationshipText, _ := utils.Find(configInfo.Relationship, *staffContact.Relationship)
	newStaffContact.RelationshipText = relationshipText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffContact.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffContact Err:", err)
		return
	}

	newStaffContact.StaffName = staff.Name
	newStaffContact.JobNum = staff.JobNum

	return
}
