package rest

import (
	"bytes"
	"io"
	"net/http"
)

type Api struct {
	checkServerUrl string 
}

func NewApi(checkServerUrl string) *Api {
	return &Api{checkServerUrl: checkServerUrl}
}

func (a *Api) SendUrls(request []byte) []byte {
	bodyReader := bytes.NewReader(request)

	req, _ := http.NewRequest(http.MethodPost, a.checkServerUrl, bodyReader)
	req.Header.Set("X-Custom-Header", "telegram-client-request")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

	return body
}