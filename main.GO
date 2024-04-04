package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "http://example.com"
	cookie := &http.Cookie{
		Name:   "your_cookie_name",
		Value:  "your_cookie_value",
		Path:   "/",
		Domain: ".example.com",
	}
	requestInterval := 5 * time.Second

	client := &http.Client{}

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.AddCookie(cookie)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		fmt.Printf("Received response with status code: %d\n", resp.StatusCode)

		resp.Body.Close()

		time.Sleep(requestInterval)
	}
}
