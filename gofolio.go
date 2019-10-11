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
	tickers := args[1]
	// url := fmt.Sprintf("https://cloud.iexapis.com/stable/stock/%s/quote?token=%s", ticker, token)
	url := fmt.Sprintf("https://cloud.iexapis.com/stable/tops?token=%s&symbols=%s", token, tickers)
	res, err := http.Get(url)
	fmt.Printf("%s", args[1])
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
