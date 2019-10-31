package router

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegAccounts(echoSrv *echo.Echo) {
	// Authentication
	echoSrv.POST("/api/v1/accounts/auth/basic", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	// Registration
	echoSrv.POST("/api/v1/accounts/reg/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})
	echoSrv.POST("/api/v1/accounts/reg/admin", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})
}
