package service

import (
	"poolServer/config"
	"poolServer/db"
	"poolServer/vo"
)

func GetDepositListService(req *vo.ReqVo) *vo.ResponsePageVo {
	res, totalSize := db.GetDepositList(req.Address, req.PageVo.PageNum, req.PageVo.PageSize)
	if res == nil && totalSize == 0 {
		return nil
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, res)
}

func GetBorrowsListService(req *vo.ReqVo) *vo.ResponsePageVo {

	switch req.Type {

	case "1":
		//要查moerios接口的 wallet
		break
	case "2":
		//deposited 我质押进去的 查token表 字段mortgagor
		//myWNFT 转换出来的专门查wnft表
		//amount 借nft 查token表 字段pool_address status=1
		//repay set delegator归还nft 查token表 字段borrower = account status=-1 pool_address
		break

	}

	res, totalSize := db.GetBorrowsList(req.Address, req.PageVo.PageNum, req.PageVo.PageSize)
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
		if value.Status == "1" {
			poolDetail.RewardsTokenName = value.TokenName
			poolDetail.RewardsTokenAddress = value.TokenAddress
		} else if value.Status == "0" {
			poolDetail.TokenName = value.TokenName
			poolDetail.TokenAddress = value.TokenAddress
		}
	}
	return vo.NewResponseVo(config.SUCCESS, poolDetail)
}
