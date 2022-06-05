package main

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

func validateGetMhsHandler(c *gin.Context) (*string, error) {
	npm := c.Param("npm")
	if npm == "" {
		err := errors.New("npm cannot be empty")
		return nil, err
	}
	return &npm, nil
}

func getMhs(npm string) (*Mahasiswa, error) {
	mhs, err := mhsDB.Get(npm).Result()
	if err != nil {
		return nil, err
	}

	mhsInfo := Mahasiswa{}
	err = json.Unmarshal([]byte(mhs), &mhsInfo)
	if err != nil {
		return nil, err
	}
	return &mhsInfo, nil
}
