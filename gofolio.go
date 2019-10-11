package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:] // Getting arguments passed after
	token := args[0]
	ticker := args[1]
	url := fmt.Sprintf("https://cloud.iexapis.com/stable/stock/%s/quote?token=%s", ticker, token)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", result)
}
