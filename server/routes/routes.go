package routes

import (
	"api-sample/controllers"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"

	"github.com/gin-gonic/gin"
)

func ConfigRoute(route *gin.Engine) *gin.Engine {
	route.Use(gintrace.Middleware("api-sample-totvs"))
	main := route.Group("api/v1")
	{
		person := main.Group("person")
		{
			person.GET("/", controllers.ShownAllPeople)
			person.GET("/:id", controllers.ShownPerson)
			person.POST("/", controllers.CreatePerson)
			person.PUT("/", controllers.UpdatePerson)
			person.DELETE("/:id", controllers.DeletePerson)
		}
		server := main.Group("server")
		{
			server.GET("/", controllers.ServerInformation)
		}
	}
	return route
}
