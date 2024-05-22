package weChat

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
	"strings"
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

	newFile, err = weChat2.WcFileResponse{}.AssembleFile(file)
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
	newFileLists, err := weChat2.WcFileResponse{}.AssembleFileList(fileLists)

	fmt.Println("newFileLists", newFileLists)

	return newFileLists, total, err
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
	newFile, err = weChat2.WcFileResponse{}.AssembleFile(newF)
	return newFile, err
}
