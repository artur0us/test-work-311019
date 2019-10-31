package mdls

type ReqAnswer struct {
	ServerStatus int `json:"server_status"`
	EntityStatus int `json:"entity_status"`

	Message string `json:"message"`

	Data interface{} `json:"data, omitempty"`
}
