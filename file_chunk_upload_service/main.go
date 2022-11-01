package main

import (
	"log"
	"net/http"
	"os"

	chunk "github.com/sevenelevenlee/go-notes/file_chunk_upload_service/chunk_upload"
)

func FileServer() {
	http.HandleFunc("/chunkUpload", chunk.ChunkFile)
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal("服务启动失败")
		os.Exit(-1)
	}
}
