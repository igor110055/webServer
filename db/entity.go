package db

import "time"

type Pool struct {
	Id            int64  `gorm:"id" json:"id"`
	CreatedTime   string `gorm:"created_time" json:"created_time"`
	UpdatedTime   string `gorm:"updated_time" json:"updated_time"`
	Deleted       int64  `gorm:"deleted" json:"deleted"`
	Address       string `gorm:"address" json:"address"` // 合约地址
	Owner         string `gorm:"owner" json:"owner"`     // 创建账户
	Url           string `gorm:"url" json:"url"`
	Name          string `gorm:"name" json:"name"`
	EffectiveTime string `gorm:"effective_time" json:"effective_time"` // 生效时间
	Rate          string `gorm:"rate" json:"rate"`                     // 基础利率值
	NewRate       string `gorm:"new_rate" json:"new_rate"`             // 新基础利率值
}

func (Pool) TableName() string {
	return "pool"
}

type PoolToken struct {
	Id             int64  `gorm:"id" json:"id"`
	CreatedTime    string `gorm:"created_time" json:"created_time"`
	UpdatedTime    string `gorm:"updated_time" json:"updated_time"`
	Deleted        int64  `gorm:"deleted" json:"deleted"` // 如果赎回就变为-1
	PoolId         string `gorm:"pool_id" json:"pool_id"`
	TokenUri       string `gorm:"token_uri" json:"token_uri"`
	TokenName      string `gorm:"token_name" json:"token_name"`
	TokenAddress   string `gorm:"token_address" json:"token_address"`
	WrapperAddress string `gorm:"wrapper_address" json:"wrapper_address"`
	WNFTAddress    string `gorm:"wnft_address" json:"wnft_address"`
	Type           string `gorm:"type" json:"type"`
}

func (PoolToken) TableName() string {
	return "pool_token"
}

type Token struct {
	Id               int64  `gorm:"id" json:"id"`
	CreatedTime      string `gorm:"created_time" json:"created_time"`
	UpdatedTime      string `gorm:"updated_time" json:"updated_time"`
	Deleted          int64  `gorm:"deleted" json:"deleted"` // 如果赎回就变为-1
	PoolAddress      string `gorm:"pool_address" json:"pool_address"`
	TokenId          string `gorm:"token_id" json:"token_id"`
	TokenAddress     string `gorm:"token_address" json:"token_address"`
	Borrower         string `gorm:"borrower" json:"borrower"`   // 借款人
	Mortgagor        string `gorm:"mortgagor" json:"mortgagor"` // 抵押人
	Status           int64  `gorm:"status" json:"status"`       // token状态初始值为0已被赎回，存入pool为1，借出为-1
	DelegatorAddress string `gorm:"delegator_address" json:"delegator_address"`
}

func (Token) TableName() string {
	return "token"
}

type Picture struct {
	Id          int64     `gorm:"id" json:"id"`
	CreatedTime time.Time `gorm:"created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"updated_time" json:"updatedTime"`
	Type        string    `gorm:"type" json:"type"`
	Url         string    `gorm:"url" json:"url"`
	Deleted     int64     `gorm:"deleted" json:"deleted"`
}

func (Picture) TableName() string {
	return "picture"
}

type Wnft struct {
	Id           int64     `gorm:"id" json:"id"`
	CreatedTime  time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime  time.Time `gorm:"updated_time" json:"updated_time"`
	Deleted      int64     `gorm:"deleted" json:"deleted"` // 如果赎回就变为-1
	PoolAddress  string    `gorm:"pool_address" json:"pool_address"`
	TokenId      string    `gorm:"token_id" json:"token_id"`
	TokenAddress string    `gorm:"token_address" json:"token_address"`
	TokenName    string    `gorm:"token_name" json:"token_name"`
	TokenUri     string    `gorm:"token_uri" json:"token_uri"`
	Owner        string    `gorm:"owner" json:"owner"`
}

func (Wnft) TableName() string {
	return "wnft"
}
