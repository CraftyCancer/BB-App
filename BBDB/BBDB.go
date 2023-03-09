package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func handler4DB(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got message from BB Server\n")
	io.WriteString(w, "DB records are being added!\n")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Error in request from BBServer: %s", err)
		return
	}

	logLocation := "C:\\Users\\29131\\Documents\\GoWorkspace\\src\\BBDB\\dblog" + time.Now().Format("2006-01-02") + ".log"
	//logLocation := "/" + time.Now().Format("2006-01-02") + ".log"
	f, errf := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY, 0644)
	if errf != nil {
		fmt.Printf("Error in creating db file: %s", err)
		return
	}
	defer f.Close()

	dbw := bufio.NewWriter(f)

	// // _, errp := fmt.Fprintf(dbw, "%s \n", body)
	// // if errp != nil {
	// // 	fmt.Printf("Error writing msg to file: %s", errp)
	// // 	return
	// // }
	fmt.Fprintf(dbw, "%s \t", body)
	fmt.Fprintf(dbw, "%s \n", time.Now().Format("2006-01-02 2:3:5 PM"))
	dbw.Flush()
	// fmt.Printf("%s \t", body)
	// fmt.Printf("%s \n", time.Now().Format("2006-01-02 2:3:5 PM"))

}

func main() {

	http.HandleFunc("/BBDB", handler4DB)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
