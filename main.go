package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Redis struct {
	Host string `json:"host"`
}

type API struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type RabbitMQ struct {
	Host     string `json:"host"`
	Vhost    string `json:"vhost"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type SensuConfig struct {
	RabbitMQ RabbitMQ `json:"rabbitmq"`
	Redis    Redis    `json:"redis"`
	API      API      `json:"api"`
}

func main() {
	apiPort := os.Getenv("SENSU_API_PORT") // getenv is string, we need an int
	port, err := strconv.Atoi(apiPort)     // convert
	if err != nil {
		log.Fatal(err)
	}
	rabbitMQ := &RabbitMQ{
		os.Getenv("SENSU_RABBITMQ_HOST"),
		os.Getenv("SENSU_RABBITMQ_VHOST"),
		os.Getenv("SENSU_RABBITMQ_USER"),
		os.Getenv("SENSU_RABBITMQ_PASSWORD"),
	}
	redis := &Redis{
		os.Getenv("SENSU_REDIS_HOST"),
	}
	api := &API{Port: port,
		Host:     os.Getenv("SENSU_API_HOST"),
		User:     os.Getenv("SENSU_API_USER"),
		Password: os.Getenv("SENSU_API_PASSWORD"),
	}
	sensuConfig := &SensuConfig{*rabbitMQ, *redis, *api}
	f, err := os.Create("/etc/sensu/conf.d/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	jsn, err := json.MarshalIndent(sensuConfig, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(jsn)
	if err != nil {
		log.Fatal(err)
	}
}
