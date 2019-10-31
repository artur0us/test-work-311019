package router

import "github.com/labstack/echo"

func RegAllRoutes(echoSrv *echo.Echo) {
	RegAccounts(echoSrv)
	RegNotes(echoSrv)
}
