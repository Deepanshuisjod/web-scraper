package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// With the status code equals to 200 it will parse the
// HTML content
func parseContent(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	tokenizer := html.NewTokenizer(strings.NewReader(string(body)))
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return
		}
		fmt.Printf("Token: %v\n", token.String())
	}
}

func noContent(resp *http.Response) {
	if resp.ContentLength == 0 {
		fmt.Println("The Response body has no content")
	}
}

func clientError() {
	//
}

func handleResponse(statusCode string, resp *http.Response) {
	code_list := strings.Split(statusCode, " ")
	code := code_list[0]

	switch {
	case code == "200": // The request is successful
		parseContent(resp)
	case code == "204": // Server successfully processed the request but has no content
		noContent(resp)
	case code == "301":
		//
	case code == "302":
		//
	case code == "304":
		//
	case code == "400": // Cannot process due to client error
		clientError()
	case code == "401":
		// Unauthorized
	case code == "403":
		// Forbidden
	case code == "404":
		// Not Found
	}

}

func main() {
	// Specify the URL correctly, including the protocol and host.
	url := "http://localhost:8080/sample.html"

	// Send the HTTP GET request and get the response.
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response.
	fmt.Println(resp)
	// Passing the status code to handleResponse
	handleResponse(resp.Status, resp)
	fmt.Println(resp.ContentLength)
}
