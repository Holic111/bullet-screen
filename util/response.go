package util

import (
	"bullet-screen/model/public"
	"github.com/gin-gonic/gin"
)

func ResponseMap(code int, data map[string]interface{}, ctx *gin.Context) {
	msg := GetMsg(code)
	ctx.JSON(OK, public.ResponseMap{Code: code, Data: data, Msg: msg})
}

func ResponseData(code int, data interface{}, ctx *gin.Context) {
	msg := GetMsg(code)
	ctx.JSON(OK, public.ResponseData{Code: code, Data: data, Msg: msg})
}