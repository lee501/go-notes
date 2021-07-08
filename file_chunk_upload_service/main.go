package main

import (
	chunk "github.com/sevenelevenlee/go-notes/file_chunk_upload_service/chunk_upload"
	"log"
	"net/http"
	"os"
)

func main() {

}

func FileServer() {
	http.HandleFunc("/chunkUpload", chunk.ChunkFile)
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal("服务启动失败")
		os.Exit(-1)
	}
}
