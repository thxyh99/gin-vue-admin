package weChat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"time"
)

type WcRankTypeService struct {
}

// GetRankTypeList 获取员工职级分类
func GetRankTypeList() (list []weChat2.WcRankTypeResponse, err error) {
	cacheKey := fmt.Sprintf("GetRankTypeList")
	cacheValue, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Result()
	if err != nil {
		fmt.Println("GetRankTypeList RedisStoreGetError:", err)
	}
	if cacheValue != "" {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("cacheValue:", cacheValue)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		err = json.Unmarshal([]byte(cacheValue), &list)
		return list, err
	} else {
		var rt weChat.WcRankType
		dRows, err := global.GVA_DB.Table(rt.TableName()).Select("id,name").Rows()
		if err != nil {
			fmt.Println("GetRankTypeList DbQueryErr:", err)
			return list, err
		} else {
			for dRows.Next() {
				var id int
				var name string
				err = dRows.Scan(&id, &name)
				if err != nil {
					fmt.Println("GetRankTypeList DbScanErr:", err)
					return list, err
				} else {
					list = append(list, weChat2.WcRankTypeResponse{
						ID:   id,
						Name: name,
					})
				}
			}

			jsonValue, err := json.Marshal(list)
			if err != nil {
				fmt.Println("GetRankTypeList JsonMarshalError:", err)
				return list, err
			}
			err = global.GVA_REDIS.Set(context.Background(), cacheKey, jsonValue, time.Hour*24).Err()
			if err != nil {
				fmt.Println("GetRankTypeList RedisStoreSetError:", err)
				return list, err
			}

		}
		return list, err
	}
}
