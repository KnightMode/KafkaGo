package cmd

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)

func IncreasePartition(kafkaClient *kafka.AdminClient, kafkaDetails []kafka.PartitionsSpecification) {
	fmt.Println("Running Increase Parition job.....")
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
		fmt.Printf("%s\n", result)
	}
}
