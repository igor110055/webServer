package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"poolServer/config"
	"poolServer/service"
	"poolServer/vo"
	"strconv"
)

func GetDepositList(c *gin.Context) {
	reqVo := vo.ReqVo{}
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	defer c.Set("req", reqVo)
	result := service.GetDepositListService(&reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetBorrowsList(c *gin.Context) {
	reqVo := vo.ReqVo{}
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	defer c.Set("req", reqVo)
	result := service.GetBorrowsListService(&reqVo)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetPoolDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	defer c.Set("req", id)

	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(err)
	}

	result := service.GetPoolDetail(idInt64)
	if result == nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INTERNAL_ERROR, nil))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetPoolList(c *gin.Context) {
	reqVo := vo.ReqVo{}
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	defer c.Set("req", reqVo)
	result := service.GetBorrowsListService(&reqVo)
	c.JSON(http.StatusOK, result)
}

func GetNFTs(c *gin.Context) {
	reqVo := vo.ReqVo{}
	err := c.ShouldBind(&reqVo)
	if err != nil {
		c.JSON(http.StatusOK, vo.NewResponseVo(config.INVALID_PARAMS, nil))
		return
	}
	defer c.Set("req", reqVo)
	result := service.GetBorrowsListService(&reqVo)
	c.JSON(http.StatusOK, result)
}
