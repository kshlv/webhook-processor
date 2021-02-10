package main

// Update ...
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

// Message ...
type Message struct {
	Date      int    `json:"date"`
	Chat      *Chat  `json:"chat"`
	MessageID int    `json:"message_id"`
	From      *User  `json:"from"`
	Text      string `json:"text"`
}

// Chat ...
type Chat struct {
	LastName  string `json:"last_name"`
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

// User ...
type User struct {
	LastName  string `json:"last_name"`
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}
