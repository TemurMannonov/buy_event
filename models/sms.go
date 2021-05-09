package models

// Content - content of the message
type Content struct {
	Text string `json:"text"`
}

// StorageSms - struct to store to database
type StorageSms struct {
	Recipients []string `json:"recipients"`
	Text       string   `json:"text"`
}

// SMS - neccessary fields to send sms
type SMS struct {
	Originator string  `json:"originator"`
	Content    Content `json:"content"`
}

// Message - structure for sms exchange in rabbitMQ
type Message struct {
	SessionID     string `json:"session_id"`
	Recipient     string `json:"recipient"`
	RouteID       string `json:"route_id"`
	RegionCode    string `json:"region_code"`
	MessageID     string `json:"message-id"`
	Type          string `json:"type"`
	ScheduledDate int    `json:"scheduled-date"`
	StoreID       string `json:"store_id"`
	SMS           SMS    `json:"sms"`
}

// Body - Array of messages
type Body struct {
	Messages []Message `json:"messages"`
}
