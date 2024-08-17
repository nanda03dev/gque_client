# Gque Client - Golang Client for Gque

Gque Client is a Golang client library designed to interact with the Gque message queue protocol. This client provides an easy-to-use interface for connecting to Gque, creating broadcast channels and queues, pushing messages, broadcasting messages, and consuming messages.

## Features

- **Connect**: Establishes a connection to the Gque server.
- **CreateBroadcast**: Creates a broadcast channel in Gque.
- **CreateQueue**: Creates a queue for message processing.
- **PushMessage**: Pushes a message into a specified queue.
- **BroadcastMessage**: Sends a message to all consumers subscribed to a broadcast channel.
- **Consume**: Consumes messages from a queue or broadcast channel.

## Installation

To install `gque_client`, run:

```bash
go get github.com/nanda03dev/gque_client
```

## Usage

### 1. Create a Queue

Define a queue and create it using the Gque client.

```go
// Define the queue with a name and a time-to-live (TTL)
var Queue1 = Queue{
	Name: "queue1",
	Time: 600, // Time-to-live in seconds
}

// Create the queue
GqueClient.CreateQueue(Queue1)
```

### 2. Create a Broadcast

Define a broadcast and associate it with the queue.

```go
// Define the broadcast with a name and associated queue names
var Broadcast1 = Broadcast{
	Name:       "all-queue",
	QueueNames: []string{Queue1.Name},
}

// Create the broadcast
GqueClient.CreateBroadcast(Broadcast1)
```

### 3. Consume Messages from a Queue

Set up a consumer to receive messages from the queue.

```go
// Create a channel to receive messages
receiveChan := make(chan MessageType, 10000)

// Define the consumer request with the queue name
consumer1Request := ConsumerRequestType{
	QueueName: Queue1.Name,
}

// Start consuming messages
GqueClient.Consume(consumer1Request, receiveChan)

// Process messages from the channel
for {
	msg, ok := <-receiveChan
	if !ok {
		break
	}
	fmt.Printf("\nQueueName: %v ---- Received Message: %v", Queue1.Name, msg)
}
```

### 4. Push a Message to a Queue

Send a message to the specified queue.

```go
// Define the message to be pushed to the queue
pushMessage := QueueMessageType{
	Name: "queue1",
	Data: MessageType{"MessageFrom": "Queue", "index": 1},
}

// Push the message to the queue
GqueClient.PushMessage(pushMessage)
```

### 5. Broadcast a Message

Send a message to all queues associated with the broadcast.

```go
// Define the broadcast message
broadcastMessage := BroadcastMessageType{
	Name: Broadcast1.Name,
	Data: MessageType{"MessageFrom": "Broadcast", "index": 1},
}

// Broadcast the message
GqueClient.BroadcastMessage(broadcastMessage)
```
