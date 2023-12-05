package network

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func DoHttpRead(request *http.Request) (string, int) {
	client := &http.Client{Timeout: time.Second * 5}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err == nil {
		ce := resp.Header.Get("Content-Encoding")
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500
		}
		if ce == "gzip" {
			buf := bytes.NewBuffer(body)
			gr, _ := gzip.NewReader(buf)
			defer gr.Close()
			body, _ = ioutil.ReadAll(gr)
		}
		return string(body), resp.StatusCode
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return err.Error(), 500
}

func DoPost(urlString string, payload any) int {
	asBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(asBytes)
	request, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return 500
	}

	_, code := DoHttpRead(request)
	return code
}
