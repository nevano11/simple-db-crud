package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) configureRoutes() {
	s.engine.GET("/", s.helloHandler)

	//s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *ApiServer) helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to site",
	})
}
