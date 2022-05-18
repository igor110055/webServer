package utils

import (
	"encoding/json"
	"poolServer/db"
	"poolServer/vo"
	"time"
)

type NftVo struct {
	Id        int64  `json:"ID"` // ID
	ProjectId int64  `json:"projectId"`
	Level     int64  `json:"level"`
	Status    string `json:"status"`
	Address   string `json:"address"`
	Number    int64  `json:"number"`
}

func TimeStampToTime(tm int64) time.Time {
	tm = tm / 1000
	timeFormat := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
	duration, _ := time.ParseInLocation("2006-01-02 15:04:05", timeFormat, time.Local)
	return duration
}

//计算年化利率
func TransferPoolListVo(dto db.PoolListDto) *vo.PoolListVo {
	listVo := vo.PoolListVo{
		Id:           dto.Id,
		CreatedTime:  dto.CreatedTime,
		UpdatedTime:  dto.UpdatedTime,
		Address:      dto.Address,
		Name:         dto.Name,
		Url:          dto.Url,
		Owner:        dto.Owner,
		TokenName:    dto.TokenName,
		TokenAddress: dto.TokenAddress,
		Type:         dto.Type,
	}

	if dto.Rate == "" {
		listVo.BorrowAPR = 0
		listVo.Apr = 0
		return &listVo
	}

	//calculateAPR
	model := vo.RateModel{}
	borrowers, total := db.GetTotalBorrower(dto.Address)
	if total == 0 {
		listVo.BorrowAPR = 0
		listVo.Apr = 0
		return &listVo
	}
	utilizationRate := float64(borrowers) / float64(total)
	json.Unmarshal([]byte(dto.Rate), &model)

	//如果使用率>边界利率
	//(基础利率+(边界利率*利率因子)) + (使用率-边界利率*利率因子*加成系数)
	if utilizationRate > model.Kink {
		listVo.BorrowAPR = (model.BaseRate + (model.Kink * model.Multiplier)) + (utilizationRate - model.Kink*model.Multiplier*model.JumpMultiplier)
	} else {
		//基础利率+(使用率*利率因子)
		listVo.BorrowAPR = model.BaseRate + (utilizationRate * model.Multiplier)
	}
	listVo.Apr = listVo.BorrowAPR * utilizationRate
	return &listVo
}
