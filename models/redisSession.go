package models

type RedisSession struct {
	Code   string `json:"code"`
    Nonce string    `json:"nonce"`
}