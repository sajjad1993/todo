package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/handlers"
	"github.com/sajjad1993/todo/internal/gateway/config"
	"github.com/sajjad1993/todo/pkg/log"
	"net/http"
)

func getRouter(handler *handlers.Handler) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	v1 := router.Group("api/v1/")
	v1.POST("signup", handler.SignUp())
	return router
}

func Serve(handler *handlers.Handler, config config.Config, logger log.Logger) {
	router := getRouter(handler)
	server := &http.Server{
		Addr:              config.GetHTTPServerAddress(),
		Handler:           router,
		ReadHeaderTimeout: config.GetHTTPServerReadHeaderTimeout(),
		ReadTimeout:       config.GetHTTPServerReadTimeout(),
		WriteTimeout:      config.GetHTTPServerWriteTimeout(),
	}
	logger.Info("listening on %s (http)", config.GetHTTPServerAddress())

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error(err.Error())
		}
	}()
}
