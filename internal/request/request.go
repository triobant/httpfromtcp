package request

import (
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

    requestLine, err := parseRequestLine(b)
    if err != nil {
        return nil, err
    }

    return requestLine, nil
}

func parseRequestLine(b string) (*RequestLine, error) {
}
