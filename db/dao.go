package db

import (
	log "github.com/sirupsen/logrus"
	"poolServer/config"
	"poolServer/vo"
)

//func GetPictures(t string) *[]Picture {
//	var result []Picture
//	res := config.DB.Table("picture").Where("type = ? and deleted = 0", t).Find(&result)
//	if res.Error != nil {
//		log.Error(res.Error)
//		return nil
//	}
//	return &result
//}

func GetPoolsByQuery(req *vo.ReqPoolVo) (*[]vo.DepositListVo, int64) {
	var result []vo.DepositListVo
	var totalSize int64
	query := config.DB.
		Table("pool p").
		Select("p.id," +
			"p.name," +
			"p.url," +
			"p.owner," +
			"p.address," +
			"p.created_time," +
			"p.updated_time," +
			"pt.token_name as rewards_token_name," +
			"pt.token_address as rewards_token_address," +
			"pt.type," +
			"pt.wrapper_address," +
			"pt.wnft_address," +
			"t.delegator_address ").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0 and type = 'erc20'").
		Joins("LEFT JOIN token t ON t.pool_address = p.address and t.deleted = 0").
		Where("p.deleted = 0").Group("p.address")

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

//TODO 按照首字母排序
func GetPoolList(pageNum, pageSize int64) (*[]PoolListDto, int64) {
	var result []PoolListDto
	var totalSize int64
	query := config.DB.
		Table("pool p").
		Select("p.id," +
			"p.name," +
			"p.url," +
			"p.owner," +
			"p.address," +
			"p.rate," +
			"p.created_time," +
			"p.updated_time," +
			"pt.token_name," +
			"pt.token_address," +
			"pt.type").
		Joins("LEFT JOIN pool_token pt ON pt.pool_id = p.id and pt.deleted = 0 and type = 'erc721'").
		Where("p.deleted = 0").Group("p.address")

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

func GetPoolById(id int64) *Pool {
	var result Pool
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
	query := config.DB.Table("token t").
		Select("t.id,"+
			"t.token_id,t.token_address,t.token_uri,t.borrower,t.mortgagor,t.status,t.delegator_address,pt.token_name").
		Where("t.deleted = 0 and t.pool_address = ?", req.PoolAddress).
		Joins("left join pool_token pt on t.token_address = pt.token_address")

	if req.Borrower != "" {
		query.Where("t.borrower = ? ", req.Borrower)
	}

	if req.Mortgagor != "" {
		query.Where("t.mortgagor = ? ", req.Mortgagor)
	}

	if req.Status != "" {
		query.Where("t.status = ? and t.borrower = '' ", req.Status)
	}
	//计算总页数
	countResult := query.Count(&totalSize)
	if countResult.Error != nil {
		log.Error(countResult.Error)
		return nil, 0
	}

	res := query.Order("t.created_time desc").Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, 0
	}
	return &result, totalSize
}

func GetWNFTs(req *vo.ReqWNFTVo) (*[]vo.TokenVo, int64) {
	var result []vo.TokenVo
	var totalSize int64

	query := config.DB.Table("wnft").Where("deleted = 0")

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

func GetTotalBorrower(poolAddress string) (int64, int64) {
	var total, borrowers int64
	b := config.DB.Table("token").Where("deleted = 0 and status = -1 and pool_address = ?", poolAddress).
		Count(&borrowers)

	if b.Error != nil {
		log.Error(b.Error)
	}

	t := config.DB.Table("token").Where("deleted = 0 and status = 1 and pool_address = ?", poolAddress).
		Count(&total)

	if t.Error != nil {
		log.Error(b.Error)
	}
	return borrowers, total
}
