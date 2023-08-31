package handler

import (
	"github.com/Kintoho/backend-trainee-assignment-2023/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("", h.createSegment)
			segments.DELETE("/:slug", h.deleteSegment)
		}
		users := api.Group("users")
		{
			users.GET("/:id/show_active_segments", h.showUserActiveSegments)
			users.POST("/:id/add_to_segment", h.addUserToSegment)
			users.DELETE("/:id/delete_from_segment", h.deleteUserFromSegment)
		}
	}

	return router
}
