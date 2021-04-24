package e

type Error struct {
	Code     int    `json:"code"`
	DevInfo  string `json:"dev_info"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
}

type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type FaildResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
}

type SuccessResponse struct {
	*Base
	Data interface{} `json:"data"`
}
