package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func downloadFile(url string, outputPath string) {
	fmt.Printf("Download: %v\n", url)
	req, _ := http.NewRequest("GET", url, nil)
	r, err := http.Get(req.URL.String())
	if err != nil {
		fmt.Printf("Download file %v: %v", url, err)
		os.Exit(-1)
	}
	defer r.Body.Close()

	filename := path.Base(req.URL.Path)
	out, err := os.Create(path.Join(outputPath, filename))
	if err != nil {
		fmt.Printf("Create file %v: %v", filename, err)
		os.Exit(-1)
	}
	defer out.Close()
	io.Copy(out, r.Body)
}
