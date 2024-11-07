package models

import (
	"message-queue-service/queue"
	"time"
)

type SendMessageInput struct {
	Subject   string
	Body      string
}

// messages
func SendMessage(queueName string, sendMessageInput SendMessageInput) (*Response, error) {
	res, err := queue.SendMessage(queueName, queue.Message{
		Subject: sendMessageInput.Subject,
		Body: sendMessageInput.Body,
		Timestamp: time.Now(),
	})
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: res,
	}, nil
}

func ReceiveMessage(queueName string) (*Response, error) {
	res, err := queue.ReceiveMessage(queueName, 10)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: res,
	}, nil
}

func DeleteMessage(queueName, receiptHandle string) (*Response, error) {
	res, err := queue.DeleteMessage(queueName, receiptHandle)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: res,
	}, nil
}