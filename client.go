package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	//dial to the server
	serverURL := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	usernameFlag := false
	if err != nil {
		log.Fatalf("Error connecting to Websocket server: %v", err)
	} else {
		fmt.Println("Enter your username: ")
	}
	defer conn.Close()
	// we will then write and read messages
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message: ", err)
				break
			}
			// Print the message  from the server
			log.Println("Message from server: ", string(msg))
		}
	}()

	for {
		var input string
		if usernameFlag {
			fmt.Println("Use `/msg <receiver> <message>` to send privately or just type a message to send to everyone: ")
		} else {
			usernameFlag = true
		}
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Println("Error reading input: ", err)
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, []byte(input))
		if err != nil {
			log.Println("Error sending message: ", err)
			break
		}
	}
}
