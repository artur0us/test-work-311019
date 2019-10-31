package httpserver

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"

	"./router"
)

var EchoServer *echo.Echo

func Start() {
	EchoServer = echo.New()

	// EchoServer.HideBanner = true;
	EchoServer.Debug = false
	EchoServer.Logger.SetLevel(echoLog.OFF)
	EchoServer.Logger.SetOutput(ioutil.Discard)

	EchoServer.Use(middleware.Recover())
	EchoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))

	router.RegAllRoutes(EchoServer)

	echoServerExtendedCfg := &http.Server{
		Addr:         ":10843",
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	EchoServer.Logger.Fatal(EchoServer.StartServer(echoServerExtendedCfg))
}
