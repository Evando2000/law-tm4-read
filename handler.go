package main

import (
	"github.com/gin-gonic/gin"
)

func readHandler(c *gin.Context) {
	npm, err := validateGetMhsHandler(c)

	if err != nil {
		failedResponseConstructor(c, err.Error())
		return
	}

	mhs, errCreate := getMhs(*npm)
	if errCreate != nil {
		failedResponseConstructor(c, errCreate.Error())
		return
	}
	getMhsSuccessResponseConstructor(c, *mhs)
}
