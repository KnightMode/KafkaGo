package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaDetails struct {
	Brokers []string  `json:"brokers"`
	Details []Details `json:"details"`
}
type Details struct {
	Topic           string `json:"topic"`
	Partition       int    `json:"partition"`
	Replication     int    `json:"replication"`
	RetentionPeriod string `json:"retention_period"`
}

func main() {
	file, _ := ioutil.ReadFile("../../details/kafka_details.json")

	data := KafkaDetails{}

	_ = json.Unmarshal([]byte(file), &data)

	if len(data.Brokers) == 0 || data.Brokers[0] == "" {
		fmt.Println("Broker Details are mandatory.Exiting.....")
		os.Exit(1)
	}

	for i := 0; i < len(data.Details); i++ {
		var detail = data.Details[i]
		if detail.Topic == "" ||
			detail.Partition == 0 ||
			detail.Replication == 0 {
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

		maxDur, err := time.ParseDuration("30s")
		if err != nil {
			panic("ParseDuration(30s)")
		}

		results, err := a.CreateTopics(
			ctx,
			[]kafka.TopicSpecification{{
				Topic:             detail.Topic,
				NumPartitions:     detail.Partition,
				ReplicationFactor: detail.Replication,
				Config: map[string]string{
					"retention.ms": detail.RetentionPeriod,
				},
			}},
			kafka.SetAdminOperationTimeout(maxDur))
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
