package endpoints_permission

import (
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"net/http"
)

type HttpEndpoints interface {
	MakeCreate() gin.HandlerFunc
	MakeList() gin.HandlerFunc
	MakeGetByApiKeyId() gin.HandlerFunc
	MakeGetById() gin.HandlerFunc
	MakeAddEndpoint() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreate() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &CreateEndpointsPermissionCommand{}
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
		cmd := &ListEndpointsPermissionsCommand{}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeGetByApiKeyId() gin.HandlerFunc {
	return func(context *gin.Context) {
		apiKeyId := context.Request.URL.Query().Get("api_key_id")
		cmd := &GetByApiKeyIdEndpointsPermissionsCommand{apiKeyId}
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
		id := context.Request.URL.Query().Get("id")
		cmd := &GetByIdEndpointsPermissionCommand{id}
		resp, err := h.ch.ExecCommand(cmd)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.JSON(http.StatusOK, resp)
	}
}

func (h *httpEndpoints) MakeAddEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &AddEndpointToEndpointsPermissionsCommand{}
		id := context.Request.URL.Query().Get("id")
		cmd.Id = id
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
		context.JSON(http.StatusOK, resp)
	}
}
