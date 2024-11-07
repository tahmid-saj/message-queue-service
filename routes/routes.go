package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// queue operations
	// ListQueues
	server.GET("/queues", listQueues)
	// GetQueueURL
	server.GET("/queues/:queueName", getQueueURL)
	// CreateQueue
	server.POST("/queues", createQueue)
	// ConfigureVisibilityTimeout
	server.PUT("/queues", configureVisibilityTimeout)
	// DeleteQueue
	server.DELETE("/queues/:queueName", deleteQueue)

	// message operations
	// ReceiveMessage
	server.GET("/queues/:queueName/messages", receiveMessage)
	// SendMessage
	server.POST("/queues/:queueName/messages", sendMessage)
	// DeleteMessage
	server.DELETE("/queues/:queueName/messages/:receiptHandle", deleteMessage)
}