package main

import (

	"net/http"
	"io"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/Vertamedia/chproxy/log"

)

type test_struct struct {
	Test string
}

func StartHTTP(){
	log.Infof("Serving https")
	http.HandleFunc("/", ParseGhPost)
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatalf("Server error : %s ", err)
	}
}

func ParseGhPost(rw http.ResponseWriter, request *http.Request) {

	if request.Body == nil {
		http.Error(rw, "Please send a request body", 400)
		return
	}

    rw.Header().Set("Content-Type", "application/json")
	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(rw, `{"alive": true}`)

	body, _ := ioutil.ReadAll(request.Body)
	io.WriteString(rw,string(body))

	g := string(body)

	if checkIfCacheExists(&g) {

	} else {
		getData(g)
	}
//	getHash(&g)
//	parseRoot(&g)
//	io.WriteString(rw,"\n")
//    io.WriteString(rw,formatRequest(request))
//	decoder := json.NewDecoder(request.Body)
	//
	//var t test_struct
	//err := decoder.Decode(&t)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(t.Test)
	rw.WriteHeader(http.StatusOK)

}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
	r.ParseForm()
	request = append(request, "\n")
	request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
