package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"poolServer/config"
	"poolServer/db"
	"poolServer/utils"
	"poolServer/vo"
	"strings"
	"time"
)

func GetTimeStamp(c *gin.Context) {
	unix := time.Now().Unix()
	defer c.Set("req", unix)
	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), unix))
}

func GetTokenByAddress(c *gin.Context) {
	address := c.Query("address")
	nft := c.Query("nft")
	if address == "" || nft == "" {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.INVALID_PARAMS, config.GetMsg(config.INVALID_PARAMS), nil))
		return
	}
	defer c.Set("req", map[string]interface{}{"address": address, "nft": nft})
	chainId := "0x61"
	req, _ := http.NewRequest("GET", "https://deep-index.moralis.io/api/v2/"+address+"/nft/"+nft+"?chain="+
		chainId, nil)
	// 设置请求头
	req.Header.Set("X-API-Key", "dXEk5PiV0qmNUHfKeEWdrf9Iu8OpanzXHRi3tLzAxxf9rwShOECim67VeQJdjwbv")
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR, config.GetMsg(config.ERROR), nil))
		return
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		//var data interface{}
		//json.Unmarshal(body, &data)
		var data vo.MoralisVo
		json.Unmarshal(body, &data)

		//临时处理主网显示Badge改为Warrior
		if data.Total > 0 {
			for _, s := range data.Result {
				if strings.EqualFold(s.TokenAddress, "0x27a9ffd030bf756965ce2a054d808435d2c61e2e") {
					s.Name = "Warrior"
					s.Symbol = "Warrior"
					s.TokenURI = "https://cloudflare-ipfs.com/ipfs/bafybeihaxnzvymml2ftjgdcfl4khey6owkfryhogunqfyuq4aqsaf5qhzy"
				}
			}
		}
		pageVo := &vo.ResponsePageVo{
			PageNum:   data.Page,
			PageSize:  data.PageSize,
			TotalSize: data.Total,
			Data:      data.Result,
		}
		c.JSON(http.StatusOK, utils.GetResponsePageVo(config.SUCCESS, config.GetMsg(config.SUCCESS), pageVo))
	}
}

func GetPictures(c *gin.Context) {
	t := c.Query("type")
	if t == "" {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.INVALID_PARAMS, config.GetMsg(config.INVALID_PARAMS), nil))
		return
	}
	defer c.Set("req", map[string]interface{}{"type": t})
	result := db.GetPictures(t)
	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), result))
}
