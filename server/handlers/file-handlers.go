package handlers

import (
	config "bobby-website/server/configurations"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//HandleFile will handle all css and js files
func HandleFile(wr http.ResponseWriter, req *http.Request) {
	var contentType, filePath = getContentType(req.URL.String())
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening file" + err.Error())
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileModTime := fileInfo.ModTime().String()
	fileSize := fileInfo.Size()
	eTag := fileModTime + strconv.Itoa(int(fileSize))

	wr.Header().Add("Content-Type", contentType)
	wr.Header().Add("ETag", eTag)

	if match := req.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, eTag) {
			wr.WriteHeader(http.StatusNotModified)
			return
		}
	}
	reader := bufio.NewReader(file)
	reader.WriteTo(wr)
}

func getContentType(file string) (contentType, path string) {
	switch true {
	case strings.HasSuffix(file, ".js"):
		contentType = "application/javascript"
		path = config.JSFILEPATH + file
	case strings.HasSuffix(file, ".css"):
		contentType = "text/css"
		path = config.CSSFILEPATH + file
	case strings.HasSuffix(file, ".png"):
		contentType = "image/png"
	case strings.HasSuffix(file, ".jpeg"), strings.HasSuffix(file, ".jpg"):
		contentType = "image/jpeg"
	case strings.HasSuffix(file, ".svg"):
		contentType = "image/svg"
	}
	return
}
