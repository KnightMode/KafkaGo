package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseBrokerDetails() string {
	brokerDetailsFromEnv := os.Getenv("BROKER_DETAILS")
	brokerDetails := strings.Split(brokerDetailsFromEnv, ",")
	if len(brokerDetails) == 0 || brokerDetails[0] == "" {
		fmt.Println("Broker Details are mandatory.Exiting.....")
		os.Exit(1)
	}
	return brokerDetails[0]
}

func GetReplicationFactor() int {
	replicationFactorFromEnv := os.Getenv("REPLICATION_FACTOR")
	replicationFactor, err := strconv.Atoi(replicationFactorFromEnv)
	if err != nil || replicationFactor == 0 {
		fmt.Println("Invalid Replication factor.Exiting.....")
		os.Exit(1)
	}
	return replicationFactor
}
