package interfaces

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	Pattern() string
	Handler() gin.HandlerFunc
}
