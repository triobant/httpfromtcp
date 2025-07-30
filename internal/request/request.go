package request

import (
    "bytes"
    "fmt"
    "io"
    "strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
    b, err := io.ReadAll(reader)
    if err != nil {
        return nil, err
    }

    requestLine, err := parseRequestLine(string(b))
    if err != nil {
        return nil, err
    }

    var req Request
    req.RequestLine = *requestLine

    return &req, nil
}

func parseRequestLine(b string) (*RequestLine, error) {
    requestLineSlice := strings.Split(b, "\r\n")
    requestSlice := strings.Split(requestLineSlice[0], " ") 
    if len(requestSlice) != 3 {
        return nil, errors.New("request-line doesn't contain three elements")
    }

    var rl RequestLine
    rl.Method = requestSlice[0]
    rl.RequestTarget = requestSlice[1]
    rl.HttpVersion = strings.Split(requestSlice[2], "/")[1]

    return &rl, nil
}
