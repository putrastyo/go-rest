package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type validationErrorMsg struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func HandleValidationError(c *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errMsg := make(map[string][]string)

		for _, err := range validationErrors {
			field := err.Field()
			tag := err.Tag()

			errMsg[field] = append(errMsg[field], tag)
		}

		errResponse := validationErrorMsg{
			Status:  "failed",
			Message: "Validation fail",
			Errors:  errMsg,
		}

		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "failed",
		"message": "internal server error",
	})
}
