package main

import (
	"io"
	"os"
	"strings"
)

type (
	DocumentsResult struct {
		HasMore       bool        `json:"hasMore"`
		NextPage      int         `json:"nextPage"`
		NumberOfPages int         `json:"numberOfPages"`
		PageIndex     int         `json:"pageIndex"`
		PageSize      int         `json:"pageSize"`
		PreviousPage  interface{} `json:"previousPage"`
		Results       int         `json:"results"`
		SortField     string      `json:"sortField"`
		SortFieldType string      `json:"sortFieldType"`
		SortOrder     string      `json:"sortOrder"`
		TotalResults  int         `json:"totalResults"`
		Documents     []Document  `json:"documents"`
	}

	Document struct {
		Actions            []Action    `json:"actions"`
		CanOpen            bool        `json:"canOpen"`
		Categories         []string    `json:"categories"`
		ContentURL         string      `json:"contentUrl"`
		Name               string      `json:"name"`
		Opened             bool        `json:"opened"`
		Payment            interface{} `json:"payment"`
		PresentationType   string      `json:"presentationType"`
		PublishDate        string      `json:"publishDate"`
		Sender             Sender      `json:"sender"`
		SenderDocumentType string      `json:"senderDocumentType"`
		ShortName          interface{} `json:"shortName"`
		URI                string      `json:"uri"`
	}
)

// Download the document's file
func (document Document) Download(configuration Configuration, path string, filename string) (int64, error) {
	var resp = DoRequest(configuration, document.ContentURL, "GET")
	defer resp.Body.Close()

	out, err := os.Create(strings.Join([]string{path, filename}, ""))
	defer out.Close()

	if err != nil {
		return 0, err
	}

	n, err := io.Copy(out, resp.Body)

	return n, err
}

// Archive the document
func (document Document) Archive(configuration Configuration) {
	for _, action := range document.Actions {
		if action.Label == "ARCHIVE" && action.Enabled {
			// This is very trustworthy at the moment...
			var resp = DoRequest(configuration, action.URL, action.Method)
			defer resp.Body.Close()
		}
	}
}
