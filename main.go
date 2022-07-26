package main

import (
	"log"

	"go-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("err：%v", err)
	}
}