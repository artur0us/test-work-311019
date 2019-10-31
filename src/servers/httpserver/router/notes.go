package router

import (
	"net/http"
	"strings"

	"../constants"
	"github.com/labstack/echo"
)

func RegNotes(echoSrv *echo.Echo) {
	// ======================================= //
	// Notes
	// ======================================= //

	// Note creation
	echoSrv.POST("/api/v1/notes/", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		//

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// All notes receiver
	echoSrv.GET("/api/v1/notes", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		//

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// Specific note receiver
	echoSrv.GET("/api/v1/notes/:note_id", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		//

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// Note update
	echoSrv.PATCH("/api/v1/notes/:id", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		//

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// All note deletion
	echoSrv.DELETE("/api/v1/notes", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		//

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// Note deletion
	echoSrv.DELETE("/api/v1/notes/:id", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		//

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})
}
