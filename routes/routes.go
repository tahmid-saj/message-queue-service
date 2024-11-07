package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// queue operations
	// ListQueues
	server.GET("/queues")
	// GetQueueURL
	server.GET("/queues/:queueName")
	// CreateQueue
	server.POST("/queues")
	// ConfigureVisibilityTimeout
	server.PUT("/queues")
	// DeleteQueue
	server.DELETE("/queues/:queueName")

	// message operations
	// ReceiveMessage
	server.GET("/queues/:queueName/messages")
	// SendMessage
	server.POST("/queues/:queueName/messages")
	// DeleteMessage
	server.DELETE("/queues/:queueName/messages/:receiptHandle")
}