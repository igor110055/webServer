package db

import (
	log "github.com/sirupsen/logrus"
	"poolServer/config"
	"poolServer/vo"
)

func GetPictures(t string) *[]Picture {
	var result []Picture
	res := config.DB.Table("picture").Where("type = ? and deleted = 0", t).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil
	}
	return &result
}

func GetDepositList(account string, pageNum, pageSize int64) (*[]vo.PoolListVo, int64) {
	var result []vo.PoolListVo
	var totalSize int64
	res := config.DB.
		Table("pool p").
		Select("p.id,"+
			"p.name,"+
			"p.uri,"+
			"p.address,"+
			"p.created_time,"+
			"p.updated_time,"+
			"pt.token_name as rewards_token_name,"+
			"pt.token_address as rewards_token_address,"+
			"pt.type,"+
			"pt.wrapper_address,"+
			"pt.wnft_address,"+
			"t.delegator_address ").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0 and pt.status = 1").
		Joins("LEFT JOIN token t ON t.pool_address = p.address and t.deleted = 0").
		Where("p.deleted = 0 and  t.mortgagor = ?", account).
		Order("p.created_time desc").
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}

	countResult := config.DB.
		Table("pool p").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0 and pt.status = 1").
		Joins("LEFT JOIN token t ON t.pool_address = p.address and t.deleted = 0").
		Where("p.deleted = 0 and  t.mortgagor = ?", account).
		Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}
	return &result, totalSize
}

func GetPoolsByQuery(req *vo.ReqPoolVo) (*[]vo.PoolListVo, int64) {
	var result []vo.PoolListVo
	var totalSize int64
	query := config.DB.
		Table("pool p").
		Select("p.id," +
			"p.name," +
			"p.uri," +
			"p.address," +
			"p.created_time," +
			"p.updated_time," +
			"pt.token_name as rewards_token_name," +
			"pt.token_address as rewards_token_address," +
			"pt.type," +
			"pt.wrapper_address," +
			"pt.wnft_address," +
			"t.delegator_address ").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0").
		Joins("LEFT JOIN token t ON t.pool_address = p.address and t.deleted = 0").
		Where("p.deleted = 0")

	if req.Mortgagor != "" {
		query.Where("t.mortgagor = ?", req.Mortgagor)
	}

	if req.Borrower != "" {
		query.Where("t.borrower = ?", req.Borrower)
	}

	if req.DelegatorAddress != "" {
		query.Where("t.delegator_address = ?", req.DelegatorAddress)
	}

	if req.TokenId != "" {
		query.Where("t.token_id = ?", req.TokenId)
	}

	//计算总页数
	countResult := query.Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}

	res := query.Order("p.created_time desc").
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}
	return &result, totalSize
}

func GetBorrowsList(account string, pageNum, pageSize int64) (*[]vo.PoolListVo, int64) {
	var result []vo.PoolListVo
	var totalSize int64
	res := config.DB.
		Table("pool p").
		Select("p.id,"+
			"p.name,"+
			"p.uri,"+
			"p.address,"+
			"p.created_time,"+
			"p.updated_time,"+
			"pt.token_name as rewards_token_name,"+
			"pt.token_address as rewards_token_address,"+
			"pt.type,"+
			"pt.wrapper_address,"+
			"pt.wnft_address,"+
			"t.delegator_address ").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0 and pt.status = 1").
		Joins("LEFT JOIN token t ON t.pool_address = p.address and t.deleted = 0").
		Where("p.deleted = 0 and  t.borrower = ?", account).
		Order("p.created_time desc").
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}

	countResult := config.DB.
		Table("pool p").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0 and pt.status = 1").
		Joins("LEFT JOIN token t ON t.pool_address = p.address and t.deleted = 0").
		Where("p.deleted = 0 and  t.mortgagor = ?", account).
		Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}
	return &result, totalSize
}

//TODO 按照首字母排序
func GetPoolList(pageNum, pageSize int64) (*[]Pool, int64) {
	var result []Pool
	var totalSize int64
	query := config.DB.
		Table("pool").
		Where("deleted = 0")

	//计算总页数
	countResult := query.Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}

	res := query.Order("created_time desc").
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}

	return &result, totalSize
}

func GetPoolById(id int64) *vo.PoolDetailVo {
	var result vo.PoolDetailVo
	res := config.DB.
		Table("pool").
		Where("deleted = 0 and id = ?", id).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil
	}
	return &result
}

func GetPoolTokenByPoolId(id int64) *[]PoolToken {
	var result []PoolToken
	res := config.DB.
		Table("pool_token").
		Where("deleted = 0 and pool_id = ?", id).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil
	}
	return &result
}

func GetToken(req *vo.ReqNFTVo) (*[]vo.TokenVo, int64) {
	var result []vo.TokenVo
	var totalSize int64
	query := config.DB.Table("token").Where("deleted = 0")

	if req.Borrower != "" {
		query.Where("borrower = ? ", req.Borrower)
	}

	if req.Mortgagor != "" {
		query.Where("mortgagor = ? ", req.Mortgagor)
	}

	if req.PoolAddress != "" {
		query.Where("pool_address = ?", req.PoolAddress)
	}
	if req.Status != "" {
		query.Where("status = ?", req.Status)
	}
	//计算总页数
	countResult := query.Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}

	res := query.Order("created_time desc").Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}
	return &result, totalSize
}

func GetWNFTs(req *vo.ReqWNFTVo) (*[]vo.TokenVo, int64) {
	var result []vo.TokenVo
	var totalSize int64

	query := config.DB.Table("token").Where("deleted = 0")

	if req.Account != "" {
		query.Where("owner = ?", req.Account)
	}
	if req.Account != "" {
		query.Where("token_address = ?", req.Address)
	}

	//计算总页数
	countResult := query.Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}

	res := query.Order("updated_time desc").Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}
	return &result, totalSize
}
