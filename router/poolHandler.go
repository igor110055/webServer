package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poolServer/config"
	"poolServer/service"
	"poolServer/vo"
)

func GetDepositList(c *gin.Context) {
	reqVo := vo.InitReqPoolVo()
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	//记录日志
	defer c.Set("req", reqVo)
	result := service.GetDepositListService(reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetBorrowsList(c *gin.Context) {
	reqVo := vo.InitReqPoolVo()
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	//记录日志
	defer c.Set("req", reqVo)

	result := service.GetBorrowsListService(reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetPoolDetail(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	//记录日志
	defer c.Set("req", map[string]string{"address": address})

	////string =>int64
	//idInt64, err := strconv.ParseInt(id, 10, 64)
	//if err != nil {
	//	log.Error(err)
	//}

	result := service.GetPoolDetail(address)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetPoolList(c *gin.Context) {
	reqVo := vo.InitReqVo()
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	//记录日志
	defer c.Set("req", reqVo)

	result := service.GetPoolListService(reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetNFTs(c *gin.Context) {
	reqVo := vo.InitReqNFTVo()
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	//记录日志
	defer c.Set("req", reqVo)

	result := service.GetNFTs(reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetWNFTs(c *gin.Context) {
	reqVo := vo.InitReqWNFTVo()
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	//记录日志
	defer c.Set("req", reqVo)

	result := service.GetWNFTs(reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}
