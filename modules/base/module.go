package base

import (
	"github.com/gin-gonic/gin"
)

type (
	ModuleInterface interface {
		GetRoutes() []Route
		GetRepository() interface{}
	}

	Route struct {
		Method  string
		Path    string
		Handler gin.HandlerFunc
	}
)

func AttachRepository(r interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Repository", r)
		c.Next()
	}
}

func InitModule(engine *gin.Engine, module ModuleInterface) {
	engine.Use(AttachRepository(module.GetRepository()))

	for _, route := range module.GetRoutes() {
		engine.Handle(route.Method, route.Path, route.Handler)
	}
}
