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

func InitReqPoolVo() *ReqPoolVo {
	vo := NewPageVo()
	return &ReqPoolVo{PageVo: *vo}
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

func InitReqNFTVo() *ReqNFTVo {
	vo := NewPageVo()
	return &ReqNFTVo{PageVo: *vo}
}

type ReqWNFTVo struct {
	PageVo  `json:"pageVo"`
	Address string `json:"address"`
	Account string `json:"account"`
}

func InitReqWNFTVo() *ReqWNFTVo {
	vo := NewPageVo()
	return &ReqWNFTVo{PageVo: *vo}
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
	Id               int64  ` json:"id"`
	TokenId          string `json:"tokenId"`
	TokenAddress     string `json:"tokenAddress"`
	TokenName        string `json:"tokenName"`
	TokenSymbol      string `json:"tokenSymbol"`
	TokenUri         string `json:"tokenUri"`
	Borrower         string `json:"borrower"`  // 借款人
	Mortgagor        string `json:"mortgagor"` // 抵押人
	Status           int64  `json:"status"`    // token状态初始值为0已被赎回，存入pool为1，借出为-1
	DelegatorAddress string `json:"delegatorAddress"`
}

type DepositListVo struct {
	Id                  int64     `json:"id"`
	CreatedTime         time.Time `json:"createdTime"`
	UpdatedTime         time.Time `json:"updatedTime"`
	Address             string    `json:"address"`
	Name                string    `json:"name"`
	Url                 string    `json:"url"`
	Owner               string    `json:"owner"`
	RewardsTokenName    string    `json:"rewardsTokenName"`
	RewardsTokenAddress string    `json:"rewardsTokenAddress"`
	//Apr                 string    `json:"apr"`
	//BorrowAPR           string    `json:"borrowAPR"`
	Type string `json:"type"`
	//WrapperAddress      string    `json:"wrapperAddress"`
	//WNFTtAddress        string    `json:"wnftAddress"`
	//DelegatorAddress    string    `json:"delegatorAddress"`
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
	EffectiveTime       string    `json:"effective_time"` // 生效时间
	Rate                RateModel `json:"rate"`           // 基础利率值
	NewRate             RateModel `json:"new_rate"`       // 新基础利率值
}

type RateModel struct {
	MultiplierPer     string `json:"multiplierPer"`
	BaseRatePer       string `json:"BaseRatePer"`
	JumpMultiplierPer string `json:"JumpMultiplierPer"`
	Kink              string `json:"kink"`
}
