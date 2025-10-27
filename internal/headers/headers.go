package headers

import (
    "bytes"
)

const crlf = "\r\n"

type Headers map[string]string

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
    idx := bytes.Index(data, []byte(crlf))
    if idx == -1 {
        return 0, false, nil
    }
    if idx == 0 {
        return 2, true, nil
    }
    line := data[:idx]
    colonIdx := bytes.IndexByte(line, ":")
    if colonIdx == -1 {
        return 0, false, err
    }

    h[key] = value
    return n, done, err
}
