package userhandlers

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func HandleUserCreate(msg *nats.Msg) {
	fmt.Printf("Received message from create: %s\n", string(msg.Data))
}
func HandleUserDelete(msg *nats.Msg) {
	fmt.Printf("Received message from delete: %s\n", string(msg.Data))
}
