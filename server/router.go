package server

import (
	"github.com/labstack/echo"
	"github.com/tonymtz/gekko/routes"
)

func router(e *echo.Echo) {
	/*
	 * Login
	 */
	e.GET("/login/:provider", routes.Login.Get)
	e.GET("/login/:provider/callback", routes.Login.Callback)
	e.GET("/login/:provider", routes.Login.Get)

	/*
	 * Home
	 */
	e.Static("/", "static")
}
