package gque_client

type Queue struct {
	Name string `json:"name"`
	Time int64  `json:"time"`
}

type Broadcast struct {
	Name       string   `json:"name"`
	QueueNames []string `json:"queueNames"`
}

type MessageType map[string]interface{}

type QueueMessageType struct {
	Name string      `json:"name"`
	Data MessageType `json:"data"`
}

type BroadcastMessageType struct {
	Name string      `json:"name"`
	Data MessageType `json:"data"`
}

type ConsumerRequestType struct {
	QueueName string `json:"queueName"`
}
