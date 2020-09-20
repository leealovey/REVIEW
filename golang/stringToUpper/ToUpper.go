package stringToUpper

import "io"
import "bytes"

type UpperWrite struct {
	io.Writer
}

func (p *UpperWrite) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}