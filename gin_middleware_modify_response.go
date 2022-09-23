package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vearne/gin-timeout/buffpool"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func main() {
	router := gin.New()

	router.Use(Handler)

	router.GET("/ping", Ping)

	server := http.Server{
		Addr:    "localhost:8999",
		Handler: router,
	}
	_ = server.ListenAndServe()
}

func Handler(c *gin.Context) {
	w := NewResponseWriter(c)
	c.Writer = w
	startTime := time.Now()
	defer w.Done(c)
	c.Next()
	duration := int(time.Since(startTime).Milliseconds())

	//status := w.Status()
	//data := w.body.Bytes()
	w.body.Reset()

	w.body.Write([]byte("hello"))
	w.Header().Set("Content-Length", strconv.Itoa(w.body.Len()))
	w.Flush()
	//c.DataFromReader(200, 5, "text/plain; charset=utf-8", io.NopCloser(bytes.NewBuffer([]byte("hello"))), map[string]string{})
	c.Header("X-Response-Time", strconv.Itoa(duration))
}

func Handler1(c *gin.Context) {
	w := NewResponseWriter(c)
	c.Writer = w
	defer w.Done(c)
	c.Next()
}

type ResponseWriter struct {
	gin.ResponseWriter
	h    http.Header
	body *bytes.Buffer

	code        int
	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
}

func NewResponseWriter(c *gin.Context) *ResponseWriter {
	buffer := buffpool.GetBuff()
	writer := &ResponseWriter{
		body:           buffer,
		ResponseWriter: c.Writer,
		h:              make(http.Header),
	}
	return writer
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.timedOut {
		return 0, nil
	}

	return w.body.Write(b)
}

func (w *ResponseWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.timedOut {
		return
	}
	w.writeHeader(code)
}

func (w *ResponseWriter) writeHeader(code int) {
	w.wroteHeader = true
	w.code = code
}

func (w *ResponseWriter) WriteHeaderNow() {}

func (w *ResponseWriter) Header() http.Header {
	return w.h
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

func (w *ResponseWriter) Done(c *gin.Context) {
	dst := w.ResponseWriter.Header()
	for k, vv := range w.Header() {
		dst[k] = vv
	}

	if !w.wroteHeader {
		w.code = http.StatusOK
	}

	w.ResponseWriter.WriteHeader(w.code)
	_, err := w.ResponseWriter.Write(w.body.Bytes())
	if err != nil {
		panic(err)
	}
	buffpool.PutBuff(w.body)
}
