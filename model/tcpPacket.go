package model


type MessagePacket struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	CreateTime float64 `json:"createTime"`
}