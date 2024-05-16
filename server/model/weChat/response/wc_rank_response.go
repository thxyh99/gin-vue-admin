package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/weChat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// WcRankResponse  员工职级结构体
type WcRankResponse struct {
	weChat.WcRank
	TypeText string `json:"typeText"` //职级类型
}

// WcRankItem 员工等级名称
type WcRankItem struct {
	ID   int    `json:"ID"`   // ID
	Name string `json:"name"` //等级名称
}

func (WcRankResponse) AssembleRankList(ranks []weChat.WcRank) (newRanks []WcRankResponse, err error) {
	var newRank WcRankResponse
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		return
	}
	rankTypeMap := make(map[int]string)
	for _, rankTypeItem := range rankTypeList {
		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
	}
	for _, rank := range ranks {
		newRank.WcRank = rank
		typeText, _ := utils.GetValueByKey(rankTypeMap, *rank.Type)
		newRank.TypeText = typeText
		newRanks = append(newRanks, newRank)
	}
	return
}

func (WcRankResponse) AssembleRank(rank weChat.WcRank) (newRank WcRankResponse, err error) {
	newRank.WcRank = rank
	rankTypeList, err := GetRankTypeList()
	if err != nil {
		return
	}
	rankTypeMap := make(map[int]string)
	for _, rankTypeItem := range rankTypeList {
		rankTypeMap[rankTypeItem.ID] = rankTypeItem.Name
	}
	typeText, _ := utils.GetValueByKey(rankTypeMap, *rank.Type)
	newRank.TypeText = typeText
	return
}
