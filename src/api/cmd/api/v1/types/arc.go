package types
type Arc struct {
	ID           int         `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Episodes     interface{} `json:"episodes"`
	Image        string      `json:"image"`
	Time         interface{} `json:"time"`
	CreatedDate  string      `json:"created_date"`
    FinishedDate string      `json:"finished_date"`
	Rate         interface{} `json:"rate"`
}