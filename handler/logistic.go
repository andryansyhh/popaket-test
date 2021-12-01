package handler

import (
	"popaket/helper"
	"popaket/service"

	"github.com/gin-gonic/gin"
)

type logisticHandler struct {
	logisticService service.LogisticService
}

func NewLogisticHandler(logisticService service.LogisticService) *logisticHandler {
	return &logisticHandler{logisticService}
}

// ShowAllLogistics godoc
// @Security Auth
// @Summary Get All Logistics
// @Description Get All Logistics
// @Tags logistic
// @Accept json
// @Produce json
// @success 200 {object} helper.Response
// @Router /popaket/logistic [get]
func (h *logisticHandler) ShowAllLogistiHandler(c *gin.Context) {
	logistic, err := h.logisticService.ShowAllLogistic()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})
		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all Logistics", 200, "status OK", logistic)
	c.JSON(200, response)
}

// ShowLogisticWithParam godoc
// @Summary Get Logistic by param
// @Description Get logistic by param
// @Tags logistic
// @Accept json
// @Produce json
// @Param origin_name query string true "origin_name"
// @Param destination_name query string true "destination_name"
// @Success 200 {object} helper.Response
// @Router /popaket/logistic/filter [get]
func (h *logisticHandler) ShowAllLogistiByParamHandler(c *gin.Context) {
	destination := c.Params.ByName("destination_name")
	origin := c.Params.ByName("origin_name")

	logistic, err := h.logisticService.ShowLogisticByParam(destination, origin)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})
		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all Logistics", 200, "status OK", logistic)
	c.JSON(200, response)
}
