package headers

import (
    "bytes"
)

const crlf = "\r\n"

type Headers map[string]string

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
    h[key] = value
    return n, done, err
}
