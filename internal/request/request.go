package request

import (
    "errors"
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
    requestSlice := strings.Split(b, " ") 
    if len(requestSlice) != 3 {
        return nil, errors.New("request-line doesn't contain three elements")
    }

    RequestLine.Method = requestSlice[0]
    RequestLine.RequestTarget = requestSlice[1]
    RequestLine.HttpVersion = requestSlice[2]

    return RequestLine, nil
}
