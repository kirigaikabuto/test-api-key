package api_key

import (
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"net/http"
)

type HttpEndpoints interface {
	MakeCreate() gin.HandlerFunc
	MakeList() gin.HandlerFunc
	MakeGetByKey() gin.HandlerFunc
	MakeGetById() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreate() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &CreateApiKeyCommand{}
		err := context.BindJSON(&cmd)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusCreated, resp)
	}
}

func (h *httpEndpoints) MakeList() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &ListApiKeyCommand{}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetByKey() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &GetByApiKeyCommand{}
		cmd.Key = context.Request.URL.Query().Get("key")
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetById() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &GetApiKeyCommand{}
		cmd.Id = context.Request.URL.Query().Get("id")
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}
