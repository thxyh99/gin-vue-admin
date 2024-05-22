package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
)

type WcStaffBankService struct {
}

// CreateWcStaffBank 创建银行卡信息记录
func (wcStaffBankService *WcStaffBankService) CreateWcStaffBank(wcStaffBank *weChat.WcStaffBank) (err error) {
	err = global.GVA_DB.Create(wcStaffBank).Error
	return err
}

// DeleteWcStaffBank 删除银行卡信息记录
func (wcStaffBankService *WcStaffBankService) DeleteWcStaffBank(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcStaffBank{}, "id = ?", ID).Error
	return err
}

// DeleteWcStaffBankByIds 批量删除银行卡信息记录
func (wcStaffBankService *WcStaffBankService) DeleteWcStaffBankByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcStaffBank{}, "id in ?", IDs).Error
	return err
}

// UpdateWcStaffBank 更新银行卡信息记录
func (wcStaffBankService *WcStaffBankService) UpdateWcStaffBank(wcStaffBank weChat.WcStaffBank) (err error) {
	err = global.GVA_DB.Save(&wcStaffBank).Error
	return err
}

// GetWcStaffBank 根据员工ID获取银行卡信息记录
func (wcStaffBankService *WcStaffBankService) GetWcStaffBank(ID string) (newStaffBank weChat2.WcStaffBankResponse, err error) {
	var staffBank weChat.WcStaffBank
	err = global.GVA_DB.Where("staff_id = ?", ID).First(&staffBank).Error
	if err != nil {
		return
	}

	newStaffBank, err = wcStaffBankService.AssembleStaffBank(staffBank)
	return
}

// GetWcStaffBankInfoList 分页获取银行卡信息记录
func (wcStaffBankService *WcStaffBankService) GetWcStaffBankInfoList(info weChatReq.WcStaffBankSearch) (list []weChat2.WcStaffBankResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcStaffBank{})
	var wcStaffBanks []weChat.WcStaffBank
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

	err = db.Find(&wcStaffBanks).Error
	if err != nil {
		return
	}

	list, err = wcStaffBankService.AssembleStaffBankList(wcStaffBanks)
	return
}

func (wcStaffBankService *WcStaffBankService) AssembleStaffBankList(staffBanks []weChat.WcStaffBank) (newStaffBanks []weChat2.WcStaffBankResponse, err error) {
	var newStaffBank weChat2.WcStaffBankResponse

	for _, staffBank := range staffBanks {
		newStaffBank.WcStaffBank = staffBank

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffBank.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffBankList Err:", err)
			return
		}
		newStaffBank.StaffName = staff.Name
		newStaffBank.JobNum = staff.JobNum

		newStaffBanks = append(newStaffBanks, newStaffBank)
	}
	return
}

func (wcStaffBankService *WcStaffBankService) AssembleStaffBank(staffBank weChat.WcStaffBank) (newStaffBank weChat2.WcStaffBankResponse, err error) {
	newStaffBank.WcStaffBank = staffBank

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", staffBank.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffBank Err:", err)
		return
	}

	newStaffBank.StaffName = staff.Name
	newStaffBank.JobNum = staff.JobNum

	return
}
