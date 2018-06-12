package main

import (
	"os"
)

func main() {

	if os.Getenv("TABLE") == "yes" {
		createTable()
	}

	startBot()
}
