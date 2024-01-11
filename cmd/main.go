package main

import (
	// "fmt"
	// "log"
	"sshClient/api"
	// "os"
	// "time"
	// "github.com/joho/godotenv"
	// "golang.org/x/crypto/ssh"
)

func main() {
	api.App()
}

// func Logger[T any, U any](fn func(arg T) U, arg T) {
// 	startTime := time.Now()
// 	result := fn(arg)
// 	endTime := time.Now()

// 	duration := endTime.Sub(startTime)

// 	log.Printf("%v Execution time: %v\n", result, duration)
// }
