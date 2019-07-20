package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/hello"

	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"name\"\r\n\r\nsurname\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	req, err := http.NewRequest("GET", url, payload)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "b48ebe37-799e-4857-8ec5-71753c25014b")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(string(body))

}
