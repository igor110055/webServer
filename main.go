package main

import (
	"poolServer/log"
	"poolServer/router"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	r := router.APIRouter()
	r.Run(":8200")
	//r.RunTLS(":8090", config.PEM, config.KEY)
}
