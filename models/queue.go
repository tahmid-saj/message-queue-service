package models

import "message-queue-service/queue"

type CreateQueueInput struct {
	QueueName string `json:"queueName"`
}

type ConfigureVisibilityTimeoutInput struct {
	QueueName          string `json:"queueName"`
	ReceiptHandle      string `json:"receiptHandle"`
	VisibilityDuration int    `json:"visibilityDuration"`
}

// queue
func ListQueues() (*Response, error) {
	res, err := queue.ListQueues()
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

func GetQueueURL(queueName string) (*Response, error) {
	res, err := queue.GetQueueURL(queueName)
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

func CreateQueue(createQueueInput CreateQueueInput) (*Response, error) {
	res, err := queue.CreateQueue(createQueueInput.QueueName)
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

func ConfigureVisibilityTimeout(configureVisibilityTimeoutInput ConfigureVisibilityTimeoutInput) (*Response, error) {
	res, err := queue.ConfigureVisibilityTimeout(
		configureVisibilityTimeoutInput.QueueName,
		configureVisibilityTimeoutInput.ReceiptHandle,
		configureVisibilityTimeoutInput.VisibilityDuration,
	)
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

func DeleteQueue(queueName string) (*Response, error) {
	res, err := queue.DeleteQueue(queueName)
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