package weChat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	weChatReq "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/request"
	weChat2 "github.com/flipped-aurora/gin-vue-admin/server/model/weChat/response"
	"time"
)

type WcRankService struct {
}

// CreateWcRank 创建职级管理记录
func (wcRankService *WcRankService) CreateWcRank(wcRank *weChat.WcRank) (err error) {
	err = global.GVA_DB.Create(wcRank).Error
	return err
}

// DeleteWcRank 删除职级管理记录
func (wcRankService *WcRankService) DeleteWcRank(ID string) (err error) {
	err = global.GVA_DB.Delete(&weChat.WcRank{}, "id = ?", ID).Error
	return err
}

// DeleteWcRankByIds 批量删除职级管理记录
func (wcRankService *WcRankService) DeleteWcRankByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]weChat.WcRank{}, "id in ?", IDs).Error
	return err
}

// UpdateWcRank 更新职级管理记录
func (wcRankService *WcRankService) UpdateWcRank(wcRank weChat.WcRank) (err error) {
	err = global.GVA_DB.Save(&wcRank).Error
	return err
}

// GetWcRank 根据ID获取职级管理记录
func (wcRankService *WcRankService) GetWcRank(ID string) (wcRank weChat2.WcRankResponse, err error) {
	var rank weChat.WcRank
	err = global.GVA_DB.Where("id = ?", ID).First(&rank).Error
	if err != nil {
		return
	}

	wcRank, err = weChat2.WcRankResponse{}.AssembleRank(rank)
	return
}

// GetWcRankInfoList 分页获取职级管理记录
func (wcRankService *WcRankService) GetWcRankInfoList(info weChatReq.WcRankSearch) (list []weChat2.WcRankResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&weChat.WcRank{})
	var wcRanks []weChat.WcRank
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

	err = db.Find(&wcRanks).Error
	if err != nil {
		return
	}

	list, err = weChat2.WcRankResponse{}.AssembleRankList(wcRanks)
	return
}

// GetRankTypeList 分页获取职级类型记录
func (wcRankService *WcRankService) GetRankTypeList(_ weChatReq.WcRankSearch) (list []weChat2.WcRankTypeResponse, total int64, err error) {
	list, err = weChat2.GetRankTypeList()
	if err != nil {
		return
	}
	total = int64(len(list))
	return
}

// GetRankListByRankType 分页获取职级记录
func (wcRankService *WcRankService) GetRankListByRankType(rankType string) (list []weChat2.WcRankItem, total int64, err error) {
	cacheKey := fmt.Sprintf("GetRankListByRankType:%s", rankType)
	cacheValue, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Result()
	if err != nil {
		fmt.Println("GetRankListByRankType RedisStoreGetError:", err)
	}
	if cacheValue != "" {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("cacheValue:", cacheValue)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		err = json.Unmarshal([]byte(cacheValue), &list)
		return list, 0, err
	} else {
		var r weChat.WcRank
		rRows, err := global.GVA_DB.Table(r.TableName()).Select("id,name").Where("type=?", rankType).Rows()
		if err != nil {
			fmt.Println("GetRankListByRankType DbQueryErr:", err)
			return list, 0, err
		} else {
			for rRows.Next() {
				var id int
				var name string
				err = rRows.Scan(&id, &name)
				if err != nil {
					fmt.Println("GetRankListByRankType DbScanErr:", err)
					return list, 0, err
				} else {
					list = append(list, weChat2.WcRankItem{
						ID:   id,
						Name: name,
					})
				}
			}

			jsonValue, err := json.Marshal(list)
			if err != nil {
				fmt.Println("GetRankListByRankType JsonMarshalError:", err)
				return list, 0, err
			}
			err = global.GVA_REDIS.Set(context.Background(), cacheKey, jsonValue, time.Hour*24).Err()
			//err = global.GVA_REDIS.Set(context.Background(), cacheKey, jsonValue, time.Minute).Err()
			if err != nil {
				fmt.Println("GetRankListByRankType RedisStoreSetError:", err)
				return list, 0, err
			}
		}
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("list", list)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	total = int64(len(list))
	return list, total, nil
}
