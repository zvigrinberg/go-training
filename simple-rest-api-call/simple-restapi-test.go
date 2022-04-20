package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		fmt.Print("Error getting from google")
	}

	req.Header.Set("Content-Type", "text/html")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error getting from google, details : %v,", err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Print(string(body))

}

