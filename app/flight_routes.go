package app

import (
	"ams-service/infrastructure/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterFlightRoutes(router *gin.Engine, flightController *controllers.FlightController) {
	flightRoute := router.Group("/flight")
	{
		flightRoute.GET("/specific", flightController.GetSpecificFlight)
		flightRoute.GET("/all", flightController.GetAllFlights)
	}
}
