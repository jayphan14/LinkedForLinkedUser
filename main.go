package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"linkedforlinked.com/user/userhandlers"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Printf("Error connecting to NATS server: %v\n", err)
		return
	}
	defer nc.Close()

	subjectHandlers := map[string]func(msg *nats.Msg){
		"user.create": userhandlers.HandleUserCreate,
		"user.delete": userhandlers.HandleUserDelete,
	}

	for subject, handler := range subjectHandlers {
		sub, err := nc.Subscribe(subject, handler)
		if err != nil {
			fmt.Printf("Error subscribing to subject: %v\n", err)
			return
		}
		defer sub.Unsubscribe()
	}

	fmt.Println("User service started successfully")

	// Run indefinitely (or replace with your service logic)
	select {}
}
