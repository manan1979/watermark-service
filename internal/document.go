package internal

type Document struct {
	TicketID  string `json:"id"`
	Content   string `json:"context"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Topic     string `json:"topic"`
	Watermark string `json:"watermark,omitempty"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type Status string

const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "InProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
