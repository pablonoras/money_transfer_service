package main

import (
	"github.com/pablonoras/money_transfer_service/cmd"
)

func main() {
	cmd.Start(cmd.NewByEnvironment())
}