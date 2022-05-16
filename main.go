package main

import (
	"poolServer/log"
	"poolServer/router"
	"time"
)

// @title Pool
// @version 1.0
// @description pool api
// @contact.name p

// @host 8.210.48.242
// @BasePath 这里写base path
func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	r := router.APIRouter()
	r.Run(":8200")
	//r.RunTLS(":8090", config.PEM, config.KEY)
}
