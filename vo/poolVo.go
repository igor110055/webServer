package vo

import "time"

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
