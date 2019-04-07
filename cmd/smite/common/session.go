package common

type Session struct {
	Message   string `json:"ret_msg"`
	Id        string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}
