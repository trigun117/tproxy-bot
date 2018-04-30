package main

import (
	"os"
	"os/exec"
)

func main() {

	//Creating command which create self-signed certificates
	cmd := exec.Command(`sh`, `-c`, `openssl req -newkey rsa:2048 -sha256 -nodes -keyout key.pem -x509 -days 365 -out cert.pem -subj "/C=US/ST=New York/L=Brooklyn/O=Example Brooklyn Company/CN=$IP"`)

	//Execute it
	cmd.Run()

	if os.Getenv("CREATE_TABLE") == "yes" {
		createTable()
	}

	start()
}
