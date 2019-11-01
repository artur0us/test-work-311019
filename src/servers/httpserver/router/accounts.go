package router

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/artur0us/test-work-311019/entities/accounts"
	"github.com/artur0us/test-work-311019/servers/httpserver/constants"
)

func RegAccounts(echoSrv *echo.Echo) {
	// ======================================= //
	// Authentication
	// ======================================= //

	// Basic authentication
	echoSrv.POST("/api/v1/accounts/auth/basic", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		flag, msg, data := accounts.BasicAuth(c.FormValue("username"), c.FormValue("password"))
		answer.Message = msg
		if !flag {
			answer.EntityStatus = -1
		} else {
			answer.Data = data
		}

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// Token authentication
	echoSrv.POST("/api/v1/accounts/auth/token", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		flag, msg, data := accounts.TokenAuth(c.Request().Header.Get("Token"))
		answer.Message = msg
		if !flag {
			answer.EntityStatus = -1
		} else {
			answer.Data = data
		}

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// ======================================= //
	// Registration
	// ======================================= //

	// Admin registration
	echoSrv.POST("/api/v1/accounts/reg/admin", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	// User registration
	echoSrv.POST("/api/v1/accounts/reg/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})
}
