package venbest

import "time"

type EventCode uint

type PPKEvent struct {
	Code       EventCode
	Additional uint
	When       time.Time
}

type PPKState struct {
	JSON
}

type State struct {
	PPKs []PPKState
	When time.Time
}

const (
	// Взятие группы под охрану
	EventCode64 EventCode = 64
	// Снятие группы с охраны
	EventCode72 EventCode = 72
	// Открыта дверца ППК
	EventCode108 EventCode = 108
	// Закрыта дверца ППК
	EventCode109 EventCode = 109
)

type JSON map[string]interface{}

type event struct {
	Data []int `json:"data"`
	Time uint  `json:"time"`
}

type messageEvents struct {
	Type   string  `json:"type"`
	PPKNum uint    `json:"ppk_num"`
	Events []event `json:"events"`
}

type PPK struct {
	PPKNum     uint   `json:"ppk_num"`
	Pwd        string `json:"pwd"`
	LicenceKey []uint `json:"license_key"`
}

type loginData struct {
	Type         string `json:"type"`
	UserName     string `json:"user_name"`
	PingInterval uint   `json:"ping_interval"`
	PPKs         []PPK  `json:"ppks"`
}
