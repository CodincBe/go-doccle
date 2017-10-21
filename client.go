package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	// API_URL Url of the API
	API_URL = "https://secure.doccle.be/doccle-euui"
)

type (
	Action struct {
		Enabled bool   `json:"enabled"`
		ID      int    `json:"id"`
		Label   string `json:"label"`
		Method  string `json:"method"`
		URL     string `json:"url"`
	}

	Sender struct {
		ReferenceId string `json:"referenceId"`
		Label       string `json:"label"`
	}
)

// GetDocuments retrieves and returns an DocumentsResult struct
func GetDocuments(configuration Configuration, options RequestOptions) DocumentsResult {
	return retrieveDocumentsResult(
		configuration,
		strings.Join([]string{API_URL, "/rest/v2/documents?", options.toQuery()}, ""),
	)
}

// GetNewDocuments retrieves and returns an DocumentsResult struct with new documents only
func GetNewDocuments(configuration Configuration, options RequestOptions) DocumentsResult {
	return retrieveDocumentsResult(
		configuration,
		strings.Join([]string{API_URL, "/rest/v2/documents/new?", options.toQuery()}, ""),
	)
}

// DoRequest makes a request
func DoRequest(configuration Configuration, url string, method string) *http.Response {
	req, err := http.NewRequest(method, url, nil)

	req.SetBasicAuth(configuration.Username, configuration.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func retrieveDocumentsResult(configuration Configuration, url string) DocumentsResult {
	var resp = DoRequest(configuration, url, "GET")
	defer resp.Body.Close()

	var data = DocumentsResult{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
	}

	return data
}
