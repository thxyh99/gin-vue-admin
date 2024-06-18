package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/employee"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type LandrayOa struct {
}

type OAStaffLeaveApplication struct {
	DocSubject  string     `json:"docSubject"`
	DocCreator  string     `json:"docCreator"`
	DocStatus   string     `json:"docStatus"`
	StaffName   string     `json:"StaffName"`
	LeaveDate   *time.Time `json:"LeaveDate"`
	LeaveType   int        `json:"LeaveType"`
	LeaveResult string     `json:"LeaveResult"`
}

// LoadConfig 加载配置文件
func LoadConfig(path string) (*ini.File, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		global.GVA_LOG.Error("LandrayOa LoadConfig 加载配置异常", zap.Error(err))
		return nil, err
	}
	return cfg, nil
}

func SendPostOaRequest(url string, data url.Values) ([]byte, error) {
	// 实现发送数据的逻辑
	client := &http.Client{}
	resp, err := client.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// CreateOAProcess 创建OA流程
func CreateOAProcess(data url.Values) ([]byte, error) {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return nil, err
	}
	// 读取OA接口配置
	oaWeb := cfg.Section("oa-web")
	oaUrl := oaWeb.Key("base-url").MustString("")
	addUrl := oaWeb.Key("addOaUrl").MustString("")

	return SendPostOaRequest(oaUrl+addUrl, data)
}

// CreateOAEmploymentApplication 创建OA入职申请
func (landrayOa *LandrayOa) CreateOAEmploymentApplication(wcStaffEmploymentApplication employee.WcStaffEmploymentApplication) (string, error) {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return "", err
	}

	eaConfig := cfg.Section("employment_application")
	// 配置OA提交

	eaData := url.Values{}
	eaData.Set(eaConfig.Key("StaffId").MustString(""), wcStaffEmploymentApplication.StaffName)
	eaData.Set(eaConfig.Key("EmploymentDate").MustString(""), wcStaffEmploymentApplication.EmploymentDate)
	eaData.Set(eaConfig.Key("EmploymentDepartmentID").MustString(""), wcStaffEmploymentApplication.EmploymentDepartmentId)
	eaData.Set(eaConfig.Key("EmploymentDepartmentName").MustString(""), wcStaffEmploymentApplication.EmploymentDepartmentName)
	eaData.Set(eaConfig.Key("JobPosition").MustString(""), wcStaffEmploymentApplication.JobPosition)
	eaData.Set(eaConfig.Key("Gender").MustString(""), wcStaffEmploymentApplication.Gender)
	eaData.Set(eaConfig.Key("Birthday").MustString(""), wcStaffEmploymentApplication.Birthday)
	eaData.Set(eaConfig.Key("NativePlace").MustString(""), wcStaffEmploymentApplication.NativePlace)
	eaData.Set(eaConfig.Key("Nationality").MustString(""), wcStaffEmploymentApplication.Nationality)
	eaData.Set(eaConfig.Key("Height").MustString(""), wcStaffEmploymentApplication.Height)
	eaData.Set(eaConfig.Key("Weight").MustString(""), wcStaffEmploymentApplication.Weight)
	eaData.Set(eaConfig.Key("Marriage").MustString(""), wcStaffEmploymentApplication.Marriage)
	eaData.Set(eaConfig.Key("PoliticalOutlook").MustString(""), wcStaffEmploymentApplication.PoliticalOutlook)
	eaData.Set(eaConfig.Key("Education").MustString(""), wcStaffEmploymentApplication.Education)
	eaData.Set(eaConfig.Key("Major").MustString(""), wcStaffEmploymentApplication.Major)
	eaData.Set(eaConfig.Key("School").MustString(""), wcStaffEmploymentApplication.School)
	eaData.Set(eaConfig.Key("GraduationDate").MustString(""), wcStaffEmploymentApplication.GraduationDate)
	eaData.Set(eaConfig.Key("Certificate").MustString(""), wcStaffEmploymentApplication.Certificate)
	eaData.Set(eaConfig.Key("IdNumber").MustString(""), wcStaffEmploymentApplication.IdNumber)
	eaData.Set(eaConfig.Key("IdAddress").MustString(""), wcStaffEmploymentApplication.IdAddress)
	eaData.Set(eaConfig.Key("BankAccount").MustString(""), wcStaffEmploymentApplication.BankAccount)
	eaData.Set(eaConfig.Key("BankName").MustString(""), wcStaffEmploymentApplication.BankName)
	eaData.Set(eaConfig.Key("ContactWechat").MustString(""), wcStaffEmploymentApplication.ContactWechat)
	eaData.Set(eaConfig.Key("Mobile").MustString(""), wcStaffEmploymentApplication.Mobile)
	eaData.Set(eaConfig.Key("HomeAddress").MustString(""), wcStaffEmploymentApplication.HomeAddress)
	//eaData.Set(eaConfig.Key("LicenseAttachment"),wcStaffEmploymentApplication.LicenseAttachment
	eaData.Set(eaConfig.Key("RelationShip").MustString(""), wcStaffEmploymentApplication.RelationShip)
	eaData.Set(eaConfig.Key("RelationName").MustString(""), wcStaffEmploymentApplication.RelationName)
	eaData.Set(eaConfig.Key("RelationMobile").MustString(""), wcStaffEmploymentApplication.RelationMobile)
	eaData.Set(eaConfig.Key("RelationAddress").MustString(""), wcStaffEmploymentApplication.RelationAddress)
	eaData.Set(eaConfig.Key("IsCeopass").MustString(""), "0")
	eaData.Set(eaConfig.Key("IsBodychecknormal").MustString(""), "2")
	eaData.Set(eaConfig.Key("JobLevel").MustString(""), wcStaffEmploymentApplication.JobLevel)
	eaData.Set(eaConfig.Key("TryPeriod").MustString(""), wcStaffEmploymentApplication.TryPeriod)
	eaData.Set(eaConfig.Key("UrgencyNotification").MustString(""), wcStaffEmploymentApplication.UrgencyNotification)
	//eaData.Set(eaConfig.Key("OnboardingOpinion"),wcStaffEmploymentApplication.OnboardingOpinion

	//
	tmplateId := eaConfig.Key("fdTemplateId").MustString("")
	oaData := url.Values{}
	oaData.Set("fdTemplateId", tmplateId)
	oaData.Set("docSubject", wcStaffEmploymentApplication.Title)
	oaData.Set("docStatus", string(wcStaffEmploymentApplication.OaStatus))
	oaData.Set("docCreator", `{"LoginName":"liuyongbo"}`)
	jsonData, err := json.Marshal(eaData)
	if err != nil {
		return "", err
	}
	oaData.Set("formValues", string(jsonData))

	dsData, err := CreateOAProcess(oaData)
	if err != nil {
		return "", err
	}
	return string(dsData), nil
}

// 创建OA转正申请
func (landrayOa *LandrayOa) CreateOAPassApplication(wcStaffPassApplication employee.WcStaffPassApplication) (string, error) {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return "", err
	}

	eaConfig := cfg.Section("pass_application")
	// 配置OA提交
	eaData := url.Values{}
	eaData.Set(eaConfig.Key("StaffId").MustString(""), "")
	eaData.Set(eaConfig.Key("EmploymentDate").MustString(""), wcStaffPassApplication.EmploymentDate)
	//eaData.Set(eaConfig.Key("EmploymentDepartmentID").MustString(""), wcStaffPassApplication.JobDepartmentId.String())
	eaData.Set(eaConfig.Key("EmploymentDepartmentName").MustString(""), wcStaffPassApplication.JobDepartment)
	eaData.Set(eaConfig.Key("JobPosition").MustString(""), wcStaffPassApplication.SourcePosition)
	eaData.Set(eaConfig.Key("Education").MustString(""), wcStaffPassApplication.Education)
	eaData.Set(eaConfig.Key("JobLevel").MustString(""), wcStaffPassApplication.SourceLevel)
	eaData.Set(eaConfig.Key("TryPeriod").MustString(""), wcStaffPassApplication.TryPeriod)

	tmplateId := eaConfig.Key("fdTemplateId").MustString("")
	oaData := url.Values{}
	oaData.Set("fdTemplateId", tmplateId)
	oaData.Set("docSubject", wcStaffPassApplication.Title)
	oaData.Set("docStatus", string(wcStaffPassApplication.OaStatus))
	oaData.Set("docCreator", `{"LoginName":"liuyongbo"}`)

	jsonData, err := json.Marshal(eaData)
	if err != nil {
		return "", err
	}
	oaData.Set("formValues", string(jsonData))

	dsData, err := CreateOAProcess(oaData)
	if err != nil {
		return "", err
	}
	return string(dsData), nil
}

// 创建OA调级申请
func (landrayOa *LandrayOa) CreateOAAdjustlevelApplication(wcStaffAdjustlevelApplication employee.WcStaffAdjustlevelApplication) (string, error) {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return "", err
	}

	eaConfig := cfg.Section("adjust_application")
	// 配置OA提交
	eaData := url.Values{}
	eaData.Set(eaConfig.Key("StaffId").MustString(""), "")
	//eaData.Set(eaConfig.Key("EmploymentDate").MustString(""), wcStaffAdjustlevelApplication.EmploymentDate)
	//eaData.Set(eaConfig.Key("EmploymentDepartmentID").MustString(""), wcStaffPassApplication.JobDepartmentId.String())
	eaData.Set(eaConfig.Key("EmploymentDepartmentName").MustString(""), wcStaffAdjustlevelApplication.SourceDepartment)
	eaData.Set(eaConfig.Key("memo").MustString(""), wcStaffAdjustlevelApplication.Memo)

	tmplateId := eaConfig.Key("fdTemplateId").MustString("")
	oaData := url.Values{}
	oaData.Set("fdTemplateId", tmplateId)
	oaData.Set("docSubject", wcStaffAdjustlevelApplication.Title)
	oaData.Set("docStatus", string(wcStaffAdjustlevelApplication.OaStatus))
	oaData.Set("docCreator", `{"LoginName":"liuyongbo"}`)

	jsonData, err := json.Marshal(eaData)
	if err != nil {
		return "", err
	}
	oaData.Set("formValues", string(jsonData))

	dsData, err := CreateOAProcess(oaData)
	if err != nil {
		return "", err
	}
	return string(dsData), nil
}

// 创建OA离职申请
func (landrayOa *LandrayOa) CreateOALeaveApplication(oaStaffLeaveApplication OAStaffLeaveApplication) (string, error) {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return "", err
	}

	eaConfig := cfg.Section("leave_application")
	// 配置OA提交
	eaData := url.Values{}
	eaData.Set(eaConfig.Key("StaffName").MustString(""), oaStaffLeaveApplication.StaffName)
	eaData.Set(eaConfig.Key("LeaveDate").MustString(""), oaStaffLeaveApplication.LeaveDate.GoString())
	eaData.Set(eaConfig.Key("LeaveType").MustString(""), strconv.Itoa(oaStaffLeaveApplication.LeaveType))
	eaData.Set(eaConfig.Key("LeaveResult").MustString(""), oaStaffLeaveApplication.LeaveResult)

	tmplateId := eaConfig.Key("fdTemplateId").MustString("")
	oaData := url.Values{}
	oaData.Set("fdTemplateId", tmplateId)
	oaData.Set("docSubject", oaStaffLeaveApplication.DocSubject)
	oaData.Set("docStatus", oaStaffLeaveApplication.DocStatus)
	oaData.Set("docCreator", `{"LoginName":"`+oaStaffLeaveApplication.DocCreator+`"}`)

	jsonData, err := json.Marshal(eaData)
	if err != nil {
		return "", err
	}
	oaData.Set("formValues", string(jsonData))
	fmt.Printf(string(jsonData))
	dsData, err := CreateOAProcess(oaData)
	if err != nil {
		return "", err
	}
	return string(dsData), nil
}

// 创建OA离职申请
func (landrayOa *LandrayOa) CreateOATransferApplication(wcStaffTransferApplication employee.WcStaffTransferApplication) (string, error) {
	cfg, err := LoadConfig("./oa-config.ini")
	if err != nil {
		return "", err
	}

	eaConfig := cfg.Section("transfer_application")
	// 配置OA提交
	eaData := url.Values{}
	eaData.Set(eaConfig.Key("TransferName").MustString(""), "刘永波")
	eaData.Set(eaConfig.Key("EmploymentDate").MustString(""), wcStaffTransferApplication.EmploymentDate)
	eaData.Set(eaConfig.Key("TransferType").MustString(""), wcStaffTransferApplication.TransferType)

	var oaTransferData string = ""
	oaTransferData += SetNameValueJson(eaConfig.Key("TransferName").MustString(""), SetIdName("165859c24b5e99309ab2c7349e7bb6ac", "刘永波"))
	oaTransferData += "," + SetNameValue(eaConfig.Key("EmploymentDate").MustString(""), wcStaffTransferApplication.EmploymentDate)
	oaTransferData += "," + SetNameValue(eaConfig.Key("TransferType").MustString(""), wcStaffTransferApplication.TransferType)
	oaTransferData += "," + SetNameValue(eaConfig.Key("TransferResult").MustString(""), wcStaffTransferApplication.TransferResult)
	oaTransferData += "," + SetNameValueJson(eaConfig.Key("TransferStaff").MustString(""), SetIdName("165859c24b5e99309ab2c7349e7bb6ac", "刘永波"))
	oaTransferData += "," + SetNameValueJson(eaConfig.Key("SourceDepartment").MustString(""), SetIdName("1644f7a70f2503bb7bf923a4c479aff1", "容大集团/集团总部/信息技术部"))
	oaTransferData += "," + SetNameValue(eaConfig.Key("SourceJobPosition").MustString(""), wcStaffTransferApplication.SourcePosition)
	oaTransferData += "," + SetNameValueJson(eaConfig.Key("NewDepartment").MustString(""), SetIdName("1644f7a70a2554c24bcefe44be68c55e", "容大集团/集团总部/综合行政部"))
	oaTransferData += "," + SetNameValue(eaConfig.Key("NewJobPosition").MustString(""), wcStaffTransferApplication.NewPosition)
	oaTransferData += "," + SetNameValue(eaConfig.Key("SourceResult").MustString(""), wcStaffTransferApplication.SourceResult)
	oaTransferData += "," + SetNameValue(eaConfig.Key("ToResult").MustString(""), wcStaffTransferApplication.ToResult)
	oaTransferData += "," + SetNameValue(eaConfig.Key("ToDate").MustString(""), wcStaffTransferApplication.ToDate)
	oaTransferData += "," + SetNameValue(eaConfig.Key("InspectionPerion").MustString(""), wcStaffTransferApplication.InspectionPerion)

	/*
		tempData := url.Values{"Id": "165859c24b5e99309ab2c7349e7bb6ac", "name": "刘永波"}
		tsData, err := json.Marshal(tempData)
	*/
	//eaData.Set(eaConfig.Key("TransferStaff").MustString(""), string(tsData))
	eaData.Set(eaConfig.Key("TransferStaff").MustString(""), wcStaffTransferApplication.TransferStaff)

	eaData.Set(eaConfig.Key("SourceDepartment").MustString(""), wcStaffTransferApplication.SourceDepartment)
	eaData.Set(eaConfig.Key("SourceJobPosition").MustString(""), wcStaffTransferApplication.SourcePosition)
	eaData.Set(eaConfig.Key("NewDepartment").MustString(""), wcStaffTransferApplication.NewDepartment)
	eaData.Set(eaConfig.Key("NewJobPosition").MustString(""), wcStaffTransferApplication.NewPosition)
	eaData.Set(eaConfig.Key("TransferResult").MustString(""), wcStaffTransferApplication.TransferResult)
	eaData.Set(eaConfig.Key("SourceResult").MustString(""), wcStaffTransferApplication.SourceResult)
	eaData.Set(eaConfig.Key("ToResult").MustString(""), wcStaffTransferApplication.ToResult)
	eaData.Set(eaConfig.Key("ToDate").MustString(""), wcStaffTransferApplication.ToDate)
	eaData.Set(eaConfig.Key("InspectionPerion").MustString(""), "1")

	tmplateId := eaConfig.Key("fdTemplateId").MustString("")
	oaData := url.Values{}
	oaData.Set("fdTemplateId", tmplateId)
	oaData.Set("docSubject", wcStaffTransferApplication.Title)
	oaData.Set("docStatus", string(wcStaffTransferApplication.OaStatus))
	oaData.Set("docCreator", `{"LoginName":"liuyongbo"}`)
	oaData.Set("attachmentForms.fdKey", `"fd_36829ea354aa5e"`)
	oaData.Set("attachmentForms.fdFileName", `"favicon.ico"`)
	fileData, err := ReadFileToBase64("/favicon.ico")
	oaData.Set("attachmentForms.fdAttachment", fileData)
	/*
		jsonData, err := json.Marshal(eaData)
		if err != nil {
			return "", err
		}
	*/
	fmt.Println(oaData)
	oaData.Set("formValues", "{"+oaTransferData+"}")
	fmt.Println(oaTransferData)
	dsData, err := CreateOAProcess(oaData)
	if err != nil {
		return "", err
	}
	return string(dsData), nil
}

// 读取文件内容为base64
func ReadFileToBase64(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	return base64.StdEncoding.EncodeToString(buffer), nil
}

// 设置ID和名称Json串
func SetIdName(id string, name string) string {
	retData := fmt.Sprintf(`{"Id":"%s","name":"%s"}`, id, name)
	return retData
}

// 设置名称和值Json串
func SetNameValue(name string, value string) string {
	retData := fmt.Sprintf(`"%s":"%s"`, name, value)
	return retData
}

// 设置名称和值Json串
func SetNameValueJson(name string, value string) string {
	retData := fmt.Sprintf(`"%s":%s`, name, value)
	return retData
}
