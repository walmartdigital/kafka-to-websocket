#!/usr/bin/env bash

set -e

docker run -d \
  --name=zookeeper \
  -e ZOOKEEPER_CLIENT_PORT=2181 \
  -p 2181:2181 \
  confluentinc/cp-zookeeper:5.1.0

docker run -d \
  --name=kafka \
  -p 9092:9092 \
  -e KAFKA_ZOOKEEPER_CONNECT=192.168.10.6:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.10.6:9092 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  confluentinc/cp-kafka:5.1.0
