package main

type User struct {
	Username string `json:"username"`
	UserId   int64  `json:"user_id"`
	ChatId   int64  `json:"chat_id"`
}
