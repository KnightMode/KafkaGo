package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func IncreasePartition(kafkaClient *kafka.AdminClient, kafkaDetails []kafka.PartitionsSpecification) {
	fmt.Println("\nIncreasing Partitions....")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur, err := time.ParseDuration("30s")
	if err != nil {
		panic("ParseDuration(30s)")
	}

	results, err := kafkaClient.CreatePartitions(
		ctx,
		kafkaDetails,
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to increase partition: %v\n", err)
		os.Exit(1)
	}

	for _, result := range results {
		if result.Error.Code() == kafka.ErrNoError {
			fmt.Println("Topic name :", result.Topic," Partition count increased successfully")
		} else if strings.HasPrefix(result.Error.String(), "Topic already has") {
			fmt.Println("Topic name :", result.Topic, result.Error)
		} else {
		fmt.Println("Topic name :", result.Topic," Error message:", result.Error)
		}
	}
}
