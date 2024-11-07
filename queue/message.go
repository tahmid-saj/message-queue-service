package queue

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Message struct {
	Subject string
	Body string
	Timestamp time.Time
}

// Message operations
func SendMessage(queueName string, message Message) (bool, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
    QueueName: &queueName,
	})
	if err != nil {
		log.Println(err)
		return false, err
	}

	queueUrl := result.QueueUrl

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Subject": &sqs.MessageAttributeValue{
				DataType: aws.String("String"),
				StringValue: aws.String(message.Subject),
			},
			"Timestamp": &sqs.MessageAttributeValue{
				DataType: aws.String("String"),
				StringValue: aws.String(message.Timestamp.String()),
			},
		},
		MessageBody: aws.String(message.Body),
		QueueUrl: queueUrl,
	})
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}