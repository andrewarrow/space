package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func DoHttpRead(request *http.Request) (string, int) {
	client := &http.Client{Timeout: time.Second * 5}
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Accept-Encoding", "application/json")
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500
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

	jsonString, code := DoHttpRead(request)
	fmt.Println(jsonString)
	return code
}
