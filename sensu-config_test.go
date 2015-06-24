package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestAPIStructMarshall(t *testing.T) {
	testjsn := []byte(`{
	"rabbitmq": {
		"host": "sensu-rabbit.rabbitmq.cluster.local",
		"vhost": "",
		"user": "",
		"password": ""
	},
	"redis": {
		"host": "sensu-redis.redis.cluster.local"
	},
	"api": {
		"host": "sensu-api.sensu.cluster.local",
		"port": 4567
	}
}`)
	api := &API{
		Port: 4567,
		Host: "sensu-api.sensu.cluster.local",
	}
	redis := &Redis{Host: "sensu-redis.redis.cluster.local"}
	rabbitMQ := &RabbitMQ{
		"sensu-rabbit.rabbitmq.cluster.local",
		os.Getenv("SENSU_RABBITMQ_VHOST"),
		os.Getenv("SENSU_RABBITMQ_USER"),
		os.Getenv("SENSU_RABBITMQ_PASSWORD"),
	}
	sensuconfig := &SensuConfig{API: api, Redis: redis, RabbitMQ: rabbitMQ}
	jsn, err := json.MarshalIndent(sensuconfig, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	if string(jsn) != string(testjsn) {
		fmt.Println(string(jsn))
		t.Error("config does not match")
	}

}

func TestNewConfig(t *testing.T) {
	os.Setenv("SENSU_API_PORT", "4567")
	os.Setenv("SENSU_RABBITMQ_HOST", "rabbitmq")
	os.Setenv("SENSU_RABBITMQ_VHOST", "/")
	os.Setenv("SENSU_RABBITMQ_USER", "guest")
	os.Setenv("SENSU_RABBITMQ_PASSWORD", "guest")
	os.Setenv("SENSU_REDIS_HOST", "redis")
	config, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	if config.API.Port != 4567 {
		t.Error("API Port not set")
	}

	err = config.Write("/tmp/sensuconfig.json")
	if err != nil {
		t.Fatal(err)
	}
}
