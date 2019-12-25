package main

type event struct {
	Data []int `json:"data"`
	Time uint  `json:"time"`
}

// {"events":[{"data":[72,16],"time":1577253372112},{"data":[8,40],"time":1577253372112}],"ppk_num":286,"type":"events"}
type messageEvents struct {
	Type   string `json:"type"`
	PPKNum uint   `json:"ppk_num"`
	Events event  `json:"events"`
}
