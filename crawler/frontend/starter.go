package main

import (
	"net/http"
	"spider/crawler/frontend/controller"
)

func main() {
	http.Handle("/", controller.CreateSearchResultHandle("crawler/frontend/view/index.html"))
	http.Handle("/search", controller.CreateSearchResultHandle("crawler/frontend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
