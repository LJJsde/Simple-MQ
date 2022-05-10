package main

import (
	"time"
)

type MQ interface {
	Send(message interface{})
	Res(size int, timeout time.Duration) []interface{}
	Size() int
	Capacity() int
}

type MyMQ struct {
	Message chan interface{}
	cap     int
}

func (receiver *MyMQ) Send(message interface{}) {
	select {
	case receiver.Message <- message:
	default:

	}
}

func (receiver *MyMQ) Res(size int, timeout time.Duration) []interface{} {
	msg := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		select {
		case res := <-receiver.Message:
			msg = append(msg, res)
		case <-time.After(timeout):
			return msg
		}
	}
	return msg
}

func (receiver *MyMQ) Size() int {
	return len(receiver.Message)
}

func (i *MyMQ) Capacity() int {
	return i.cap
}

func CreateMQ(capacity int) MQ {
	mq := &MyMQ{
		Message: make(chan interface{}, capacity),
		cap:     capacity,
	}
	return mq
}

func main() {

}
