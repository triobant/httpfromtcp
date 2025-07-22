package main

import (
    "fmt"
    "net"
    "log"
)

const addr = "localhost:42069"

func main() {
    resolv, err := net.ResolveUDPAddr("udp", addr)
    if err != nil {
        log.Fatalf("error resolving address: %s\n", err.Error())
    }
    conn, err := net.DialUDP("udp", nil, resolv)
    if err != nil {
        log.Fatalf("error opening udp connection: %s\n", err.Error())
    }
    defer conn.Close()
}
