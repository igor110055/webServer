package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"poolServer/config"
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

func GetMoralis(address, nft string) *vo.ResponsePageVo {
	var data vo.MoralisVo
	req, _ := http.NewRequest("GET", config.MORALIS.Url+address+"/nft/"+nft+"?chain="+
		config.MORALIS.ChainId, nil)
	// 设置请求头
	req.Header.Set("X-API-Key", config.MORALIS.Key)
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		return nil
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &data)
	}

	//重新组装数据
	var tokenVos []vo.TokenVo
	for _, v := range data.Result {
		tokenVo := vo.TokenVo{
			TokenId:      v.TokenId,
			TokenAddress: v.TokenAddress,
			TokenSymbol:  v.Symbol,
			TokenName:    v.Name,
			TokenUri:     v.TokenURI.(string),
		}
		tokenVos = append(tokenVos, tokenVo)
	}
	return vo.NewResponsePageVo(data.Page, data.PageSize, data.Total, config.SUCCESS, tokenVos)
}
