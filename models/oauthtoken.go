	package models

	import "time"



	type AccessToken struct {
		Token   string    `xorm:"token VARCHAR(750) pk"`
		Sub    string      `xorm:"user_id VARCHAR(255)"`
	    Aud  string    `xorm:"client_id VARCHAR(255)"`
		Nonce     string    `xorm:"nonce VARCHAR(255)"`
		Exp time.Time `xorm:"expires_at DATETIME"`
		Iat time.Time `xorm:"Issued_at DATETIME"`
		Scope string `xorm:"scope VARCHAR(255)"`
	}

	
