package dom

import (
	"syscall/js"
)

// ArrayBufferToBlob convert the buf of [ArrayBuffer, ArrayBuffer, ...] to blob
func ArrayBufferToBlob(buf js.Value) (ret js.Value, err error) {
	blobConstructor := js.Global().Get("Blob")
	ret = blobConstructor.New(buf)
	return
}

// BlobToText convert the val of blob data to utf-8 text
// if encoding="", then default is utf-8
func BlobToText(val js.Value, encoding string) (ret string, err error) {
	ctr := js.Global().Get("FileReader")
	if encoding == "" {
		encoding = "utf-8"
	}
	reader := ctr.New()

	done := make(chan any)
	handleLoad := func(this js.Value, args []js.Value) any {
		defer close(done)
		ret = reader.Get("result").String()
		return nil
	}
	reader.Set("onload", js.FuncOf(handleLoad))
	reader.Call("readAsText", val, encoding)
	<-done
	return
}
