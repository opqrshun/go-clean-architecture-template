package extapi

import (
	// "log"
	// "net/http"

	mapper "github.com/ttaki/go-clean-architecture-template/internal/infrastructure/extapi/extapimapper"
	"github.com/ttaki/go-clean-architecture-template/model"
)

type Auth struct {
	*APIHandler
}

func NewAuth() *Auth {
	return &Auth{
		APIHandler: NewAPIHandler(),
	}
}

// Process - 破壊対策レイヤーのサンプルを示す。
// コアモデル(アプリケーション固有)と外部モデル（レガシーアプリケーションのモデル）のマッピングを行う。外部モデルの複雑な要素はアプリケーション内部に入れない。
// 外部用のモデルを生成し、レガシーアプリケーション関連の処理を行うメソッドに移譲する。
func (h *Auth) Process(coreModel model.Base) (model.Base, error) {
	//.コアモデルを外部モデルにマッピング
	var u mapper.User

	// 処理の移譲
	_, err := h.service.Complex(&u)
	if err != nil {
		return model.Base{}, BuildAPIError(err, "")
	}

	//.外部モデルをコアモデルにマッピング
	var resultCoreModel model.Base
	return resultCoreModel, nil
}
