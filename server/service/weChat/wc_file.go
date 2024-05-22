package weChat

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
	"strings"
	"time"
)

type WcFileService struct {
}

// CreateWcFile 创建文件管理记录
func (wcFileService *WcFileService) CreateWcFile(wcFile *weChat.WcFile) (err error) {
	err = global.GVA_DB.Create(wcFile).Error
	return err
}

// DeleteWcFile 删除文件管理记录
func (wcFileService *WcFileService) DeleteWcFile(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcFile{}, "id = ?", ID).Error
	return err
}

// DeleteWcFileByIds 批量删除文件管理记录
func (wcFileService *WcFileService) DeleteWcFileByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcFile{}, "id in ?", IDs).Error
	return err
}

// UpdateWcFile 更新文件管理记录
func (wcFileService *WcFileService) UpdateWcFile(wcFile weChat.WcFile) (err error) {
	err = global.GVA_DB.Save(&wcFile).Error
	return err
}

// GetWcFile 根据ID获取文件管理记录
func (wcFileService *WcFileService) GetWcFile(ID string) (wcFile weChat.WcFile, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wcFile).Error
	return
}

// GetWcFileInfoList 分页获取文件管理记录
func (wcFileService *WcFileService) GetWcFileInfoList(info weChatReq.WcFileSearch) (list []weChat.WcFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcFile{})
	var wcFiles []weChat.WcFile
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

	err = db.Find(&wcFiles).Error
	return wcFiles, total, err
}

func (wcFileService *WcFileService) Upload(file weChat.WcFile) (weChat.WcFile, error) {
	err := global.GVA_DB.Create(&file).Error
	return file, err
}

func (wcFileService *WcFileService) FindFile(ID uint) (newFile weChat2.WcFileResponse, err error) {
	var file weChat.WcFile
	err = global.GVA_DB.Where("id=?", ID).First(&file).Error
	if err != nil {
		return newFile, err
	}

	newFile, err = wcFileService.AssembleFile(file)
	return newFile, err
}

func (wcFileService *WcFileService) DeleteFile(file weChat.WcFile) (err error) {
	fileResponse, err := wcFileService.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileResponse.Key); err != nil {
		return errors.New("文件删除失败")
	}

	err = global.GVA_DB.Where("id = ?", fileResponse.ID).Unscoped().Delete(&file).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (wcFileService *WcFileService) EditFileName(file weChat.WcFile) (err error) {
	var fileFromDb weChat.WcFile
	return global.GVA_DB.Where("id=?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

func (wcFileService *WcFileService) GetFileRecordInfoList(info request.PageInfo, fileType, staffId int) (list []weChat2.WcFileResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.GVA_DB.Model(&weChat.WcFile{})
	var fileLists []weChat.WcFile
	db = db.Where("`type`=? and staff_id=?", fileType, staffId)
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	if err != nil {
		return
	}
	newFileLists, err := wcFileService.AssembleFileList(fileLists)

	fmt.Println("newFileLists", newFileLists)

	return newFileLists, total, err
}

// GetFileRecordInfoListByStaffId 通过员工ID获取证件材料
func (wcFileService *WcFileService) GetFileRecordInfoListByStaffId(staffId int) (list []weChat2.WcFileResponse, err error) {
	db := global.GVA_DB.Model(&weChat.WcFile{})
	var fileLists []weChat.WcFile
	db = db.Where("staff_id=?", staffId)
	err = db.Find(&fileLists).Error
	if err != nil {
		return
	}
	list, err = wcFileService.AssembleFileList(fileLists)
	return
}

func (wcFileService *WcFileService) UploadFile(header *multipart.FileHeader, fileType, staffId int) (newFile weChat2.WcFileResponse, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return newFile, uploadErr
	}
	s := strings.Split(header.Filename, ".")
	f := weChat.WcFile{
		Url:     filePath,
		Name:    header.Filename,
		Tag:     s[len(s)-1],
		Key:     key,
		StaffId: &staffId,
		Type:    &fileType,
	}

	newF, err := wcFileService.Upload(f)
	newFile, err = wcFileService.AssembleFile(newF)
	return newFile, err
}

func (wcFileService *WcFileService) AssembleFileList(files []weChat.WcFile) (newFiles []weChat2.WcFileResponse, err error) {
	configInfo := config.GetConfigInfo()
	var newFile weChat2.WcFileResponse
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

func (wcFileService *WcFileService) AssembleFile(file weChat.WcFile) (newFile weChat2.WcFileResponse, err error) {
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
