package models


type Client struct {
	ID          string `xorm:"varchar(255) notnull pk 'client_id'" json:"id"`
	Secret      string `xorm:"varchar(255) notnull 'client_secret'" json:"secret"`
	RedirectURI string `xorm:"varchar(255) 'redirect_uri'" json:"redirect_uri"`
	Client_name string `xorm:"varchar(255) 'client_name'" json:"name"`
}
