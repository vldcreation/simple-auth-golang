package delivery

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
)

func (ox *GinObject) InitRoutes(context context.Context) {
	env := entity.ENV()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.POST("/user/account-creation", SetupUser(context, ox.features.SetupUser))
	router.POST("/user/account-login", AccountLogin(context, ox.features.AccountLogin))

	router.Run(env.Get(constants.AppHost) + ":" + env.Get(constants.AppPort))
}
