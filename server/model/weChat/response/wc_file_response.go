package weChat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
	"time"
)

// WcFileResponse  文件结构体
type WcFileResponse struct {
	weChat.WcFile
	StaffName        string `json:"staffName"`        //员工名称
	JobNum           string `json:"jobNum"`           //员工工号
	TypeText         string `json:"typeText"`         //文件分类
	PrivateAccessURL string `json:"privateAccessURL"` //私有空间文件下载的 URL
}

func (WcFileResponse) AssembleFileList(files []weChat.WcFile) (newFiles []WcFileResponse, err error) {
	configInfo := config.GetConfigInfo()
	var newFile WcFileResponse
	qiNiuConfig := global.GVA_CONFIG.Qiniu
	mac := auth.New(qiNiuConfig.AccessKey, qiNiuConfig.SecretKey)
	for _, file := range files {
		newFile.WcFile = file
		//文件类型
		fileTypeText, _ := utils.Find(configInfo.FileType, *file.Type)
		newFile.TypeText = fileTypeText

		//获取员工名称工号
		var staff weChat.WcStaff
		err = global.GVA_DB.Table(staff.TableName()).Where("id=?", file.StaffId).First(&staff).Error
		if err != nil {
			fmt.Println("AssembleStaffBank Err:", err)
			return
		}
		newFile.StaffName = staff.Name
		newFile.JobNum = staff.JobNum

		//私有空间文件下载的 URL
		deadline := time.Now().Add(time.Second * 3600).Unix()
		newFile.PrivateAccessURL = storage.MakePrivateURL(mac, qiNiuConfig.ImgPath, file.Key, deadline)

		newFiles = append(newFiles, newFile)
	}
	return
}

func (WcFileResponse) AssembleFile(file weChat.WcFile) (newFile WcFileResponse, err error) {
	configInfo := config.GetConfigInfo()
	newFile.WcFile = file

	//文件类型
	fileTypeText, _ := utils.Find(configInfo.FileType, *file.Type)
	newFile.TypeText = fileTypeText

	//获取员工名称工号
	var staff weChat.WcStaff
	err = global.GVA_DB.Table(staff.TableName()).Where("id=?", file.StaffId).First(&staff).Error
	if err != nil {
		fmt.Println("AssembleStaffBank Err:", err)
		return
	}
	newFile.StaffName = staff.Name
	newFile.JobNum = staff.JobNum

	//私有空间文件下载的 URL
	qiNiuConfig := global.GVA_CONFIG.Qiniu
	mac := auth.New(qiNiuConfig.AccessKey, qiNiuConfig.SecretKey)
	deadline := time.Now().Add(time.Second * 3600).Unix()
	newFile.PrivateAccessURL = storage.MakePrivateURL(mac, qiNiuConfig.ImgPath, file.Key, deadline)

	return
}
