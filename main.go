package main

import (
	"money_transfer_service/cmd"
)

func main() {
	cmd.Start(cmd.NewByEnvironment())
}