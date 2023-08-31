package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userSegments struct {
	User_id int      `json:"user_id"`
	Slugs   []string `json:"slugs"`
}

func (h *Handler) validateUserIdParam(paramName string, c *gin.Context) (int, bool) {
	user_id, err := strconv.Atoi(c.Param(paramName))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return 0, false
	}

	exists, err := h.services.Authorization.UserExists(user_id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 0, false
	}

	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`User with id '%d' does not exist`, user_id))
		return 0, false
	}

	return user_id, true
}

func (h *Handler) showUserActiveSegments(c *gin.Context) {
	user_id, valid := h.validateUserIdParam("id", c)

	if !valid {
		return
	}

	activeSegments, err := h.services.User.GetActiveSegment(user_id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var output userSegments

	output.User_id = user_id
	output.Slugs = []string{}

	for _, m := range activeSegments {
		output.Slugs = append(output.Slugs, m.Slug)
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) addUserToSegment(c *gin.Context) {
	user_id, valid := h.validateUserIdParam("id", c)

	if !valid {
		return
	}

	var input []string

	if err := c.BindJSON(&input); err != nil || len(input) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	for _, slug := range input {
		exists, err := h.services.Segment.Exists(slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' does not exist`, slug))
			return
		}

		exists, err = h.services.User.SegmentRelationExists(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Relation ('%d':'%s') already exists`, user_id, slug))
			return
		}
	}

	for _, slug := range input {
		_, err := h.services.User.AddToSegment(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, statusResponse{"ok"})
}

func (h *Handler) deleteUserFromSegment(c *gin.Context) {
	user_id, valid := h.validateUserIdParam("id", c)

	if !valid {
		return
	}

	var input []string

	if err := c.BindJSON(&input); err != nil || len(input) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	for _, slug := range input {
		exists, err := h.services.Segment.Exists(slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Segment with name '%s' does not exist`, slug))
			return
		}

		exists, err = h.services.User.SegmentRelationExists(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !exists {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf(`Relation ('%d':'%s') does not exist`, user_id, slug))
			return
		}
	}

	for _, slug := range input {
		err := h.services.User.DeleteSegmentRelation(user_id, slug)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, statusResponse{"ok"})
}
