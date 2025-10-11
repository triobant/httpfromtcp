package main

import (
    "errors"
    "fmt"
    "io"
    "log"
    "net"
    "strings"
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
	    log.Fatalf("failed to accept connection: %s\n", err)
	    continue
	}

	fmt.Printf("Reading data from %s\n", conn)
	fmt.Println("=====================================")

//	linesChan := getLinesChannel(conn)
	linesChan := RequestFromReader(conn)

	for line := range linesChan {
	    fmt.Printf("read: %s\n", line)
	}
    }
}

func getLinesChannel(f io.ReadCloser) <-chan string {
    lines := make(chan string)
    go func() {
        defer f.Close()
        defer close(lines)
        currentLineContents := ""
        for {
            buffer := make([]byte, 8, 8)
            n, err := f.Read(buffer)
            if err != nil {
                if currentLineContents != "" {
                    lines <- currentLineContents
                }
                if errors.Is(err, io.EOF) {
                    break
                }
                fmt.Printf("error: %s\n", err.Error())
                break
            }
            str := string(buffer[:n])
            parts := strings.Split(str, "\n")
            for i := 0; i < len(parts) - 1; i++ {
                lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
                currentLineContents = ""
            }
            currentLineContents += parts[len(parts) - 1]
        }
    }()

    return lines
}
