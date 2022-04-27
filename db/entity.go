package db

import "time"

type Assets struct {
	Id           int64     `gorm:"id" json:"id"`
	CreatedTime  time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime  time.Time `gorm:"updated_time" json:"updated_time"`
	Deleted      int64     `gorm:"deleted" json:"deleted"`
	Type         string    `gorm:"type" json:"type"`
	TokenAddress string    `gorm:"token_address" json:"token_address"`
	TokenName    string    `gorm:"token_name" json:"token_name"`
	PoolAddress  string    `gorm:"pool_address" json:"pool_address"`
}

func (Assets) TableName() string {
	return "assets"
}

type Pool struct {
	Id          int64     `gorm:"id" json:"id"`
	CreatedTime time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"`
	Address     string    `gorm:"address" json:"address"` // 合约地址
	Deleted     int64     `gorm:"deleted" json:"deleted"`
	Owner       string    `gorm:"owner" json:"owner"` // 创建账户
	Uri         string    `gorm:"uri" json:"uri"`
	PoolName    string    `gorm:"pool_name" json:"pool_name"`
}

func (Pool) TableName() string {
	return "pool"
}

type Token struct {
	Id           int64     `gorm:"id" json:"id"`
	CreatedTime  time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime  time.Time `gorm:"updated_time" json:"updated_time"`
	Deleted      int64     `gorm:"deleted" json:"deleted"`
	PoolAddress  string    `gorm:"pool_address" json:"pool_address"`
	TokenId      int64     `gorm:"token_id" json:"token_id"`
	TokenAddress string    `gorm:"token_address" json:"token_address"`
	Borrower     string    `gorm:"borrower" json:"borrower"`   // 借款人
	Mortgagor    string    `gorm:"mortgagor" json:"mortgagor"` // 抵押人
	Status       int64     `gorm:"status" json:"status"`       // token状态初始值为0，存入pool为1，借出为-1
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
