package entity

import "time"

type Asset struct {
	Name      string
	Uid       int64
	Data      []byte
	CreatedAt time.Time
}
