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

	borrowerAPR, apr := CalculateAPR(dto.Rate, dto.Address)
	listVo.BorrowAPR = borrowerAPR
	listVo.Apr = apr
	return &listVo
}

//计算年化利率
func TransferPoolDetailVo(dto *db.Pool) *vo.PoolDetailVo {
	var rateMode, newRateModel vo.RateModel

	res := vo.PoolDetailVo{
		Id:            dto.Id,
		CreatedTime:   dto.CreatedTime,
		UpdatedTime:   dto.UpdatedTime,
		Address:       dto.Address,
		Name:          dto.Name,
		Url:           dto.Url,
		Owner:         dto.Owner,
		EffectiveTime: dto.EffectiveTime,
	}
	poolTokens := db.GetPoolTokenByPoolId(dto.Id)

	if poolTokens == nil {
		return &res
	}

	//组装Vo
	for _, value := range *poolTokens {
		//0 need 1 rewards
		if value.Type == "erc20" {
			res.RewardsTokenName = value.TokenName
			res.RewardsTokenAddress = value.TokenAddress
		} else if value.Type == "erc721" {
			res.TokenName = value.TokenName
			res.TokenAddress = value.TokenAddress
		}
	}
	borrowerAPR, apr := CalculateAPR(dto.Rate, dto.Address)
	res.BorrowAPR = borrowerAPR
	res.Apr = apr

	json.Unmarshal([]byte(dto.Rate), &rateMode)
	json.Unmarshal([]byte(dto.NewRate), &newRateModel)
	res.Rate = rateMode
	res.NewRate = newRateModel
	return &res
}

func CalculateAPR(rate, poolAddress string) (float64, float64) {
	var borrowerAPR, apr float64
	if rate == "" {
		return 0, 0
	}

	//calculateAPR
	borrowers, total := db.GetTotalBorrower(poolAddress)
	if total == 0 {
		return 0, 0
	}
	utilizationRate := float64(borrowers) / float64(total)
	model := vo.RateModel{}
	json.Unmarshal([]byte(rate), &model)

	//如果使用率>边界利率
	//(基础利率+(边界利率*利率因子)) + (使用率-边界利率*利率因子*加成系数)
	if utilizationRate > model.Kink {
		borrowerAPR = (model.BaseRate + (model.Kink * model.Multiplier)) + (utilizationRate - model.Kink*model.Multiplier*model.JumpMultiplier)
	} else {
		//基础利率+(使用率*利率因子)
		borrowerAPR = model.BaseRate + (utilizationRate * model.Multiplier)
	}
	apr = borrowerAPR * utilizationRate
	return borrowerAPR, apr

}
