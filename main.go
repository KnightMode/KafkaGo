package main

import (
	"fmt"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sgbcoder/KafkaGo/cmd"
	"github.com/sgbcoder/KafkaGo/utils"
)
func main() {
	const kafkaBroker = "bootstrap.servers"
	const kafkaConfigFilePath = "./details/kafka_details.json"
	fmt.Println("Running kafka actions job....")
	brokerDetail := utils.ParseBrokerDetails()
	kafkaDetails := cmd.ParseDetails(kafkaConfigFilePath)
	kafkaAdminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{kafkaBroker: brokerDetail})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	cmd.CreateTopic(kafkaAdminClient, kafkaDetails.ToTopicSpecification())
	cmd.IncreasePartition(kafkaAdminClient, kafkaDetails.ToPartitionSpecification())
	kafkaAdminClient.Close()
}
