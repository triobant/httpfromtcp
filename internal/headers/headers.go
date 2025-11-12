package headers

import (
    "bytes"
    "fmt"
    "strings"
    "unicode"
)

const crlf = "\r\n"

type Headers map[string]string

func NewHeaders() Headers {
    return map[string]string{}
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
    idx := bytes.Index(data, []byte(crlf))
    if idx == -1 {
        return 0, false, nil
    }
    if idx == 0 {
        return 2, true, nil
    }

    parts := bytes.SplitN(data[:idx], []byte(":"), 2)
    key := string(parts[0])

    if key != strings.TrimRight(key, " ") {
        return 0, false, fmt.Errorf("invalid header name: %s", key)
    }

    value := bytes.TrimSpace(parts[1])
    key = strings.TrimSpace(key)

    key = strings.ToLower(key)
    for _, r := range key {
	if unicode.IsLetter(r) {
            fmt.Printf("It is a letter: %q", r)
	}
	if unicode.IsDigit(r) {
            fmt.Printf("It is a digit: %q", r)
	}
	fmt.Printf("No letter but...\nrune: %q byte: %d\n", r, len([]byte(string(r))))
    }

    h.Set(key, string(value))
    return idx + 2, false, nil
}

func (h Headers) Set(key, value string) {
    h[key] = value
}
