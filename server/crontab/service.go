package crontab

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/crontab/internal"
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

func (s *Service) SyncDepartment() {
	fmt.Println("Sync department success:", time.Now())
	accessToken, tokenErr := s.getToken()
	if tokenErr != nil {
		return
	}
	//获取部门列表
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s", accessToken)
	//获取子部门ID列表
	//url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))

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
