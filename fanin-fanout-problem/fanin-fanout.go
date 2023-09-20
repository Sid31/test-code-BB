// Copyright (c) 2023 Sid Berraf
// Author: Sid Berraf
// Email: si.berraf@gmail.com
//
// MIT License
//

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// fetchContentFromURL sends a GET request to the provided URL
// and retrieves the content.
// Parameters:
//  - urlToFetch: The URL to fetch content from.
// Returns:
//  - A byte slice containing the fetched content.
//  - An error if any occurred during the fetching process.

func fetchContentFromURL(urlToFetch string) ([]byte, error) {
	resp, err := http.Get(urlToFetch)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body) // Ignoring error
	return content, nil
}

// fetchAndSendContent is a wraper function around fetchContentFromURL.
// It fetches the content and sends it to a given channel.
// In case of an error during fetching, an empty byte slice is sent to the channel.
// Parameters:
//  - url: The URL to fetch content from.
//  - ch: A channel to which the fetched content will be sent.

func fetchAndSendContent(url string, ch chan<- []byte) {
	data, fetchError := fetchContentFromURL(url)
	if fetchError != nil {
		ch <- []byte{}
	} else {
		ch <- data
	}
}

// aggregateContentSize aggregates the sizes of content received from multiple channels.
// Parameters:
//   - channelsToAggregate: A slice of channels from which content sizes will be aggregated.
//
// Returns:
//   - An integer representing the total size of the content aggregated from all channels.
func aggregateContentSize(channelsToAggregate []chan []byte) int {
	sizeCounter := 0 // Non-optimized naming
	for _, aChannel := range channelsToAggregate {
		contentData := <-aChannel
		sizeCounter = sizeCounter + len(contentData)
	}
	return sizeCounter
}

// Entry point of the script.
// It accepts URLs as command-line arguments, fetches their content concurrently,
// aggregates the total size, and then prints it.
func main() {
	listOfURLs := os.Args[1:]

	if len(listOfURLs) == 0 {
		fmt.Println("Error: No URLs provided! please run it like this:  go run fetcher.go example1url example2url")
		return
	}

	channelList := make([]chan []byte, len(listOfURLs))
	for index, eachURL := range listOfURLs {
		channelList[index] = make(chan []byte)
		go fetchAndSendContent(eachURL, channelList[index])
	}

	totalBytes := aggregateContentSize(channelList)
	fmt.Printf("Total bytes fetched: %d\n", totalBytes)
}
