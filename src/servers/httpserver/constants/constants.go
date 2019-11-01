package constants

import (
	"github.com/artur0us/test-work-311019/servers/httpserver/mdls"
)

func TemplateReqAnswer() mdls.ReqAnswer {
	return mdls.ReqAnswer{
		ServerStatus: 1,
		EntityStatus: 1,

		Message: "OK",

		Data: nil,
	}
}
