package models

import (
	"time"
)

type Cookie struct {
	CookieID     string    `xorm:"cookie_id VARCHAR(255) pk"`
	UserID       int       `xorm:"user_id INT"`
	SessionID    string    `xorm:"session_id VARCHAR(255)"`
	CreationTime time.Time `xorm:"creation_time TIMESTAMP"`
	ExpiryTime   time.Time `xorm:"expiry_time TIMESTAMP"`
	Data         string    `xorm:"data TEXT"`
}