package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sgbcoder/KafkaGo/utils"
)

func ParseDetails(kafkaConfigFilePath string) KafkaDetails {
	replicationFactor := utils.GetReplicationFactor()
	file, _ := ioutil.ReadFile(kafkaConfigFilePath)

	data := KafkaDetails{}

	_ = json.Unmarshal([]byte(file), &data)
	for index := range data.Details {
		data.Details[index].Replication = replicationFactor
	}
	fmt.Println("Parsed Data from kafkaConfig.json: ", data)
	return data
}