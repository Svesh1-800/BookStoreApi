package utils

import (
	"encoding/json"
	"log"
)

type Message struct {
	Msg string
}

func JsonMessageByte(msg string) []byte {
	errMessage := Message{msg}
	byteContent, _ := json.Marshal(errMessage)
	return byteContent
}

func CheckError(err error) {
	if err != nil {
		log.Printf("Error - %v", err)
	}

}
