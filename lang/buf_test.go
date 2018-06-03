package lang

import (
	"bytes"
	"io"
	"testing"
)

func TestBuf(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	buf.Write([]byte{1, 2, 3, 4})

	i := 0
	for _, err := buf.ReadByte(); err != io.EOF; _, err = buf.ReadByte() {
		i++
	}
	t.Logf("eof at %d ", i)
}
