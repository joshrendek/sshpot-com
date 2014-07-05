package main

import "time"

type SshLogin struct {
	Id         int64
	RemoteAddr string
	Username   string
	Password   string
	CreatedAt  time.Time
	DeletedAt  time.Time
}
