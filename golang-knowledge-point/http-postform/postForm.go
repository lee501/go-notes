package http_postform

import (
	"net/http"
	"net/url"
	"strings"
)

func PostWithForm() {
	payload := make(url.Values)
	payload.Add("name", "lee")
	payload.Add("password", "")
	req, err := http.NewRequest(
			http.MethodPost,
			"",
			strings.NewReader(payload.Encode()),
		)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r, err := http.DefaultClient.Do(req)
}

//use http postform
func postForm() {
	payload := make(url.Values)
	payload.Add("name", "poloxue")
	payload.Add("password", "123456")
	r, err := http.PostForm("http://httpbin.org/post", payload)
}