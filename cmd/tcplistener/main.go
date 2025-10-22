package main

import (
    "fmt"
    "log"
    "net"
    "httpfromtcp/internal/request"
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

	req, err := request.RequestFromReader(conn)
	if err != nil {
	    log.Fatalf("error parsing request: %s\n", err.Error())
	}
	fmt.Println("Request line:")
	fmt.Printf("- Method: %s\n", req.RequestLine.Method)
	fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
	fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)
    }
}
