package main

import (
	"html/template"
	"log"
	"net/http"
	"runtime"

	"github.com/garyburd/redigo/redis"
)

var templates = template.Must(template.ParseFiles("main.html"))

type info struct {
	System  string
	CPU     string
	Counter int64
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	redi, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer redi.Close()

	res, err := redi.Do("incr", "counter")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	input := info{
		System: runtime.GOOS,
		CPU:    runtime.GOARCH,
	}
	if res, ok := res.(int64); ok {
		input.Counter = res
	}

	log.Printf("counter = %d", input.Counter)

	err = templates.ExecuteTemplate(w, "main.html", input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
