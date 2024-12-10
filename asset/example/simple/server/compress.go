package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/andybalholm/brotli"
	"github.com/gin-gonic/gin"
)

// Compress is a middleware for gin
// if possible, use brotli first, otherwise use gzip
func Compress() gin.HandlerFunc {
	handler := new(compressHandler)
	handler.brPool = sync.Pool{
		New: func() interface{} {
			return brotli.NewWriterLevel(io.Discard, 3)
		},
	}
	handler.gzPool = sync.Pool{
		New: func() interface{} {
			gz, err := gzip.NewWriterLevel(io.Discard, gzip.DefaultCompression)
			if err != nil {
				panic(err)
			}
			return gz
		},
	}
	return handler.Handle
}

type compressHandler struct {
	brPool sync.Pool
	gzPool sync.Pool
}

func (p *compressHandler) Handle(c *gin.Context) {
	alg, ok := p.shouldCompress(c.Request)
	if !ok {
		return
	}

	//var writer io.WriteCloser
	switch alg {
	case "br":
		//br := p.brPool.Get().(*brotli.Writer)
		//defer p.brPool.Put(br)
		//defer br.Reset(io.Discard)
		//br.Reset(c.Writer)
		//writer = br

		c.Header("Content-Encoding", "br")
		c.Header("Vary", "Accept-Encoding")

	case "gzip":
		//gz := p.gzPool.Get().(*gzip.Writer)
		//defer p.gzPool.Put(gz)
		//defer gz.Reset(io.Discard)
		//gz.Reset(c.Writer)
		//writer = gz
		//
		//c.Header("Content-Encoding", "gzip")
		//c.Header("Vary", "Accept-Encoding")
	}

	//rawWriter := c.Writer
	//c.Writer = &compressWriter{c.Writer, writer}
	defer func() {
		//writer.Close()
		//c.Writer = rawWriter
		c.Writer.Header().Del("Content-Length")
		c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
	}()
	c.Next()
}

func (p *compressHandler) shouldCompress(req *http.Request) (alg string, ok bool) {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") ||
		strings.Contains(req.Header.Get("Connection"), "Upgrade") ||
		strings.Contains(req.Header.Get("Accept"), "text/event-stream") {
		return
	}
	if !strings.Contains(req.URL.Path, ".wasm") {
		return
	}

	if strings.Contains(req.Header.Get("Accept-Encoding"), "br") {
		alg = "br"
		ok = true
		return
	}
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		alg = "gzip"
		ok = true
		return
	}
	return
}
