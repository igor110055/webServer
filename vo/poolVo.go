package vo

import "time"

type pageReqVo struct {
	PageVo  `json:"pageVo"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Id      string `json:"id"`
}

type SuppliesVo struct {
	PageNum     int64     `json:"pageNum"`
	PageSize    int64     `json:"pageSize"`
	TotalSize   int64     `json:"totalSize"`
	Id          int64     `json:"id"`
	PoolName    string    `json:"poolName"`
	PoolLogo    string    `json:"poolLogo"`
	PoolAddress string    `json:"poolAddress"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
	Token       TokenVo   `json:"token"`
}

type TokenVo struct {
	Id           int64
	TokenId      string
	TokenAddress string
	status       int64
	Borrower     string
	Mortgagor    string
}

type PoolDetailVo struct {
}
