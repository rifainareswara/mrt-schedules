package station

import "github.com/gin-gonic/gin"

func Initiate(group gin.RouterGroup) {
	station := group.Group("/stations")

	station.GET("", GetStations)
	station.GET("/:id", GetStation)
	station.POST("", CreateStation)
	station.PUT("/:id", UpdateStation)
	station.DELETE("/:id", DeleteStation)
}
