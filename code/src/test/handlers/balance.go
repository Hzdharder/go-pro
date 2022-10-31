package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test/utils"
	"test/xhttp"
)

func OnGetAvailableBalance(ctx *gin.Context) {
	// 获取token
	//httpClient := xhttp.NewClient()
	//httpClient.Header.Set("Authorization",utils.Authorization)
	//_, bs, err := httpClient.Post(utils.AuthUrl).EndBytes()
	//if err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "获取token失败"})
	//	return
	//}
	//var respAuth utils.RespAuth
	//if err := json.Unmarshal(bs, &respAuth); err != nil {
	//	log.Print("解析获取token返回值失败")
	//}
	//
	//// 获取可用余额
	//httpClient.Header.Set("Authorization", "bearer " + respAuth.ResponseBody.AccessToken)
	//_, bs, err = httpClient.Get(utils.BalanceUrl).EndBytes()
	//if err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "获取可用余额失败"})
	//	return
	//}
	//var respBalance utils.RespBalance
	//if err := json.Unmarshal(bs, &respBalance); err != nil {
	//	log.Print("解析获取可用余额失败")
	//}
	fmt.Print("1")
	//ctx.JSON(http.StatusOK, gin.H{"errno": "0", "errmsg": "获取成功", "data": respBalance.ResponseBody})
	ctx.JSON(http.StatusOK, gin.H{"errno": "0", "errmsg": "获取成功", "data": "123"})
}

