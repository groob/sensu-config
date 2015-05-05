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
	Port int `json:"port"`
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
	api := &API{port}
	sensuConfig := &SensuConfig{*rabbitMQ, *redis, *api}
	f, err := os.Create("config.json")
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
