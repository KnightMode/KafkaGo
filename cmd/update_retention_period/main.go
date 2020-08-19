package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"os"
)

type KafkaDetails struct {
	Brokers []string  `json:"brokers"`
	Details []Details `json:"details"`
}
type Details struct {
	Topic     string `json:"topic"`
	Retention string `json:"retention"`
}

func main() {
	file, _ := ioutil.ReadFile("../../details/update_retention_period.json")

	data := KafkaDetails{}

	_ = json.Unmarshal([]byte(file), &data)

	if len(data.Brokers) == 0 || data.Brokers[0] == "" {
		fmt.Println("Broker Details are mandatory.Exiting.....")
		os.Exit(1)
	}

	for i := 0; i < len(data.Details); i++ {
		var detail = data.Details[i]
		if detail.Topic == "" ||
			detail.Retention == "" {
			fmt.Println("Missing Required Details.Exiting.....")
			os.Exit(1)
		}

		a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": data.Brokers[0]})
		if err != nil {
			fmt.Printf("Failed to create Admin client: %s\n", err)
			os.Exit(1)
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		results, err := a.AlterConfigs(
			ctx,
			[]kafka.ConfigResource{{
				Type: kafka.ResourceTopic,
				Name: detail.Topic,
				Config: []kafka.ConfigEntry{
					{
						Name:  "retention.ms",
						Value: detail.Retention,
					},
				},
			}})
		if err != nil {
			fmt.Printf("Failed to create topic: %v\n", err)
			os.Exit(1)
		}

		for _, result := range results {
			fmt.Printf("%s\n", result)
		}

		a.Close()
	}
}
