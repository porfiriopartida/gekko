package main

import (
	"fmt"
	"os"
	"./server"
)

func main() {
	fmt.Println(os.Getenv("GOPATH"))
	server.Start()
}
