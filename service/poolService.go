package service

import (
	"poolServer/config"
	"poolServer/db"
	"poolServer/utils"
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
	poolDtos, totalSize := db.GetPoolList(req.PageVo.PageNum, req.PageVo.PageSize)
	if poolDtos == nil && totalSize == 0 {
		return nil
	}

	var res []*vo.PoolListVo
	//封装数据，计算apr
	for _, listVo := range *poolDtos {
		poolListVo := utils.TransferPoolListVo(listVo)
		res = append(res, poolListVo)
	}
	return vo.NewResponsePageVo(req.PageVo.PageNum, req.PageVo.PageSize, totalSize, config.SUCCESS, res)
}

func GetPoolDetail(id int64) *vo.ResponseVo {
	poolDetail := db.GetPoolById(id)

	res := utils.TransferPoolDetailVo(poolDetail)

	return vo.NewResponseVo(config.SUCCESS, res)
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
