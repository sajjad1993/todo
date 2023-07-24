package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Succeed bool   `json:"succeed"`
	Message string `json:"message"`
	Results any    `json:"results"`
	Metas   any    `json:"metas"`
}

func GeneralResponse(ctx *gin.Context, statusCode int, succeed bool, message string, results any, metas any) {
	ctx.JSON(
		statusCode,
		Response{
			Succeed: succeed,
			Message: message,
			Results: results,
			Metas:   metas,
		},
	)
}
func AbortResponse(ctx *gin.Context, statusCode int, succeed bool, message string, results any, metas any) {
	ctx.AbortWithStatusJSON(statusCode,
		Response{
			Succeed: succeed,
			Message: message,
			Results: results,
			Metas:   metas,
		})

}

func BadRequestResponse(ctx *gin.Context, error string, data any) {
	GeneralResponse(ctx, http.StatusBadRequest, false, error, data, nil)
}
func FailedResponse(ctx *gin.Context, status int, error string) {
	AbortResponse(ctx, status, false, error, nil, nil)
}

func SuccessResponse(ctx *gin.Context, message string, results any) {
	GeneralResponse(ctx, http.StatusOK, true, message, results, nil)
}

func SuccessResponseWithMeta(ctx *gin.Context, message string, results, metas any) {
	GeneralResponse(ctx, http.StatusOK, true, message, results, metas)
}

func SuccessCreateResponse(ctx *gin.Context) {
	GeneralResponse(ctx, http.StatusCreated, true, "", nil, nil)
}
func OKResponse(ctx *gin.Context) {
	GeneralResponse(ctx, http.StatusOK, true, "", nil, nil)
}

func NotFoundResponse(ctx *gin.Context, error string) {
	GeneralResponse(ctx, http.StatusNotFound, false, error, nil, nil)
}

func UnauthorizedResponse(ctx *gin.Context, statusCode int, error string) {
	GeneralResponse(ctx, http.StatusUnauthorized, false, error, nil, nil)
}
