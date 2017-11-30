package main

import (

	"net/http"
	"io"
	"io/ioutil"
	"strings"
	"fmt"
	"net/url"
)

func ParseClientPost(rw http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(rw, "Please send a request body", 400)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(request.Body)
	io.WriteString(rw, string(body))


	fmt.Println("Request_body: ", string(body))
    q := parse(body)

    fmt.Println("Query: ", toString(q))
	rangeList:=getHourRanges(q)

	fmt.Println("hour ranges: ",rangeList)

	for _, element := range rangeList {


		query_template := query{q.Query,element.StartHour,element.EndHour,element.StartDay,element.EndDay,[]string{""}}

		fmt.Println(toString(query_template))

		fmt.Println(url.QueryEscape(toString(query_template)))


		req, err := http.NewRequest("GET", "http://localhost:9090?query="+url.QueryEscape(toString(query_template)),nil)


		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.StatusCode)
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
		}
		processing(response)
		//fmt.Println(string(response))
	}

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


