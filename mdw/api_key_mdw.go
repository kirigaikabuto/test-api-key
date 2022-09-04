package mdw

import (
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"github.com/kirigaikabuto/test-api-key/api_key"
	"net/http"
)

type ApiKeyMdw interface {
	MakeApiKeyMiddleware() gin.HandlerFunc
}

type apiKeyMdw struct {
	apiKeyStore api_key.Store
}

func NewApiKeyMdw(apiKeyStore api_key.Store) ApiKeyMdw {
	return &apiKeyMdw{apiKeyStore: apiKeyStore}
}

func (a *apiKeyMdw) MakeApiKeyMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		apiKeyVal := context.Request.Header.Get("Api-Key")
		if apiKeyVal == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, setdata_common.ErrToHttpResponse(ErrNoApiKeyHeaderValue))
			return
		}
		apiKey, err := a.apiKeyStore.GetByKey(apiKeyVal)
		if err != nil && err == api_key.ErrApiKeyNotFound {
			context.AbortWithStatusJSON(http.StatusBadRequest, setdata_common.ErrToHttpResponse(ErrIncorrectApiKey))
			return
		} else if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, setdata_common.ErrToHttpResponse(err))
			return
		}
		context.Set("api_key_id", apiKey.Id)
		context.Next()
	}
}
