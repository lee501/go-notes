* File
---
    file is an interface to access the file part of a multipart message
    type File interface {
        io.Reader
        io.ReaderAt
        io.Seeker
        io.Closer
    }
    
* FileHeader封装了文件的基本信息

---
    A FileHeader describes a file part of a multipart request.
    type FileHeader struct {
        Filename string					//文件名
        Header   textproto.MIMEHeader	//MIME信息
        Size     int64					//文件大小,单位bit
        content []byte					//文件内容,类型[]byte
        tmpfile string					//临时文件
    }