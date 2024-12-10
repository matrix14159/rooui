package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	listen = flag.String("listen", ":12000", "listen address")
	dir    = flag.String("dir", "./public", "directory to serve")
)

func main() {
	flag.Parse()
	//log.Printf("listening on %q...", *listen)
	//err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	//log.Fatalln(err)

	r := gin.Default()
	//r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(Compress())

	r.Static("/", "D:\\codespace\\rooui\\asset\\example\\simple\\server\\public")

	if err := r.Run(*listen); err != nil {
		log.Fatalf("artroo start web server error:%v", err)
	}
}
