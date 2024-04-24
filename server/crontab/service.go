package crontab

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/crontab/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type Department struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	NameEN           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentID         int      `json:"parentid"`
	Order            int      `json:"order"`
}

func (s *Service) SyncDepartment() {
	fmt.Println("Sync department success:", time.Now())
	accessToken, tokenErr := s.getToken()
	if tokenErr != nil {
		return
	}
	//获取部门列表,废弃了不让请求
	//url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s", accessToken)
	//获取子部门ID列表
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s", accessToken)
	body, err := utils.SendGetRequest(url)
	if err != nil {
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 获取 department_id 数组
	departmentID, ok := data["department_id"].([]interface{})
	if !ok {
		fmt.Println("Error getting department_id")
		return
	}

	// 遍历 department_id 数组
	for _, department := range departmentID {
		departmentMap, ok := department.(map[string]interface{})
		if !ok {
			fmt.Println("Error converting department to map")
			continue
		}
		id := departmentMap["id"].(float64)
		parentID := departmentMap["parentid"].(float64)
		order := departmentMap["order"].(float64)
		fmt.Printf("ID: %d, Parentid: %d, Order: %d\n", int(id), int(parentID), int(order))

		////获取单个部门详情
		// 从2022年8月15日10点开始，新增IP不能访问了
		//urlInfo := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/get?access_token=%s&id=%s", accessToken, id)
		//bodyInfo, err := utils.SendGetRequest(urlInfo)
		//if err != nil {
		//	return
		//}
		//var dataInfo map[string]interface{}
		//err = json.Unmarshal(bodyInfo, &dataInfo)
		//if err != nil {
		//	fmt.Println("Error decoding JSON:", err)
		//	return
		//}
		//
		//fmt.Println(dataInfo)
		//
		//// 获取 department 数组
		//departmentData, ok := dataInfo["department"].(map[string]interface{})
		//fmt.Println(departmentData)
		//if !ok {
		//	fmt.Println("Error getting department")
		//	return
		//}
		//// 解析 department 字段到 Department 结构体
		//var department Department
		//departmentJSON, err := json.Marshal(departmentData)
		//if err != nil {
		//	fmt.Println("Error decoding department data:", err)
		//	return
		//}
		//err = json.Unmarshal(departmentJSON, &department)
		//if err != nil {
		//	fmt.Println("Error decoding department JSON:", err)
		//	return
		//}
		//
		//// 输出解析结果
		//fmt.Printf("Department ID: %d\n", department.ID)
		//fmt.Printf("Department Name: %s\n", department.Name)
		//fmt.Printf("Department Name (EN): %s\n", department.NameEN)
		//fmt.Printf("Department Leaders: %v\n", department.DepartmentLeader)
		//fmt.Printf("Department Parent ID: %d\n", department.Parentid)
		//fmt.Printf("Department Order: %d\n", department.Order)
	}

}

func (s *Service) getToken() (string, error) {
	// 构建请求URL
	corpID := internal.CorpId
	corpSecret := internal.Secret
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpID, corpSecret)

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// 解析JSON响应
	var accessTokenResp AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResp)
	if err != nil {
		fmt.Println("Error Unmarshal JSON response:", err)
		return "", err
	}

	if accessTokenResp.ErrCode != 0 {
		fmt.Println("Error Get Token response:", accessTokenResp.ErrMsg)
		return "", errors.New(accessTokenResp.ErrMsg)
	}
	fmt.Println("Access Token:", accessTokenResp.AccessToken)
	fmt.Println("Expires In:", accessTokenResp.ExpiresIn)
	return accessTokenResp.AccessToken, nil
}
