package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func handler4alice(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got message from Alice\n")
	io.WriteString(w, "Message received from Alice!\n")

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var msg map[string]string
	err1 := json.Unmarshal([]byte(body), &msg)

	if err1 != nil {
		fmt.Printf("Error in Alice's request: %s", err1)
		return
	}

	fmt.Printf("Message received from Alice: %s \n", msg["post"])

	//c := http.Client{Timeout: time.Duration(1) * time.Second}
	c := http.Client{}
	postData := bytes.NewBuffer([]byte(msg["post"]))
	resp, err := c.Post("http://localhost:3333/BBDB", "application/json", postData)
	if err != nil {
		fmt.Printf("Error in request from Alice to BBDB: %s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error in response from BBDB: %s", err)
		return
	}

	fmt.Printf("Response from BBDB for Alice: %s \n", rbody)

}

func handler4bob(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got message from Bob\n")
	io.WriteString(w, "Message received from Bob!\n")

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var msg map[string]string
	err2 := json.Unmarshal([]byte(body), &msg)

	if err2 != nil {
		fmt.Printf("Error in Bob's request: %s", err2)
		return
	}

	fmt.Printf("Message received from Bob: %s \n", msg["post"])

	//c := http.Client{Timeout: time.Duration(1) * time.Second}
	c := http.Client{}
	postData := bytes.NewBuffer([]byte(msg["post"]))
	resp, err := c.Post("http://localhost:3333/BBDB", "application/json", postData)
	if err != nil {
		fmt.Printf("Error in request from Bob to BBDB: %s", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error in response from BBDB: %s", err)
		return
	}

	fmt.Printf("Response from BBDB for Bob: %s \n", rbody)
}

func main() {
	http.HandleFunc("/Alice", handler4alice)
	http.HandleFunc("/Bob", handler4bob)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
