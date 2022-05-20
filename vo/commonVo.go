package vo

import "poolServer/config"

type ReqMoralisVo struct {
	Address      string `json:"address"`
	TokenAddress string `json:"tokenAddress"`
	Cursor       string `json:"cursor"`
	PageSize     string `json:"pageSize"`
}

func InitReqMoralisVo() *ReqMoralisVo {
	return &ReqMoralisVo{Cursor: "", PageSize: "10"}
}

type MoralisResVo struct {
	Cursor string  `json:"cursor"`
	Tokens []Token `json:"tokens"`
}
type Token struct {
	TokenId      string `json:"tokenId"`
	TokenAddress string `json:"tokenAddress"`
	TokenName    string `json:"tokenName"`
	TokenSymbol  string `json:"tokenSymbol"`
	TokenUri     string `json:"tokenUri"`
}
type ReqVo struct {
	PageVo           `json:"pageVo"`
	Id               string `json:"id"`
	Address          string `json:"address"`
	NFTAddress       string `json:"nftAddress"`
	TokenId          string `json:"tokenId"`
	DelegatorAddress string `json:"delegatorAddress"`
	Type             string `json:"type"`
}

func InitReqVo() *ReqVo {
	vo := NewPageVo()
	return &ReqVo{PageVo: *vo}
}

type PageVo struct {
	PageNum   int64 `json:"pageNum"`
	PageSize  int64 `json:"pageSize"`
	TotalSize int64 `json:"totalSize"`
}

func NewPageVo() *PageVo {
	return &PageVo{PageNum: 1, PageSize: 10}
}

type ResponsePageVo struct {
	PageNum   int64       `json:"pageNum"`
	PageSize  int64       `json:"pageSize"`
	TotalSize int64       `json:"totalSize"`
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func NewResponsePageVo(pageNum, pageSize, totalSize, code int64, data interface{}) *ResponsePageVo {
	return &ResponsePageVo{PageNum: pageNum, PageSize: pageSize, TotalSize: totalSize, Code: code, Message: config.GetMsg(code), Data: data}
}

type ResponseVo struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseVo(code int64, data interface{}) *ResponseVo {
	return &ResponseVo{Code: code, Message: config.GetMsg(code), Data: data}
}

type MoralisVo struct {
	Total    int64  `json:"total"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	Cursor   string `json:"cursor"`
	Result   []struct {
		TokenAddress      string `json:"token_address"`
		TokenId           string `json:"token_id"`
		BlockNumberMinted string `json:"block_number_minted"`
		OwnerOf           string `json:"owner_of"`
		BlockNumber       string `json:"block_number"`
		Amount            string `json:"amount"`
		ContractType      string `json:"contract_type"`
		Name              string `json:"name"`
		Symbol            string `json:"symbol"`
		TokenURI          string `json:"token_uri"`
		Metadata          string `json:"metadata"`
		SyncedAt          string `json:"synced_at"`
		IsValid           int64  `json:"is_valid"`
		Syncing           int64  `json:"syncing"`
		Frozen            int64  `json:"frozen"`
	} `json:"result"`
	Status string `json:"status"`
}
