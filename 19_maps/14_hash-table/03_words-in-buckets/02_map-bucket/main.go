package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Create a map with a key of int
	// and with a value of another map
	// with a key of string, which will be the word
	// and a valie of int, which will be the number of times the word occurs
	buckets := make(map[int]map[string]int)
	// Create maps to hold words
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n][word]++
	}
	// Print words in a bucket
	for k, v := range buckets[6] {
		fmt.Println(v, " \t- ", k)
	}
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
