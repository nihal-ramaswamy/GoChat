package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleErrorAndAbortWithError(ctx *gin.Context, err error, log *zap.Logger) {
	if nil != err {
		log.Error(err.Error())
		err = ctx.AbortWithError(http.StatusInternalServerError, err)
		if nil != err {
			log.Error("error producing error", zap.Error(err))
		}
	}
}
