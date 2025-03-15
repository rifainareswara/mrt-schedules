package station

import "github.com/gin-gonic/gin"

func Initiate(group gin.RouterGroup) {
	stationService := NewService
	station := group.Group("/stations")
	station.GET("", func(c, *gin.Context) {
		GetStations(c, stationService())
	})
}

func GetStations(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		// handle error
	}

	// handle response
}
