package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client *http.Client

func NewHttpClient(timeout int) {
	client = &http.Client{Timeout: time.Duration(timeout) * time.Second}
}

func main() {
	NewHttpClient(50000)
	http.HandleFunc("/proxy", ForwardHttpRequest)
	http.ListenAndServe(":5000", nil)
}

func ForwardHttpRequest(w http.ResponseWriter, r *http.Request) {

	headers := r.Header.Clone()

	log.Printf("Received at URL [%v]", r.URL)
	log.Printf("Path [%v]", r.URL.RawPath)
	log.Printf("RawQuery [%v]", r.URL.RawQuery)
	log.Printf("Newhost [%v]", os.Getenv("SERVERHOST"))

	fwd2Host := os.Getenv("SERVER_HOST")
	fwd2Port := os.Getenv("SERVER_PORT")

	u, err := url.Parse("http://" + fwd2Host + ":" + fwd2Port + "/books")
	log.Printf("NEW URL [%v]", u)

	request := &http.Request{
		Method: r.Method,
		URL:    u,
		Header: headers,
		Body:   r.Body,
	}

	log.Printf("forwarded to service url [%v]", u)

	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	log.Printf("response received from forwarded service [%v]", err)

	w.WriteHeader(response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	} else {
		w.Write([]byte([]byte(body)))
	}

}
