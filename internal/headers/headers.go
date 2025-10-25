package headers

import (
    "bytes"
)

const crlf = "\r\n"

type Headers map[string]string

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
    idx := bytes.Index(data, []byte(crlf))
    h[key] = value
    return n, done, err
}
