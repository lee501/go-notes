package chunk_upload

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const tmpFilePath = "../tmp/"

var wait sync.WaitGroup

func ChunkFile(w http.ResponseWriter, r *http.Request) {
	//设置跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("context-type", "application/json")

	res, err := mergeChunk(r)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	} else {
		// 保存上传节点
		_, _ = w.Write([]byte("上传成功:" + strconv.Itoa(res)))
	}
}

func mergeChunk(r *http.Request) (int, error) {
	size, _ := strconv.Atoi(r.FormValue("fileSize"))
	total, _ := strconv.Atoi(r.FormValue("chunkSum"))
	//上传文件到缓存路径下
	chunkUpload(r)
}

func chunkUpload(r *http.Request) (int, error) {
	//获取分片号和文件part
	chunkIndex := r.FormValue("chunkindex")
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return 0, errors.New("fetch form file failed")
	}

	//创建临时文件保存chunk
	filePath := fmt.Sprintf("%s%s_%s", tmpFilePath, fileHeader.Filename, chunkIndex)
	_, err = createFile(filePath)
	if err != nil {
		return 0, err
	}
	//获取当前文件大小
	os.Stat(filePath)
}

//创建文件
func createFile(filepath string) (bool, error) {
	filebool, err := isExistFile(filepath)
	if filebool && err == nil {
		return true, errors.New("file exist")
	} else {
		file, err := os.Create(filepath)
		defer file.Close()
		if err != nil {
			return false, errors.New("file create failed")
		}
	}
	return true, nil
}

//check 文件是否已经创建
func isExistFile(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
