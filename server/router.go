package server

import (
	"github.com/labstack/echo"
	//"github.com/tonymtz/gekko/routes"
	"../routes"
)

func router(e *echo.Echo) {
	/*
	 * Login
	 */
	e.GET("/login/:provider", routes.Login.Get)
	e.GET("/login/:provider/callback", routes.Login.Callback)
	e.GET("/login/:provider", routes.Login.Get)

	/**
	*
	* User REST methods
	*/
	e.GET("/user", routes.User.Get)
	e.GET("/user/:id", routes.User.Get)
	e.POST("/user/:id", routes.User.Post)

	/*
	 * Home
	 */
	e.Static("/", "static")
}
