package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice of slice of string to hold slices of words
	buckets := make([][]string, 12)
	// Create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print len of each bucket (how many words it contains)
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	// Print the words in one of the buckets
	//fmt.Println(buckets[6])
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
