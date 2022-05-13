package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"poolServer/config"
	"poolServer/db"
	"poolServer/vo"
	"time"
)

func GetTimeStamp(c *gin.Context) {
	unix := time.Now().Unix()
	defer c.Set("req", unix)
	c.JSON(http.StatusOK, vo.NewResponseVo(config.SUCCESS, unix))
}

func GetTokenByAddress(c *gin.Context) {
	var data vo.MoralisVo
	reqVo := vo.ReqMoralisVo{}
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}
	defer c.Set("req", reqVo)
	req, _ := http.NewRequest("GET", config.MORALIS.Url+reqVo.Address+"/nft/"+reqVo.TokenAddress+"?chain="+
		config.MORALIS.ChainId+"&cursor="+reqVo.Cursor+"&limit="+reqVo.Limit, nil)
	// 设置请求头
	req.Header.Set("X-API-Key", config.MORALIS.Key)
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		return
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
	c.JSON(http.StatusOK, vo.NewResponsePageVo(data.Page, data.PageSize, data.Total, config.SUCCESS, tokenVos))
}

func GetPictures(c *gin.Context) {
	t := c.Query("type")
	if t == "" {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}
	defer c.Set("req", map[string]interface{}{"type": t})
	result := db.GetPictures(t)
	c.JSON(http.StatusOK, vo.NewResponseVo(config.SUCCESS, result))
}
