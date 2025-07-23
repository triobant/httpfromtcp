package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    serverAddr = "localhost:42069"

    udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error resolving UDP address: %v\n", err)
	os.Exit(1)
    }

    conn, err := net.DialUDP("udp", nil, udpAddr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error dialing UDP: %v\n", err)
	os.Exit(1)
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print(">")
	stdin, err := reader.ReadString('\n')
	if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	    os.Exit(1)
	}

	_, err = conn.Write([]byte(stdin))
	if err != nil {
            fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
	    os.Exit(1)
	}
    }
}
