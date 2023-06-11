package helper

import (
	"log"
	"os"
)

func InitLog() *os.File {
	f, err := os.OpenFile("cmd/product-http/product-http-log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// log.SetOutput(f)
	return f
}
