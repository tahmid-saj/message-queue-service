package routes

import (
	"message-queue-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// queue
func listQueues(context *gin.Context) {
	res, err := models.ListQueues()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not list queues"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func getQueueURL(context *gin.Context) {
	queueName := context.Param("queueName")

	res, err := models.GetQueueURL(queueName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get queue URL"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func createQueue(context *gin.Context) {
	var createQueueInput models.CreateQueueInput

	err := context.ShouldBindJSON(&createQueueInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse queue creation input"})
		return
	}

	res, err := models.CreateQueue(createQueueInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create queue"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func configureVisibilityTimeout(context *gin.Context) {
	var configureVisibilityTimeoutInput models.ConfigureVisibilityTimeoutInput

	err := context.ShouldBindJSON(&configureVisibilityTimeoutInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse queue configuration input"})
		return
	}

	res, err := models.ConfigureVisibilityTimeout(configureVisibilityTimeoutInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not configure queue visibility timeout"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func deleteQueue(context *gin.Context) {
	queueName := context.Param("queueName")

	res, err := models.DeleteQueue(queueName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete queue"})
		return
	}

	context.JSON(http.StatusOK, res)
}