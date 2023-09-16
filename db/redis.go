package db

import (
	"cas/models"
	"cas/conn"
	"encoding/json"
)

// Redis CRUD for RedisSession

func CreateRedisSession(session *models.RedisSession) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return conn.RedisClient.Set(conn.Ctx, session.Code, data, 0).Err()
}

func GetRedisSessionByCode(code string) (*models.RedisSession, error) {
	data, err := conn.RedisClient.Get(conn.Ctx, code).Result()
	if err != nil {
		return nil, err
	}

	session := &models.RedisSession{}
	err = json.Unmarshal([]byte(data), session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func UpdateRedisSession(session *models.RedisSession) error {
	return CreateRedisSession(session) // 使用相同的键更新数据
}

func DeleteRedisSession(code string) error {
	return conn.RedisClient.Del(conn.Ctx, code).Err()
}
