package main

import (
    "errors"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

const inputFilePath = "messages.txt"

func main() {
    file, err := os.Open(inputFilePath)
    if err != nil {
        log.Fatalf("could not open %s: %s\n", inputFilePath, err)
    }

    fmt.Printf("Reading data from %s\n", inputFilePath)
    fmt.Println("=====================================")

    linesChan := getLinesChannel(file)

    for line := range linesChan {
        fmt.Printf("read: %s\n", line)
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
