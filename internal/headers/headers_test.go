package headers

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeadersParse(t *testing.T) {
    // Test: Valid single header
    headers := NewHeaders()
    data := []byte("Host: localhost:42069\r\n\r\n")
    n, done, err := headers.Parse(data)
    require.NoError(t, err)
    require.NotNil(t, headers)
    assert.Equal(t, "localhost:42069", headers["Host"])
    assert.Equal(t, 23, n)
    assert.False(t, done)

    // Test: Invalid spacing header
    headers = NewHeaders()
    data = []byte("       Host : localhost:42069       \r\n\r\n")
    n, done, err = headers.Parse(data)
    require.Error(t, err)
    assert.Equal(t, 0, n)
    assert.False(t, done)

    // Test: Valid single header with extra whitespace
    headers = NewHeaders()
    data = []byte("Host: localhost   :   42069\r\n\r\n")
    n, done, err = headers.Parse(data)
    require.NoError(t, err)
    require.NotNil(t, headers)
    assert.Equal(t, "localhost   :   42069", headers["Host"])
    assert.Equal(t, 29, n)
    assert.False(t, done)

    // Test: Valid two headers with existing headers
    headers = NewHeaders()
    n, done, err = headers.Parse([]byte("Host: localhost:42069\r\n\r\n"))
    require.NoError(t, err)
    assert.Equal(t, "localhost:42069", headers["Host"])
    assert.Equal(t, 23, n)
    assert.False(t, done)

    n, done, err = headers.Parse([]byte("Content-Type: text/plain\r\n\r\n"))
    require.NoError(t, err)
    require.NotNil(t, headers)
    assert.Equal(t, "text/plain", headers["Content-Type"])
    assert.Equal(t, 26, n)
    assert.False(t, done)
}
