package msg

type ClientMessage struct {
	Event     string `json:"event"`
	EventType string `json:"event_type"`
	Data      string `json:"data"`
}
