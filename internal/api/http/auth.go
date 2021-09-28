package http

import (
	"github.com/gin-gonic/gin"
	"gobackend/internal/controller"
	"gobackend/pkg/errors"
	"gobackend/pkg/logger"
)

var Controller = controller.New(logger.GetLogger())

// Auth - 認証(認可)を行うミドルウェア。
// headerをバリデーションし、アクセスキーの検証を行う。
// 検証に成功した場合、Contextにアクセスキー検証状態を保持。
// 検証に失敗した場合、処理を中断し、Unauthorized Statusでレスポンスする。
// TODO 具体的な処理の定義をControllerに移譲する
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessKey := ""
		var err error

		// If the verification fails, abort and respond with Unauthorized status
		if err != nil {
			aerr := errors.Wrapf(err, "token is invalid, method: Auth").Unauthorized()
			Controller.RespondWithError(c, aerr)
			c.Abort()
			return
		}

		c.Set("accessKey", accessKey)
		c.Next()

	}
}

type Header struct {
	accessKey string `header:"authorization" binding:"required"`
}
