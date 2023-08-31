package handler

import (
	"net/http"

	"github.com/Kintoho/backend-trainee-assignment-2023/structure"

	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input structure.User

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, successBaseResponse{
		Id: id,
	})
}
