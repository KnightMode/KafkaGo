package cmd

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaDetails struct {
	Brokers []string `json:"brokers"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Topic           string `json:"topic"`
	Partition       int    `json:"partition"`
	Replication     int    `json:"replication"`
	RetentionPeriod string `json:"retention_period"`
}

func (detail Detail) IsValid() bool {
	if detail.Topic == "" ||
		detail.Partition == 0 ||
		detail.Replication == 0 {
		fmt.Println("Missing Required Details.Exiting.....")
		fmt.Println("Topic: ", detail.Topic, " Partition: ", detail.Partition, " Retention Period: ", detail.RetentionPeriod)
		return false
	}
	return true
}
func (details KafkaDetails) IsValid() bool {
	if len(details.Brokers) == 0 || details.Brokers[0] == "" {
		fmt.Println("Broker Details are mandatory.Exiting.....")
		return false
	}
	return true
}

func (details KafkaDetails) ToTopicSpecification() []kafka.TopicSpecification {
	var topicSpecifications []kafka.TopicSpecification
	for _, detail := range details.Details {
		topicSpecifications = append(topicSpecifications, kafka.TopicSpecification{
			Topic:             detail.Topic,
			NumPartitions:     detail.Partition,
			ReplicationFactor: detail.Replication,
			Config: map[string]string{
				"retention.ms": detail.RetentionPeriod,
			},
		})
	}
	return topicSpecifications
}

func (details KafkaDetails) ToPartitionSpecification() []kafka.PartitionsSpecification {
	var topicSpecifications []kafka.PartitionsSpecification
	for _, detail := range details.Details {
		topicSpecifications = append(topicSpecifications, kafka.PartitionsSpecification{
			Topic:      detail.Topic,
			IncreaseTo: detail.Partition,
		})
	}
	return topicSpecifications
}
