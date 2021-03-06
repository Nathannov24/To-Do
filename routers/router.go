package routers

import (
	"net/http"
	"todo/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete},
	}))

	e.POST("/post", controllers.PostActivityController)
	e.GET("/", controllers.GetActivityController)
	e.DELETE("/delete/:id", controllers.DeleteActivityController)
	e.PUT("/update-status/:id", controllers.UpdateStatusController)

	return e
}
