echo "Hello, Kafka!" | docker exec -i wbtech-kafka-1 kafka-console-producer.sh --broker-list localhost:9092 --topic orders
