package routes

import (
	controllers "api-echo/Controllers"
	"api-echo/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/index", controllers.IndexPegawai, middleware.IsAuthenticated)
	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai/store", controllers.StorePegawai)
	e.PUT("/pegawai/update", controllers.UpdatePegawai)
	e.DELETE("/pegawai/delete", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/test-validate", controllers.TestStructValidation)

	return e
}
