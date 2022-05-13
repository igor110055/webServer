package main

import (
	"poolServer/db"
	"poolServer/log"
	"poolServer/vo"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	//r := router.APIRouter()
	//r.Run(":9999")
	//r.RunTLS(":8090", config.PEM, config.KEY)
	nftVo := vo.ReqNFTVo{
		PoolAddress: "123",
		Mortgagor:   "961026",
	}
	db.GetToken(&nftVo)
}
