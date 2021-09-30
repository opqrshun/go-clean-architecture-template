package extapi

import (
	// "log"
	"net/http"

	"github.com/ttaki/go-clean-architecture-template/internal/config"
	apimodel "github.com/ttaki/go-clean-architecture-template/internal/infrastructure/extapi/extapimodel"
	sv "github.com/ttaki/go-clean-architecture-template/internal/infrastructure/extapi/extapiservice"
	"github.com/ttaki/go-clean-architecture-template/model"
	"github.com/ttaki/go-clean-architecture-template/pkg/errors"
)

type APIHandler struct {
	client  *http.Client
	URL     string
	service *sv.Service
}

var client *http.Client

func init() {
	client = newClient()
}

func newClient() *http.Client {
	return &http.Client{}
}

func getClient() *http.Client {
	return client
}

func NewAPIHandler() *APIHandler {
	config := config.GetExtAPIConfig()

	return &APIHandler{
		client:  getClient(),
		URL:     config.URL,
		service: sv.New(getClient()),
	}
}

func (h *APIHandler) Ping() error {
	req, err := http.NewRequest("Ping", h.URL, nil)
	if err != nil {
		return BuildAPIError(err, "Ping")
	}

	_, err = client.Do(req)
	if err != nil {
		return BuildAPIError(err, "Ping")
	}

	return nil
}

// Process - 破壊対策レイヤーのサンプルを示す。
// コアモデル(アプリケーション固有)と外部モデル（レガシーアプリケーションのモデル）のマッピングを行う。外部モデルの複雑な要素はアプリケーション内部に入れない。
// 外部用のモデルを生成し、レガシーアプリケーション関連の処理を行うメソッドに移譲する。
func (h *APIHandler) Process(coreModel model.Base) (model.Base, error) {
	//.コアモデルを外部モデルにマッピング
	var u apimodel.User

	// 処理の移譲
	_, err := h.service.Complex(&u)
	if err != nil {
		return model.Base{}, BuildAPIError(err, "")
	}

	//.外部モデルをコアモデルにマッピング
	var resultCoreModel model.Base
	return resultCoreModel, nil
}

func BuildAPIError(err error, methodName string) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, http.ErrNotSupported) {
		return errors.Wrapf(err, "extapi error, method: "+methodName).Internal()
	}
	return errors.Wrapf(err, "extapi error, method: "+methodName).Internal()
}
