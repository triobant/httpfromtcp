package main

import (
    "errors"
    "fmt"
    "io"
    "log"
    "net"
    "strings"
    "github.com/triobant/httpfromtcp/internal/request"
)

const port = ":42069"

func main() {
    listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
    }
    defer listener.Close()

    fmt.Println("Listening TCP traffic on", port)
    for {
        conn, err := listener.Accept()
	if err != nil {
	    log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Println("Accepted connection from", conn.RemoteAddr())
    }
}
