package tool

import (
	"io"
	"net/http"
)

func HttpPostRequest(urlValue map[string][]string) ([]byte, error) {
	resp, err := http.PostForm("http://127.0.0.1/api.php", urlValue)
	if err != nil {
		panic(err)
	}

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return byteResp, err
}
