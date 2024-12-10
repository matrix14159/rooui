package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

type compressWriter struct {
	gin.ResponseWriter
	writer io.Writer
}

func (p *compressWriter) WriteString(s string) (int, error) {
	p.Header().Del("Content-Length")
	return p.writer.Write([]byte(s))
}

func (p *compressWriter) Write(data []byte) (int, error) {
	p.Header().Del("Content-Length")
	return p.writer.Write(data)
}

// Fix: https://github.com/mholt/caddy/issues/38
func (p *compressWriter) WriteHeader(code int) {
	p.Header().Del("Content-Length")
	p.ResponseWriter.WriteHeader(code)
}
