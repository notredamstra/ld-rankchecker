package ldmdare

import "net/http"

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{BaseURL: baseURL, httpClient: httpClient}
}

func NewLDGame(id int, name string, description string, grade float64, cover string, eventId int, eType string) *LDGame {
	return &LDGame{Id: id, Name: name, Description: description, Grade: grade, Cover: cover, EventId: eventId, Type: eType}
}

