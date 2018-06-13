package main

import (
	"os"
	"time"
)

func main() {

	time.Sleep(60 * time.Second)

	if os.Getenv("TABLE") == "create" {
		createTable()
	}

	startBot()
}
