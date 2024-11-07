package routes

import (
	"message-queue-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// messages
func sendMessage(context *gin.Context) {
	queueName := context.Param("queueName")

	var sendMessageInput models.SendMessageInput

	err := context.ShouldBindJSON(&sendMessageInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse message input"})
		return
	}

	res, err := models.SendMessage(queueName, sendMessageInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not send message"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func receiveMessage(context *gin.Context) {
	queueName := context.Param("queueName")

	res, err := models.ReceiveMessage(queueName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not receive message"})
		return
	}

	context.JSON(http.StatusOK, res)
}

func DeleteMessage(context *gin.Context) {
	queueName := context.Param("queueName")
	receiptHandle := context.Param("receiptHandle")

	res, err := models.DeleteMessage(queueName, receiptHandle)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete message"})
		return
	}

	context.JSON(http.StatusOK, res)
}