package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://textfiles.com/etext/AUTHORS/TWAIN/sawyr10.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)
}
