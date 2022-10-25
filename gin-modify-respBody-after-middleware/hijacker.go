package hijacker

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/vearne/gin-timeout/buffpool"
)

/*
	修改gin middleware after 的hijacker
*/
type ResponseWriter struct {
	gin.ResponseWriter
	h    http.Header
	Body *bytes.Buffer

	code        int
	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
}

func NewResponseWriter(c *gin.Context) *ResponseWriter {
	buffer := buffpool.GetBuff()
	writer := &ResponseWriter{
		Body:           buffer,
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

	return w.Body.Write(b)
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
	_, err := w.ResponseWriter.Write(w.Body.Bytes())
	if err != nil {
		panic(err)
	}
	buffpool.PutBuff(w.Body)
}

//middleware
/*
w := NewResponseWriter(ctx)
ctx.Writer = w
defer w.Done(ctx)
ctx.Next()
m.doAfter(ctx, beforeRes, w)

w.Body.Reset()
			w.Body.Write([]byte(resp.RawBody))
			w.Header().Set("Content-Length", strconv.Itoa(w.Body.Len()))
*/
