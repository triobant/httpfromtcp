package request

import (
    "bytes"
    "fmt"
    "io"
    "strings"
)

type Request struct {
	RequestLine RequestLine
	State	      int
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

const (
    _ int = iota
    initialized
    done
)

const crlf = "\r\n"

func RequestFromReader(reader io.Reader) (*Request, error) {
    buf := make([]byte, 8, 8)
    if err != nil {
        return nil, err
    }
    requestLine, err := parseRequestLine(rawBytes)
    if err != nil {
        return nil, err
    }
    return &Request{
        RequestLine: *requestLine,
    }, nil
}

func parseRequestLine(data []byte) (int, error) {
    if !bytes.Contains(data, []byte(crlf)) {
        return 0, nil
    }
    idx := bytes.Index(data, []byte(crlf))
    requestLineText := string(data[:idx])
    requestLine, err := requestLineFromString(requestLineText)
    if err != nil {
        return nil, err
    }
    numBytes := len(crlf) + len(data[:idx])
    return numBytes, nil
}

func requestLineFromString(str string) (*RequestLine, error) {
    parts := strings.Split(str, " ") 
    if len(parts) != 3 {
        return nil, fmt.Errorf("poorly formatted request-line: %s", str)
    }

    method := parts[0]
    for _, c := range method {
        if c < 'A' || c > 'Z' {
	    return nil, fmt.Errorf("invalid method: %s", method)
	}
    }

    requestTarget := parts[1]

    versionParts := strings.Split(parts[2], "/")
    if len(versionParts) != 2 {
        return nil, fmt.Errorf("malformed start-line: %s", str)
    }

    httpPart := versionParts[0]
    if httpPart != "HTTP" {
        return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpPart)
    }
    version := versionParts[1]
    if version != "1.1" {
        return nil, fmt.Errorf("unrecognized HTTP-version: %s", version)
    }

    return &RequestLine{
        Method:		method,
	RequestTarget:  requestTarget,
	HttpVersion:	versionParts[1],
    }, nil
}

func (r *Request) parse(data []byte) (int, error) {
    switch r.State {
        case 1:
	    parsedBytes, err := r.parseRequestLine(data)
            if err != nil {
                return nil, fmt.Errorf("State isn't initialized: %s", err)
            }
	    if parsedBytes == 0 {
	        return 0, nil
	    }
	    r.State = done
	case 2:
	    return nil, fmt.Errorf("error: trying to read data in a done state: %s", err)
	default:
	    return nil, fmt.Errorf("error: unknown state: %s", err)
    }
}
