package utils

import (
	"poolServer/vo"

	"time"
)

func GetResponsePageVo(code int64, message string, data *vo.ResponsePageVo) *vo.ResponsePageVo {
	data.Code = code
	data.Message = message
	return data
}

func GetResponseVo(code int64, message string, data interface{}) *vo.ResponseVo {
	responseVo := vo.ResponseVo{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return &responseVo
}

type NftVo struct {
	Id        int64  `json:"ID"` // ID
	ProjectId int64  `json:"projectId"`
	Level     int64  `json:"level"`
	Status    string `json:"status"`
	Address   string `json:"address"`
	Number    int64  `json:"number"`
}

func TimeStampToTime(tm int64) time.Time {
	tm = tm / 1000
	timeFormat := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
	duration, _ := time.ParseInLocation("2006-01-02 15:04:05", timeFormat, time.Local)
	return duration
}
