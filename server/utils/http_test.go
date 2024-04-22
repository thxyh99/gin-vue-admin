package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSendGetRequest(t *testing.T) {
	url := "https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"
	body, err := SendGetRequest(url)
	if err != nil {
		return
	}
	fmt.Println(body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println(data)
	for key, value := range data {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func TestSendPostRequest(t *testing.T) {
	url := "https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"
	body, err := SendGetRequest(url)
	if err != nil {
		return
	}
	fmt.Println(body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println(data)
	for key, value := range data {
		fmt.Printf("%s: %v\n", key, value)
	}
}
