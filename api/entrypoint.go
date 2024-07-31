package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	r := app.Group("/api")

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
