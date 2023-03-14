package utils

import (
	"github.com/mochafiqri/monorepo/fetch/commons/constants"
	"io"
	"net"
	"net/http"
	"time"
)

type HttpUtils struct {
	Client *http.Client
}

type Response struct {
	Code int
	Body []byte
}

func setTransport() *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 60 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}
}

func setHttpClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * 60,
		Transport: setTransport(),
	}
}

func GetHttpClient() HttpUtils {
	return HttpUtils{
		setHttpClient(),
	}
}

func (h *HttpUtils) DoRequest(method, url, contentType string, body io.Reader, header map[string]string) (Response, error) {
	var resp = Response{}
	request, err := http.NewRequest(method,
		url, body)
	if err != nil {
		return resp, err
	}

	if contentType != "" {
		request.Header.Add(constants.ContentType, contentType)
	}

	if len(header) > 0 {
		for i, v := range header {
			request.Header.Add(i, v)
		}
	}

	response, err := h.Client.Do(request)
	if err != nil {
		return resp, err
	}

	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}
	resp.Body = res
	resp.Code = response.StatusCode
	return resp, nil
}
