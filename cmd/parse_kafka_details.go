package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ParseDetails(kafkaConfigFilePath string) KafkaDetails {
	file, _ := ioutil.ReadFile(kafkaConfigFilePath)

	data := KafkaDetails{}

	_ = json.Unmarshal([]byte(file), &data)

	if !data.IsValid() {
		os.Exit(1)
	}
	return data
}