package models

type User struct {
	Sub          string `xorm:"varchar(255)"`
	ID           int    `xorm:"int autoincr pk"`
	Username     string `xorm:"varchar(255)"`
	Password string `xorm:"varchar(255)"`
	Email        string `xorm:"varchar(255)"`
}
