package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
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

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(">")
	stdin, err := reader.ReadString('\n')
	if err != nil {
	    log.Fatalf("error reading input string: %s", err.Error())
	}

	_, err = conn.Write([]byte(stdin))
	if err != nil {
	    log.Fatalf("error writing input to udp connection: %s", err.Error())
	}
    }
}
