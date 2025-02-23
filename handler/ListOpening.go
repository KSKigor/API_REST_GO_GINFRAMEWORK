package handler

import (
	"net/http"

	"github.com/KSKigor/api_rest_go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePAth /api/v1

// @Summary List openings
// @Description list existents Openings
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Succes 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSucces(ctx, "list-openings", openings)
}
