package controller

import (
	"github.com/gin-gonic/gin"
)

func GetUser(ct *gin.Context){
	ct.String(200, "scusses")
}