package main

import "time"

type SshLogin struct {
	Id         int64     `json:"id"`
	RemoteAddr string    `json:"remote_addr"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"-"`
}

type ApiStat struct {
	Id      int64
	Counter int64
	DateKey string
}

type ApiResponse struct {
	page  int64      `json:"page"`
	total int64      `json:"total"`
	data  []SshLogin `json:"data"`
}
