package zcpt

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Gzip(data []byte) []byte {
	bf := NewBuff()
	defer bf.Free()
	gz := gzip.NewWriter(bf)
	_, _ = gz.Write(data)
	_ = gz.Flush()
	_ = gz.Close()
	return bf.Clone()
}
func UnGzip(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	return io.ReadAll(gz)
}
