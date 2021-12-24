package neuralHash

import (
    "io/ioutil"
    "net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type NHResponse struct {
	Hash string `json:"hash"`
}

type ClientError struct {
  statusCode int
  message string
}

func (e *ClientError) Error() string {
	return "NH Client request failed with status code: " + strconv.Itoa(e.statusCode) + " and message: " + e.message
}

type NeuralHashClient struct {
	baseUrl string
	linkPath string
	uploadPath string
}

func (c *NHClient) getHashFromUrl(url string) (string, error) {
	postBody, _ := json.Marshal(map[string]string{
		"url":  url,
	})
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(c.baseUrl + c.linkPath, "application/json", requestBody)
    if err != nil {
		return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", &ClientError{resp.StatusCode, string(body)}
	}

	// snippet only
	var result = new(NHResponse)
	if err := json.Unmarshal(body, &result); err != nil {   // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

    return result.Hash, nil   
}

func Default() (*NeuralHashClient) {
	return &NeuralHashClient{"https://hash.kikkia.dev/api/", "link", "upload"}
}