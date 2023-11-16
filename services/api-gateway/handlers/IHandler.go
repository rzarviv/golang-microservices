package handlers

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	newHandler(endpointAddr string) (*IHandler, error)
	nddRoutes(router *gin.Engine)
	CloseHandler()
}