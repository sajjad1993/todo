package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/services/gateway/adapter/restapi/handlers"
	"github.com/sajjad1993/todo/services/gateway/adapter/restapi/middleware"
	"github.com/sajjad1993/todo/services/gateway/app/query"
	"github.com/sajjad1993/todo/services/gateway/config"
	"net/http"
)

func getRouter(handler *handlers.Handler, authService *query.CheckToken) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	v1 := router.Group("api/v1/")
	v1.POST("signup", handler.SignUp())
	v1.POST("signin", handler.SignIn())
	v1.POST("todo-list", middleware.CheckToken(authService), handler.CreateTodoList())
	v1.GET("todo-list", middleware.CheckToken(authService), handler.ListTodoList())
	v1.DELETE("todo-list/:id", middleware.CheckToken(authService), handler.DeleteTodoList())
	v1.PATCH("todo-list/:id", middleware.CheckToken(authService), handler.UpdateTodoList())
	v1.POST("todo", middleware.CheckToken(authService), handler.CreateTodo())
	v1.PATCH("todo/:id", middleware.CheckToken(authService), handler.UpdateTodo())
	v1.DELETE("todo/:id", middleware.CheckToken(authService), handler.DeleteTodo())
	return router
}

func Serve(handler *handlers.Handler, config config.Config, logger log.Logger, authService *query.CheckToken) {
	router := getRouter(handler, authService)
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
