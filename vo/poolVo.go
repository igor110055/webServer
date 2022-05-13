package vo

import "time"

type ReqPoolVo struct {
	PageVo           `json:"pageVo"`
	Id               int64  `json:"id"`
	TokenId          string `json:"tokenId"`
	DelegatorAddress string `json:"delegatorAddress"`
	Borrower         string `json:"borrower"`  // 借款人
	Mortgagor        string `json:"mortgagor"` // 抵押人
}

type ReqNFTVo struct {
	PageVo           `json:"pageVo"`
	TokenId          string `json:"tokenId"`
	DelegatorAddress string `json:"delegatorAddress"`
	Status           string `json:"status"`
	PoolAddress      string `json:"poolAddress"`
	Borrower         string `json:"borrower"`  // 借款人
	Mortgagor        string `json:"mortgagor"` // 抵押人
	Address          string `json:"address"`
	NFTAddress       string `json:"nftAddress"`
}
type ReqWNFTVo struct {
	PageVo  `json:"pageVo"`
	Address string `json:"address"`
	Account string `json:"account"`
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
	Id int64 ` json:"id"`
	//CreatedTime      string `gorm:"created_time" json:"created_time"`
	//UpdatedTime      string `gorm:"updated_time" json:"updated_time"`
	//PoolAddress      string `gorm:"pool_address" json:"pool_address"`
	TokenId          string `json:"tokenId"`
	TokenAddress     string `json:"tokenAddress"`
	TokenName        string `json:"tokenName"`
	TokenSymbol      string `json:"tokenSymbol"`
	TokenUri         string `json:"tokenUri"`
	Borrower         string `json:"borrower"`  // 借款人
	Mortgagor        string `json:"mortgagor"` // 抵押人
	Status           int64  `json:"status"`    // token状态初始值为0已被赎回，存入pool为1，借出为-1
	DelegatorAddress string `json:"delegatorAddress"`
	From             string `json:"from"`
	To               string `json:"to"`
}

type PoolListVo struct {
	Id                  int64     `json:"id"`
	CreatedTime         time.Time `json:"createdTime"`
	UpdatedTime         time.Time `json:"updatedTime"`
	Address             string    `json:"address"`
	Name                string    `json:"name"`
	Url                 string    `json:"url"`
	Owner               string    `json:"owner"`
	TokenAddress        string    `json:"tokenAddress"`
	TokenName           string    `json:"tokenName"`
	RewardsTokenName    string    `json:"rewardsTokenName"`
	RewardsTokenAddress string    `json:"rewardsTokenAddress"`
	Apr                 string    `json:"apr"`
	BorrowAPR           string    `json:"borrowAPR"`
	Type                string    `json:"type"`
	WrapperAddress      string    `json:"wrapperAddress"`
	WNFTtAddress        string    `json:"wnftAddress"`
	DelegatorAddress    string    `json:"delegatorAddress"`
}

type PoolDetailVo struct {
	Id                  string    `json:"id"`
	CreatedTime         time.Time `json:"createdTime"`
	UpdatedTime         time.Time `json:"updatedTime"`
	Address             string    `json:"address"`
	Name                string    `json:"name"`
	Url                 string    `json:"url"`
	Owner               string    `json:"owner"`
	TokenAddress        string    `json:"tokenAddress"`
	TokenName           string    `json:"tokenName"`
	RewardsTokenName    string    `json:"rewardsTokenName"`
	RewardsTokenAddress string    `json:"rewardsTokenAddress"`
	BaseRate            string    `json:"baseRate"`
	InterestFactor      string    `json:"interestFactor"`
	KinkRate            string    `json:"kinkRate"`
	JumpMultiplier      string    `json:"jumpMultiplier"`
	EffectiveTime       string    `json:"effectiveTime"`
	NewBaseRate         string    `json:"newBaseRate"`
	NewInterestFactor   string    `json:"newInterestFactor"`
	NewKinkRate         string    `json:"newKinkRate"`
	NewJumpMultiplier   string    `json:"newJumpMultiplier"`
}
