package service

import "poolServer/db"

func GetSuppliesListService() {
	db.GetSuppliesList()
}
func GetBorrowsListService() {
	db.GetBorrowsList()
}
func GetPoolListService() {
	db.GetPoolList()
}
func GetPoolDetail() {
	db.GetPoolDetail()
}
