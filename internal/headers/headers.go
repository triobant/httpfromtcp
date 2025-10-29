package headers

import (
    "bytes"
    "fmt"
)

const crlf = "\r\n"

type Headers map[string]string

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
    idx := bytes.Index(data, []byte(crlf))
    if idx == -1 {
        return 0, false, nil
    }
    if idx == 0 {
        return len(crlf), true, nil
    }

    line := data[:idx]

    colonIdx := bytes.IndexByte(line, ':')
    if colonIdx == -1 {
        return 0, false, fmt.Errorf("missing colon")
    }

    rawKey := line[:colonIdx]
    if len(rawKey) == 0 || rawKey[len(rawKey)-1] == ' ' {
        return 0, false, fmt.Errorf("space before colon in field-name")
    }

    rawVal := line[colonIdx+1:]
    key := string(bytes.TrimSpace(rawKey))
    val := string(bytes.TrimSpace(rawVal))

    h[key] = val
    return idx + len(crlf), false, nil
}
