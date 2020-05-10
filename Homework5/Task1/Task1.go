package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// type Request represents request info
type Request struct{
	Host       string  `json:"host"`
	UserAgent  string  `json:"user_agent"`
	RequestURL string  `json:"request_url"`
	Headers    Headers `json:"headers"`
}

// type Headers represents request headers
type Headers struct{
	Accept    string `json:"Accept"`
	UserAgent string `json:"User-Agent"`
}

// function SomeHandler handles request,
// returns request info in JSON format.
func someHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/alex" {
		http.NotFound(w, r)
		return
	}
	request := Request{
		r.Host,
		r.URL.Path,
		r.UserAgent(),
		Headers{
				r.Header.Get("Accept"),
				r.Header.Get("User-Agent"),
		},
	}
	jsonReqInfo, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err = w.Write(jsonReqInfo); err != nil{
		log.Fatal(err)
		return
	}
}

func main() {
	http.HandleFunc("/user/alex", someHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
		return
	}
}
