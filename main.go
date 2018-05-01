package main

import (
	"os"
)

func main() {

	if os.Getenv("CREATE_TABLE") == "yes" {
		createTable()
	}

	startBot()
}
