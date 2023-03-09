package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//c := http.Client{Timeout: time.Duration(1) * time.Second}
	c := http.Client{}

	for {
		postData := bytes.NewBuffer([]byte(`{"post":"Alice says Hi"}`))
		resp, err := c.Post("http://localhost:8080/Alice", "application/json", postData)
		if err != nil {
			fmt.Printf("Error in Alice client %s", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("Error in response read %s", err)
			return
		}

		fmt.Printf("Response from BBServer: %s", body)

		time.Sleep(3 * time.Second)
	}
}
