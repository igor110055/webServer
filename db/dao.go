package db

import (
	log "github.com/sirupsen/logrus"
	"poolServer/config"
)

func InsertToken() {

}

func UpdatedToken() {

}

func GetPictures(t string) *[]Picture {
	var result []Picture
	res := config.DB.Table("PICTURE").Where("type = ? and deleted = 0", t).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
	}
	return &result
}

func GetSuppliesList() {

}

func GetBorrowsList() {

}

func GetPoolList() {

}

func GetPoolDetail() {

}
