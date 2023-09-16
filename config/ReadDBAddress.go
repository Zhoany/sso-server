package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"cas/models"
)





func LoadMysqlConfig() (*models.MysqlConfig, error) {
	data, err := ioutil.ReadFile("./conf/mysql.yaml")
	
	if err != nil {
		log.Fatalf("Error reading MySQL config: %v", err)
		return nil, err
	}

	var config models.MysqlConfig
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Fatalf("Error unmarshalling MySQL config: %v", err)
		return nil, err
	}

	return &config, nil
}

func LoadRedisConfig() (*models.RedisConfig, error) {
	data, err := ioutil.ReadFile("./conf/redis.yaml")
	if err != nil {
		log.Fatalf("Error reading Redis config: %v", err)
		return nil, err
	}

	var config models.RedisConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling Redis config: %v", err)
		return nil, err
	}

	return &config, nil
}
