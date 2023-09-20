package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchContentFromURL(t *testing.T) {
	server100 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(make([]byte, 100)))
	}))
	defer server100.Close()

	server1000 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(make([]byte, 1000)))
	}))
	defer server1000.Close()

	fmt.Println("Testing server with 100 bytes response...")
	if content, _ := fetchContentFromURL(server100.URL); len(content) != 100 {
		t.Errorf("Expected 100 bytes, got %d bytes", len(content))
	} else {
		fmt.Printf("Successfully fetched 100 bytes from %s\n", server100.URL)
	}

	fmt.Println("Testing server with 1000 bytes response...")
	if content, _ := fetchContentFromURL(server1000.URL); len(content) != 1000 {
		t.Errorf("Expected 1000 bytes, got %d bytes", len(content))
	} else {
		fmt.Printf("Successfully fetched 1000 bytes from %s\n", server1000.URL)
	}

}
