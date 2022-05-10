package service

import (
	"poolServer/db"
	"poolServer/vo"
)

func GetSuppliesListService(req *vo.ReqVo) {
	db.GetSuppliesList(req.Account, req.PageVo.PageSize, req.PageVo.PageNum)
}

func GetBorrowsListService(req *vo.ReqVo) {
	db.GetBorrowsList(req.Account, req.PageVo.PageSize, req.PageVo.PageNum)
}

func GetPoolListService(req *vo.ReqVo) {
	db.GetPoolList(req.PageVo.PageSize, req.PageVo.PageNum)
}

func GetPoolDetail(id int64) {
	db.GetPoolDetail(id)
}
