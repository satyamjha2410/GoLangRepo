package main

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	BootstrapServers string `yaml:"bootstrap_servers"`
	GroupID          string `yaml:"group_id"`
}

func loadConfig(env string) (*Config, error) {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config map[string]Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config[env], nil
}

func main() {
	// Specify the environment (e.g., "development" or "production").
	env := "development"

	// Load configuration based on the specified environment.
	config, err := loadConfig(env)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Create a new Kafka consumer configuration using the loaded settings.
	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"group.id":          config.GroupID,
		"auto.offset.reset": "earliest",
	}

	// Rest of the code remains unchanged...

	// Close the consumer when done.
	consumer.Close()
}
