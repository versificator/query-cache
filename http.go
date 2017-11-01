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
	io.WriteString(rw, `{"alive": true}`)
	body, _ := ioutil.ReadAll(request.Body)
	io.WriteString(rw, string(body))

	response := string(body)

	if checkIfCacheExists(&response) {

	} else {
		getDataFromDB(response)
	}
	rw.WriteHeader(http.StatusOK)
}
