package handler

import (
	"fmt"
	"net/http"

	"github.com/Kintoho/backend-trainee-assignment-2023/structure"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input structure.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	exists, err := h.services.Segment.Exists(input.Slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' already exists`, input.Slug))
		return
	}

	id, err := h.services.Segment.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, successBaseResponse{
		Id: id,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {
	slug := c.Param("slug")

	exists, err := h.services.Segment.Exists(slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' does not exist`, slug))
		return
	}

	h.services.Segment.Delete(slug)

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
