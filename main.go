package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello, welcome to hzh's blog"))

	var indexData IndexData
	indexData.Title = "hzh's blog"
	indexData.Desc = "現在是入門教程"
	jsonstr, _ := json.Marshal(indexData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonstr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "hzh's blog"
	indexData.Desc = "现在是入门教程"

	//拿到当前路径
	t := template.New("index.html")
	path, _ := os.Getwd()
	//fmt.Printf("path = ", path)
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
