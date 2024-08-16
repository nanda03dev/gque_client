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
