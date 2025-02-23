package handler

import (
	"net/http"

	"github.com/KSKigor/api_rest_go/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePAth /api/v1

// @Summary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Succes 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	logger.Infof("request received: %+v", request)

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error create opening: %v", err.Error())
		return
	}
	sendSucces(ctx, "create-opening", opening)
}
