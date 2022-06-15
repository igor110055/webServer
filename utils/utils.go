package utils

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"poolServer/db"
	"poolServer/vo"
)

type NftVo struct {
	Id        int64  `json:"ID"` // ID
	ProjectId int64  `json:"projectId"`
	Level     int64  `json:"level"`
	Status    string `json:"status"`
	Address   string `json:"address"`
	Number    int64  `json:"number"`
}

//func TimeStampToTime(tm int64) time.Time {
//	tm = tm / 1000
//	timeFormat := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
//	duration, _ := time.ParseInLocation("2006-01-02 15:04:05", timeFormat, time.Local)
//	return duration
//}

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

	borrowerAPR, apr := calculateAPR(dto.Rate, dto.Address)
	listVo.BorrowAPR = borrowerAPR
	listVo.Apr = apr
	return &listVo
}

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
		LiquidateLine: dto.LiquidateLine,
	}
	poolTokens := db.GetPoolTokenByPoolAddress(dto.Address)

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
			res.WNFTAddress = value.WnftAddress
			res.WrapperAddress = value.WrapperAddress
		}
	}

	//计算年化利率
	borrowerAPR, apr := calculateAPR(dto.Rate, dto.Address)
	res.BorrowAPR = borrowerAPR
	res.Apr = apr
	if dto.Rate != "" {
		json.Unmarshal([]byte(dto.Rate), &rateMode)
		res.Rate = &rateMode
	}

	if dto.NewRate != "" {
		json.Unmarshal([]byte(dto.NewRate), &newRateModel)
		res.NewRate = &newRateModel
	}

	return &res
}

func calculateAPR(rate, poolAddress string) (string, string) {
	//var borrowerAPR, apr string
	borrowerAPR := decimal.NewFromFloat(0)
	utilizationRate := decimal.NewFromFloat(0)
	if rate == "" {
		return "0", "0"
	}

	//calculateAPR
	borrowers, total := db.GetTotalBorrower(poolAddress)
	//防止分母为0
	if total != 0 {
		utilizationRate = decimal.NewFromFloat(float64(borrowers)).Div(decimal.NewFromFloat(float64(total)))
	}

	model := vo.RateModel{}
	json.Unmarshal([]byte(rate), &model)
	base := decimal.NewFromFloat(model.BaseRate)
	kink := decimal.NewFromFloat(model.Kink)
	multiplier := decimal.NewFromFloat(model.Multiplier)
	jumpMultiplier := decimal.NewFromFloat(model.JumpMultiplier)

	if utilizationRate.Cmp(kink) == 1 {
		//如果使用率>边界利率
		//(基础利率+(边界利率*利率因子)) + (使用率-边界利率)*利率因子*加成系数
		//borrowerAPR = (model.BaseRate + (model.Kink * model.Multiplier)) + (utilizationRate-model.Kink)*model.Multiplier*model.JumpMultiplier
		borrowerAPR = base.Add(kink.Mul(multiplier)).Add(utilizationRate.Sub(kink).Mul(multiplier).Mul(jumpMultiplier))
		fmt.Println(borrowerAPR.String())
	} else {
		//如果使用<=边界利率
		//基础利率+(使用率*利率因子)
		//borrowerAPR = model.BaseRate + (utilizationRate * model.Multiplier)
		borrowerAPR = base.Add(utilizationRate.Mul(multiplier))
		fmt.Println(borrowerAPR.String())
	}
	//apr = borrowerAPR * utilizationRate
	apr := borrowerAPR.Mul(utilizationRate).String()
	return borrowerAPR.String(), apr

}
