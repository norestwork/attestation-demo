package main

import (
	"demo/server"
	"fmt"
)

func main() {
	server := server.NewHTTPServer()

	if err := server.Run(); err != nil {
		fmt.Println("Server error:", err)
		return
	}
}
