package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/andybalholm/brotli"
)

func TestCC(t *testing.T) {
	f, err := os.Open("D:\\codespace\\rooui\\asset\\example\\simple\\server\\public\\main.wasm")
	if err != nil {
		t.Fatalf("Error opening test data file: %v", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error reading test data: %v", err)
	}

	out := bytes.Buffer{}
	e := brotli.NewWriterOptions(&out, brotli.WriterOptions{Quality: 11})
	n, err := e.Write(data)
	if err != nil {
		t.Errorf("Error compressing data: %v", err)
	}
	if int(n) != len(data) {
		t.Errorf("Write() n=%v, want %v", n, len(data))
	}
	if err := e.Close(); err != nil {
		t.Errorf("Close Error after writing %d bytes: %v", n, err)
	}

	d, err := os.Create("D:\\codespace\\rooui\\asset\\example\\simple\\server\\public\\main2.wasm")
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	_, err = d.Write(out.Bytes())
	if err != nil {
		t.Fatal(err)
	}
}
