package db

import "time"

type Pool struct {
	Id               int64     `gorm:"id" json:"id"`
	CreatedTime      time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime      time.Time `gorm:"updated_time" json:"updated_time"`
	Deleted          int64     `gorm:"deleted" json:"deleted"`
	Address          string    `gorm:"address" json:"address"` // 合约地址
	Owner            string    `gorm:"owner" json:"owner"`     // 创建账户
	Url              string    `gorm:"url" json:"url"`
	Name             string    `gorm:"name" json:"name"`
	EffectiveTime    time.Time `gorm:"effective_time" json:"effective_time"`         // 生效时间
	Rate             string    `gorm:"rate" json:"rate"`                             // 基础利率值
	NewRate          string    `gorm:"new_rate" json:"new_rate"`                     // 新基础利率值
	LiquidateLine    float64   `gorm:"liquidate_line" json:"liquidate_line"`         // 清算线
	RateModelAddress string    `gorm:"rate_model_address" json:"rate_model_address"` // 利率模型地址
	ParaAddress      string    `gorm:"para_address" json:"para_address"`             // 参数合约地址
}

func (Pool) TableName() string {
	return "pool"
}

type PoolToken struct {
	Id             int64     `gorm:"id" json:"id"`
	CreatedTime    time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime    time.Time `gorm:"updated_time" json:"updated_time"`
	Deleted        int64     `gorm:"deleted" json:"deleted"`
	PoolAddress    string    `gorm:"pool_address" json:"pool_address"`
	TokenName      string    `gorm:"token_name" json:"token_name"`
	TokenSymbol    string    `gorm:"token_symbol" json:"token_symbol"`
	TokenAddress   string    `gorm:"token_address" json:"token_address"`
	Status         string    `gorm:"status" json:"status"` // 0 need/1 rewards
	WrapperAddress string    `gorm:"wrapper_address" json:"wrapper_address"`
	WnftAddress    string    `gorm:"wnft_address" json:"wnft_address"`
	Type           string    `gorm:"type" json:"type"`
	Principal      string    `gorm:"principal" json:"principal"` // 锚定价格
}

func (PoolToken) TableName() string {
	return "pool_token"
}

type Token struct {
	Id               int64     `gorm:"id" json:"id"`
	CreatedTime      time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime      time.Time `gorm:"updated_time" json:"updated_time"`
	Deleted          int64     `gorm:"deleted" json:"deleted"` // 如果赎回就变为-1
	PoolAddress      string    `gorm:"pool_address" json:"pool_address"`
	TokenId          string    `gorm:"token_id" json:"token_id"`
	TokenAddress     string    `gorm:"token_address" json:"token_address"`
	TokenUri         string    `gorm:"token_uri" json:"token_uri"`
	Borrower         string    `gorm:"borrower" json:"borrower"`   // 借款人
	Mortgagor        string    `gorm:"mortgagor" json:"mortgagor"` // 抵押人
	Status           int64     `gorm:"status" json:"status"`       // token状态初始值为0已被赎回，存入pool为1，借出为-1
	DelegatorAddress string    `gorm:"delegator_address" json:"delegator_address"`
}

func (Token) TableName() string {
	return "token"
}

//type Picture struct {
//	Id          int64     `gorm:"id" json:"id"`
//	CreatedTime time.Time `gorm:"created_time" json:"createdTime"`
//	UpdatedTime time.Time `gorm:"updated_time" json:"updatedTime"`
//	Type        string    `gorm:"type" json:"type"`
//	Url         string    `gorm:"url" json:"url"`
//	Deleted     int64     `gorm:"deleted" json:"deleted"`
//}
//
//func (Picture) TableName() string {
//	return "picture"
//}

type PoolListDto struct {
	Id           int64     `json:"id"`
	CreatedTime  time.Time `json:"createdTime"`
	UpdatedTime  time.Time `json:"updatedTime"`
	Address      string    `json:"address"`
	Name         string    `json:"name"`
	Url          string    `json:"url"`
	Owner        string    `json:"owner"`
	TokenName    string    `json:"tokenName"`
	TokenAddress string    `json:"tokenAddress"`
	Rate         string    `json:"rate"`
	Type         string    `json:"type"`
}
