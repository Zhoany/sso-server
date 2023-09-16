package conn

import (
	"fmt"
	"context"
	"github.com/go-xorm/xorm"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"cas/models"
	"cas/config"
)

var Engine *xorm.Engine
var RedisClient *redis.Client
var Ctx = context.Background()

func InitDB() error {
	mysqlConfig, err := config.LoadMysqlConfig()
	if err != nil {
		return err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to create engine: %w", err)
	}

	// Check if the database exists
	hasDB, err := engine.SQL(fmt.Sprintf("SHOW DATABASES LIKE '%s'", mysqlConfig.Database)).QueryString()
	if err != nil || len(hasDB) == 0 {
		_, err = engine.Exec(fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci", mysqlConfig.Database))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
	}

	// Close the initial engine and connect to the new database
	engine.Close()

	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to create engine: %w", err)
	}

	err = Engine.Sync2(new(models.User), new(models.AccessToken),  new(models.Cookie), new(models.Client))
	if err != nil {
		return fmt.Errorf("failed to sync database tables: %w", err)
	}

	return nil
}

func InitRedis() error {
	redisConfig, err := config.LoadRedisConfig()
	if err != nil {
		return err
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})

	_, err = RedisClient.Ping(Ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}
	return nil
}
