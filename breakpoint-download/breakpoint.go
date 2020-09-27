package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

//文件分片
type filePart struct {
	Index int   //分片的序号
	From  int	//开始的byte
	To    int	//解决的byte
	Data  []byte  //http下载得到的文件内容
}

//文件下载
type FileDownloader struct {
	fileSize	int
	url			string
	filename 	string
	totalPart	int
	outFileDir		string
	doneFilePart []filePart
}

func NewFileDownloader(url, filename, outFileDir string, totalPart int) *FileDownloader {
	if outFileDir == "" {
		wd, _ := os.Getwd()
		outFileDir = wd
	}
	return &FileDownloader{
		fileSize: 0,
		url: url,
		filename: filename,
		totalPart: totalPart,
		outFileDir: outFileDir,
		doneFilePart: make([]filePart, totalPart),
	}
}

func main() {
	startTime := time.Now()
	var url string
	url = "https://download.jetbrains.com/go/goland-2020.2.2.dmg"
	downloader := NewFileDownloader(url, "", "", 10)
	if err := downloader.Run(); err != nil {
		// fmt.Printf("\n%s", err)
		log.Fatal(err)
	}
	fmt.Printf("\n 文件下载完成耗时: %f second\n", time.Now().Sub(startTime).Seconds())
}

func (f *FileDownloader) head() (int, error) {
	r, err := f.getNewRequest("HAED")
	if err != nil {
		return 0, err
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	if res.StatusCode > 299 {
		return 0, errors.New(fmt.Sprintf("Can't process, response is %v", res.StatusCode))
	}
//检测是否支持断点续传
   if res.Header.Get("Accept-Ranges") != "bytes" {
   		return 0, errors.New("服务器不支持断点续传")
   }
   f.filename = parseFileInfo(res)
   return strconv.Atoi(res.Header.Get("Content-Length"))
}

func (f *FileDownloader) Run() error {
	fileSize, err := f.head()
	if err != nil {
		return err
	}
	f.fileSize = fileSize
	//开启分片下载
	jobs := make([]filePart, f.totalPart)
	eachSize := fileSize / f.totalPart
	//fmt.Println(fileSize, eachSize)
	for i := range jobs {
		jobs[i].Index = i
		if i == 0 {
			jobs[i].From = 0
		} else {
			jobs[i].From = jobs[i-1].To + 1
		}
		if i < f.totalPart - 1 {
			jobs[i].To = jobs[i].From + eachSize
		} else {
			jobs[i].To = fileSize - 1
		}
	}
	//fmt.Println(jobs)

	var wg sync.WaitGroup
	for _, j := range jobs {
		wg.Add(1)
		go func(job filePart) {
			defer wg.Done()
			err := f.download(job)
			if err != nil {
				log.Println("下载文件失败:", err, job)
			}
		}(j)
	}
	wg.Wait()
	return f.mergeFileParts()
}

func (f * FileDownloader) getNewRequest(method string) (*http.Request, error) {
	r, err := http.NewRequest(
		method,
		f.url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	return r, nil
}

func parseFileInfo(res *http.Response) string {
	contentDisposition := res.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err != nil {
			panic(err)
		}
		return params["filename"]
	}
	filename := filepath.Base(res.Request.URL.Path)
	return filename
}

func (f *FileDownloader) download(part filePart) error {
	r, err := f.getNewRequest("GET")
	if err != nil {
		return err
	}
	log.Printf("开始[%d]下载from:%d to:%d\n", part.Index, part.From, part.To)
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", part.From, part.To))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("服务器错误状态码: %v", resp.StatusCode))
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(bs) != (part.To - part.From + 1) {
		return errors.New("下载文件分片长度错误")
	}
	part.Data = bs
	f.doneFilePart[part.Index] = part
	return nil
}

func (f *FileDownloader) mergeFileParts() error {
	log.Println("开始合并文件")
	path := filepath.Join(f.outFileDir, f.filename)
	file, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	hash := sha256.New()
	totalSize := 0

	for _, filepart := range f.doneFilePart {
		file.Write(filepart.Data)
		hash.Write(filepart.Data)
		totalSize += len(filepart.Data)
	}

	if totalSize != f.fileSize {
		return errors.New("文件不完整")
	}
	if hex.EncodeToString(hash.Sum(nil)) != "3af4660ef22f805008e6773ac25f9edbc17c2014af18019b7374afbed63d4744" {
		return errors.New("文件损坏")
	} else {
		log.Println("文件SHA-256校验成功")
	}
	return nil
}