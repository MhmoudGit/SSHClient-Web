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

// func getOs(connection *ssh.Client) string {
// 	session, err := connection.NewSession()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer session.Close()

// 	// Run a command to get information about the OS
// 	cmd := "sudo systemctl status alameda"
// 	output, err := session.CombinedOutput(cmd)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Extract the OS information from the command output
// 	value := fmt.Sprintf("Connected to %v successfully. Remote OS: %v", connection.RemoteAddr(), string(output))

// 	return value
// }

// func Logger[T any, U any](fn func(arg T) U, arg T) {
// 	startTime := time.Now()
// 	result := fn(arg)
// 	endTime := time.Now()

// 	duration := endTime.Sub(startTime)

// 	log.Printf("%v Execution time: %v\n", result, duration)
// }
