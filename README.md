# Kafka to Websocket

## Using a local Kafka within a minikube cluster

Deploy Kafka

```sh
kubectl apply -f ./kafka.yaml
```

Expose kafka broker in order to connect from your local machine for using
this app and also other tools like kafkacat

```sh
kubectl port-forward -nk2w services/kafka 9092
```

## Use kafkacat for consuming and producing messages

To consume messages

```sh
kafkacat -b <kafka_brokers> -G <consumer_group> <...topic>
```

To produce messages

```sh
echo <message> | kafkacat -P -b <kafka_brokers> -t <topic> -p <partition>
```
