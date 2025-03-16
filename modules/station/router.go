package station

import (
	"github.com/gin-gonic/gin"
	"mrt-schedules/common/response"
	"net/http"
)

func Initiate(group *gin.RouterGroup) {
	stationService := NewService
	station := group.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetStations(c, stationService())
	})
	station.GET("/:id", func(c *gin.Context) {
		GetStationById(c, stationService())
	})
	station.GET("/:id/schedule", func(c *gin.Context) {
		GetStationSchedule(c, stationService())
	})
}

func GetStations(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	c.JSON(http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Success GET all stations",
			Data:    datas,
		})
}

func GetStationById(c *gin.Context, service Service) {
	id := c.Param("id")
	data, err := service.GetStationById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	c.JSON(http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Success GET station by ID",
			Data:    data,
		})
}

func GetStationSchedule(c *gin.Context, service Service) {
	id := c.Param("id")
	data, err := service.GetStationSchedule(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	c.JSON(http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Success GET station schedule",
			Data:    data,
		})
}
