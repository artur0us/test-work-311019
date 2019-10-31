package constants

import (
	"../mdls"
)

func TemplateReqAnswer() mdls.ReqAnswer {
	return mdls.ReqAnswer{
		ServerStatus: 1,
		EntityStatus: 1,

		Message: "OK",

		Data: nil,
	}
}
