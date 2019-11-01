package router

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/artur0us/test-work-311019/entities/accounts"
	"github.com/artur0us/test-work-311019/entities/notes"
	notesMdls "github.com/artur0us/test-work-311019/entities/notes/mdls"
	"github.com/artur0us/test-work-311019/servers/httpserver/constants"
	"github.com/labstack/echo"
)

func RegNotes(echoSrv *echo.Echo) {
	// ======================================= //
	// Notes
	// ======================================= //

	// Note creation
	echoSrv.POST("/api/v1/notes", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		if !accounts.IsTokenAuthed(c.Request().Header.Get("Token")) {
			answer.EntityStatus = -1
			answer.Message = "Authentication error!"
		} else {
			note := notesMdls.NewNote{
				Title: c.FormValue("title"),
				Body:  c.FormValue("body"),
			}
			if note.Title == "" || note.Body == "" {
				answer.EntityStatus = -2
				answer.Message = "Title or body of creating note is empty!"
			} else {
				flag, msg := notes.Create(c.Request().Header.Get("Token"), note)
				answer.Message = msg
				if !flag {
					answer.EntityStatus = -3
				}
			}
		}

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// All notes receiver
	echoSrv.GET("/api/v1/notes", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		if !accounts.IsTokenAuthed(c.Request().Header.Get("Token")) {
			answer.EntityStatus = -1
			answer.Message = "Authentication error!"
		} else {
			flag, msg, data := notes.GetAll(c.Request().Header.Get("Token"))
			answer.Message = msg
			if !flag {
				answer.EntityStatus = -2
			} else {
				answer.Data = data
			}
		}

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// Specific note receiver
	echoSrv.GET("/api/v1/notes/:note_id", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		if !accounts.IsTokenAuthed(c.Request().Header.Get("Token")) {
			answer.EntityStatus = -1
			answer.Message = "Authentication error!"
		} else {
			noteID, err := strconv.ParseInt(c.Param("note_id"), 10, 64)
			if err != nil {
				answer.EntityStatus = -2
				answer.Message = "Invalid note id specified in URL!"
			} else {
				flag, msg, data := notes.GetByID(c.Request().Header.Get("Token"), noteID)
				answer.Message = msg
				if !flag {
					answer.EntityStatus = -3
				} else {
					answer.Data = data
				}
			}
		}

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

		if !accounts.IsTokenAuthed(c.Request().Header.Get("Token")) {
			answer.EntityStatus = -1
			answer.Message = "Authentication error!"
		} else {
			flag, msg := notes.DeleteAll(c.Request().Header.Get("Token"))
			answer.Message = msg
			if !flag {
				answer.EntityStatus = -2
			}
		}

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})

	// Note deletion
	echoSrv.DELETE("/api/v1/notes/:id", func(c echo.Context) error {
		answer := constants.TemplateReqAnswer()

		if !accounts.IsTokenAuthed(c.Request().Header.Get("Token")) {
			answer.EntityStatus = -1
			answer.Message = "Authentication error!"
		} else {
			noteID, err := strconv.ParseInt(c.Param("note_id"), 10, 64)
			if err != nil {
				answer.EntityStatus = -2
				answer.Message = "Invalid note id specified in URL!"
			} else {
				flag, msg := notes.DeleteByID(c.Request().Header.Get("Token"), noteID)
				answer.Message = msg
				if !flag {
					answer.EntityStatus = -3
				}
			}
		}

		if strings.Contains(strings.ToLower(c.Request().Header.Get("Answer-Type")), "xml") {
			return c.XML(http.StatusOK, answer)
		} else {
			return c.JSON(http.StatusOK, answer)
		}
	})
}
