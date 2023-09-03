package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"simple-db-crud/internal/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", h.welcome)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		cars := api.Group("/car")
		{
			cars.GET("/", h.getAllCars)
			cars.GET("/:id", h.getCarById)
			cars.POST("/", h.createCar)
			cars.PUT("/:id", h.editCar)
			cars.DELETE("/:id", h.deleteCar)
		}
	}
	return router
}

func (h *Handler) welcome(ctx *gin.Context) {
	io.WriteString(ctx.Writer, "Welcome")
}
