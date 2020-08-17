package model

import (
	"time"
)

type Demo struct {
	Id          int
	Name        string
	Description string
	CreateTime  time.Time
}
