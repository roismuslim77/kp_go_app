package infrahttp

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Middleware struct {
}

func NewBuilderMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) AddHeader() gin.HandlerFunc {
	log.Println("header")
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}

func (m Middleware) GetSessionCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
