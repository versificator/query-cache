package main

import (

	"net/http"
	"io"
	"io/ioutil"
	"github.com/Vertamedia/chproxy/log"
	"strings"
	"fmt"
	"net/url"
)

func StartHTTP(){
	log.Infof("Serving https")
	http.HandleFunc("/", ParseClientPost)
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatalf("Server error : %s ", err)
	}
}

func ParseClientPost(rw http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(rw, "Please send a request body", 400)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	io.WriteString(rw, `{"alive": true}`)
	body, _ := ioutil.ReadAll(request.Body)
	io.WriteString(rw, string(body))

	request_body := string(body)

	fmt.Println("Request_body: ", request_body)
    q := parse(request_body)

    fmt.Println(toString(q))
	ff:=getHourRanges(q.Filter)
	for _, element := range ff {
		fmt.Print(element.StartDay," ")
		fmt.Print(element.EndDay," ")
		fmt.Print(element.StartHour," ")
		fmt.Println(element.EndHour," ")

		d := query{q.Query,element}

		fmt.Println(toString(d))


		req, err := http.NewRequest("GET", "http://localhost:9090?query="+url.QueryEscape(toString(d)),nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		f, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(f))
	}

	req, err := http.NewRequest("GET", "http://localhost:9090?query=select%201%20union%20all%20select%202", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.StatusCode)
	response,err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(response))



	rw.WriteHeader(http.StatusOK)
}
///////////////

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


