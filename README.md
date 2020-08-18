# KafkaGo

  Uses the confluent-go library to abstract the topic creation and partition increase operations.


#Prerequisities

  Kafka Broker -> Running on port 9092
  Zookeeper -> Running on port 2181
  
  
  To Run:
  
  For creating topic(Update the topic, partition, retention period details in the json in details directory prior to running):
   
   
   cd cmd
   
   
   cd create_topic
   
   
   go run .
