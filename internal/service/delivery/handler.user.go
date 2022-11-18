package delivery

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/feature"
	sconstants "github.com/vldcreation/simple-auth-golang/internal/service/constants"
	"github.com/vldcreation/simple-auth-golang/internal/service/gin_"
)

func SetupUser(ctx context.Context, feat feature.SetupUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response gin_.Response
		var err error

		var req feature.SetupUserRequest
		if err = entity.JSON.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
			response, err := gin_.NewResponseError(err, 500, sconstants.MsgUnexpected, sconstants.Text(500), sconstants.FailedToUnmarshall, gin_.UnwrapFirstError(err))
			ctx.IndentedJSON(response.StatusCode, err)

			return
		}

		log.Printf("[SetupUser] - req: %+v", req)

		res, htppcode, err := feat.SetupUser(ctx, req)
		if err != nil {
			response, err = gin_.NewResponseError(err, htppcode, sconstants.MsgUnexpected, sconstants.Text(500), sconstants.Text(500), gin_.UnwrapFirstError(err))

			ctx.IndentedJSON(response.StatusCode, err)

			return
		}

		response = gin_.Response{Message: sconstants.MsgSuccess, StatusCode: 200, Data: res}

		ctx.IndentedJSON(response.StatusCode, response)
	}
}

func AccountLogin(ctx context.Context, feat feature.AccountLogin) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response gin_.Response
		var err error

		var req feature.AccountLoginRequest
		if err = entity.JSON.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
			response, err := gin_.NewResponseError(err, 500, sconstants.MsgUnexpected, sconstants.Text(500), sconstants.FailedToUnmarshall, gin_.UnwrapFirstError(err))
			ctx.IndentedJSON(response.StatusCode, err)

			return
		}

		log.Printf("[AccountLogin] - req: %+v", req)

		res, htppcode, err := feat.AccountLogin(ctx, req)
		if err != nil {
			response, err = gin_.NewResponseError(err, htppcode, sconstants.MsgUnexpected, sconstants.Text(htppcode), sconstants.Text(htppcode), gin_.UnwrapFirstError(err))

			ctx.IndentedJSON(response.StatusCode, err)

			return
		}

		response = gin_.Response{Message: sconstants.MsgSuccess, StatusCode: 200, Data: res}

		ctx.IndentedJSON(response.StatusCode, response)
	}
}
