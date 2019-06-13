package uiresource

type APIResponse struct {
	Header Header      `json:"header"'`
	Data   interface{} `json:"data"`
}
type Header struct {
	Status int `json:"status"`
	Reason string `json:"reason"`
}
