package cmd

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"time"
)

func CreateTopic(kafkaClient *kafka.AdminClient, kafkaDetails []kafka.TopicSpecification) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeoutPeriod, err := time.ParseDuration("30s")
	if err != nil {
		panic("ParseDuration(30s)")
	}
	fmt.Println("Creating topics....")
	results, kafkaTopicError := kafkaClient.CreateTopics(
		ctx,
		kafkaDetails,
		kafka.SetAdminOperationTimeout(timeoutPeriod),
	)
	if kafkaTopicError != nil {
		fmt.Printf("Failed to create topic: %v\n", kafkaTopicError.Error())
		os.Exit(1)
	}

	for _, result := range results {
		if result.Error.Code() == kafka.ErrNoError {
			fmt.Println("Topic name: ", result.Topic, " created successfully")
		} else if result.Error.Code() == kafka.ErrTopicAlreadyExists {
			fmt.Println("Topic name: ", result.Topic, " already exists.")
		} else {
			fmt.Printf("Topic name : %s, Error message: %v \n", result.Topic, result.Error)
		}
	}
}
