package main

import (
	"poolServer/db"
	"poolServer/log"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	//r := router.APIRouter()
	//r.Run(":9999")
	db.GetDepositList("961026", 1, 10)
	//r.RunTLS(":8090", config.PEM, config.KEY)
}
