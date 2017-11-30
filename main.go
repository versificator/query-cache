package main

import( "net/http"
"github.com/Vertamedia/chproxy/log")

func main() {
	log.Infof("Serving https")
	http.HandleFunc("/", ParseClientPost)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error : %s ", err)
	}
}


