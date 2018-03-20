package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"os"
	"bufio"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path[1:]
	log.Println(path)
	f, err := os.Open(path)

	data, err := ioutil.ReadFile(string(path))
	if err == nil{
		bufferedReader := bufio.NewReader(f)
		var contentType string
		if strings.HasSuffix(path, ".css"){
			contentType = "text/css"
		}else if strings.HasSuffix(path, ".html"){
			contentType = "text/html"
		}else if strings.HasSuffix(path, ".js"){
			contentType = "text/js"
		}else if strings.HasSuffix(path, ".png"){
			contentType = "image/png"
		}else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		}else if strings.HasSuffix(path, ".mp4"){
			contentType = "video/mp4"
		}else{
			contentType = "text/plain"
		}
		w.Header().Add("Content Type", contentType)
		//w.Write(data)
		bufferedReader.WriteTo(w)
	}else{
		w.WriteHeader(404)
		w.Write([]byte("404 Mi amor - " + http.StatusText(404)))
	}
}

func main()  {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080",nil)
}