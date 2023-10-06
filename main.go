package main

import (
	"fmt"
	"os"
)

func main() {
	goPath := os.Getenv("GOPATH")
	connectorPath := os.Getenv("CREDENTIALSERVICE")
	filePath := goPath + connectorPath
	fmt.Println("The path is: ", filePath)
}
