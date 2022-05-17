package service

import (
	"poolServer/config"
	"poolServer/db"
	"poolServer/vo"
)

func GetDepositListService(req *vo.ReqPoolVo) *vo.ResponsePageVo {
	res, totalSize := db.GetPoolsByQuery(req)
	if res == nil && totalSize == 0 {
		return nil
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, res)
}

func GetBorrowsListService(req *vo.ReqPoolVo) *vo.ResponsePageVo {
	res, totalSize := db.GetPoolsByQuery(req)
	if res == nil && totalSize == 0 {
		return nil
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, res)
}

func GetPoolListService(req *vo.ReqVo) *vo.ResponsePageVo {
	res, totalSize := db.GetPoolList(req.PageVo.PageNum, req.PageVo.PageSize)
	if res == nil && totalSize == 0 {
		return nil
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, res)
}

func GetPoolDetail(id int64) *vo.ResponseVo {
	poolDetail := db.GetPoolById(id)
	poolTokens := db.GetPoolTokenByPoolId(id)

	if poolDetail == nil || poolTokens == nil {
		return nil
	}

	//组装Vo
	for _, value := range *poolTokens {
		//0 need 1 rewards
		if value.Type == "erc20" {
			poolDetail.RewardsTokenName = value.TokenName
			poolDetail.RewardsTokenAddress = value.TokenAddress
		} else if value.Type == "erc721" {
			poolDetail.TokenName = value.TokenName
			poolDetail.TokenAddress = value.TokenAddress
		}
	}
	return vo.NewResponseVo(config.SUCCESS, poolDetail)
}

func GetNFTs(req *vo.ReqNFTVo) *vo.ResponsePageVo {
	//deposited 我质押进去的 查token表 入参poolAddress mortgagor
	//amount 我可以借的nft 查token表 字段poolAddress status=1
	//repay set delegator归还nft 查token表 字段borrower poolAddress
	tokens, totalSize := db.GetToken(req)
	if tokens == nil && totalSize == 0 {
		return nil
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, tokens)
}

func GetWNFTs(req *vo.ReqWNFTVo) *vo.ResponsePageVo {
	tokens, totalSize := db.GetWNFTs(req)
	if tokens == nil && totalSize == 0 {
		return nil
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, tokens)
}
