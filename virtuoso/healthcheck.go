package main

import (
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://synbiohub:7777/rootCollections")
	if err != nil {
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		os.Exit(1)
	}
}
