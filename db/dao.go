package db

import (
	log "github.com/sirupsen/logrus"
	"poolServer/config"
)

func GetPictures(t string) *[]Picture {
	var result []Picture
	res := config.DB.Table("picture").Where("type = ? and deleted = 0", t).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
	}
	return &result
}

func GetSuppliesList(account string, pageSize, pageNum int64) (*[]Pool, int64) {
	var result []Pool
	var totalSize int64
	res := config.DB.
		Select("pool.*").
		Joins("left join token on token.pool_address = pool.address"+
			"and token.deleted = 0"+
			"and token.mortgagor = ?", account).
		Where("pool.deleted = 0 and pool.status = 1 ").
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
	}

	countResult := config.DB.
		Table("pool").
		Joins("left join token on token.pool_address = pool.address"+
			"and token.deleted = 0"+
			"and token.borrower = ?", account).
		Where("pool.deleted = 0 and pool.status = -1 ").
		Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
	}
	return &result, totalSize
}

func GetBorrowsList(account string, pageSize, pageNum int64) (*[]Pool, int64) {
	var result []Pool
	var totalSize int64
	res := config.DB.
		Select("pool.*").
		Joins("left join token on token.pool_address = pool.address"+
			"and token.deleted = 0"+
			"and token.borrower = ?", account).
		Where("pool.deleted = 0 and pool.status = -1 ").
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
	}

	countResult := config.DB.
		Table("pool").
		Joins("left join token on token.pool_address = pool.address"+
			"and token.deleted = 0"+
			"and token.borrower = ?", account).
		Where("pool.deleted = 0 and pool.status = -1 ").
		Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
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
	}

	countResult := config.DB.
		Table("pool").
		Where("deleted = 0").
		Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
	}
	return &result, totalSize
}

func GetPoolDetail(id int64) *[]Pool {
	var result []Pool
	res := config.DB.
		Table("pool").
		Where(" id = ? and deleted = 0", id).
		Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
	}
	return &result
}
