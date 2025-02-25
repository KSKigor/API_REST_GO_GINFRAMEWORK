package handler

import (
	"net/http"

	"github.com/KSKigor/api_rest_go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePAth /api/v1

// @Summary Update Opening
// @Description Updating an existing opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Indentification"
// @Param opening body UpdateOpeningRequest true "Opening data to update"
// @Succes 200 {object} UpdateOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSucces(ctx, "update opening", opening)

}
