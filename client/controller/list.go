package controller

import (
	"github.com/gin-gonic/gin"
)

func GetList(ct *gin.Context){
	ct.String(200, "List Success")
}