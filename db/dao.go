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

func GetBorrowsList(account string, pageSize, pageNum int64) (*[]vo.PoolListVo, int64) {
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

func GetPoolList(pageSize, pageNum int64) (*[]Pool, int64) {
	var result []Pool
	var totalSize int64
	res := config.DB.
		Table("pool").
		Where("deleted = 0").
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}

	countResult := config.DB.
		Table("pool").
		Where("deleted = 0").
		Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
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
