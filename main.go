package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type FileHandler struct {
	Config Config
}

func (f FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var ff *os.File
	var err error

	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	filename := path.Clean(upath)

	routes := f.Config.Routes
	fmt.Println(filename)

	if filename == "/" {
		ff, err = os.Open("./static/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer ff.Close()
		b, _ := ioutil.ReadAll(ff)

		for _, r := range routes {
			a, _ := regexp.Compile(r.Src)
			isMatch := a.MatchString(filename)
			if isMatch {
				for k, v := range r.Headers {
					w.Header().Add(k, v)
				}
			}
		}
		w.Write(b)
		return
	}

	filename = fmt.Sprintf("./static/%s", strings.TrimPrefix(filename, "/"))
	fmt.Println(filename)
	ff, err = os.Open(filename)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer ff.Close()
	bb, _ := ioutil.ReadAll(ff)

	for _, r := range routes {
		if r.Src == "/" {
			continue
		}
		a, _ := regexp.Compile(r.Src)
		isMatch := a.MatchString(filename)
		if isMatch {
			for k, v := range r.Headers {
				w.Header().Set(k, v)
			}
		}
	}

	ctype := mime.TypeByExtension(filepath.Ext(filename))
	w.Header().Set("Content-Type", ctype)

	fmt.Println(ctype)

	w.Write(bb)
}

func main() {

	var err error
	var config Config

	r := mux.NewRouter()

	b, err := ioutil.ReadFile("now.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}

	fileHandler := &FileHandler{
		Config: config,
	}

	r.PathPrefix("/").Handler(fileHandler)
	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}

}
