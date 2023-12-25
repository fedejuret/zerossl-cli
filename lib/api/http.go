package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const API_URI = "https://api.zerossl.com"

func Post(path string, content interface{}) []byte {

	buf, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", API_URI+path+"?access_key="+os.Getenv("ZEROSSL_API_KEY"), bytes.NewBuffer(buf))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		response.Body.Close()
	}

	body, _ := io.ReadAll(response.Body)
	response.Body.Close()

	return body
}

func Get(path string, query map[string]string) []byte {

	u, err := url.Parse(API_URI + path)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("access_key", os.Getenv("ZEROSSL_API_KEY"))

	for key, value := range query {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()

	request, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		response.Body.Close()
	}

	body, _ := io.ReadAll(response.Body)
	response.Body.Close()

	return body
}
