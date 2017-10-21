package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var savePath = flag.String("save-path", "", "Path to archive the files")
	var onlyNew = flag.Bool("only-new", false, "Only retrieve the new files")
	var markArchive = flag.Bool("archive", false, "Archive the retrieved documents")

	flag.Parse()

	var configuration = ReadConfiguration()
	var requestOptions = DefaultRequestOptions()

	var documentsResult DocumentsResult

	if *onlyNew {
		documentsResult = GetNewDocuments(configuration, requestOptions)
	} else {
		documentsResult = GetDocuments(configuration, requestOptions)
	}

	if len(documentsResult.Documents) > 0 {
		for _, document := range documentsResult.Documents {
			var filename = strings.Join([]string{strings.Replace(document.Name, "/", "-", 999), ".pdf"}, "")
			filename = strings.Join([]string{document.Sender.Label, filename}, " - ")

			fmt.Printf("Performing actions on document %s\n", filename)

			if *markArchive {
				fmt.Print(" - Archiving ...\n")
				document.Archive(configuration)
			}
			if len(*savePath) > 0 {
				fmt.Printf(" - Downloading to %s%s\n", *savePath, filename)
				_, err := document.Download(configuration, *savePath, filename)
				if err != nil {
					fmt.Printf(" -- Could not store the file: %s\n", err)
					break
				}
			}
		}
	} else {
		fmt.Printf("There are no documents to be processed\n")
	}

}
