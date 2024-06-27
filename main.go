package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	websocketproxy "github.com/relaytools/interceptor-proxy"
)

func main() {
	port := os.Getenv("INTERCEPTOR_PORT")
    if port == "" {
        port = "9696" // default port if not specified
    }

	// This is the endpoint interceptor will query for downstream proxy and pubkey ACLs
	endpoint := os.Getenv("INTERCEPTOR_CONFIG_URL")
	if endpoint == "" {
		endpoint = "http://127.0.0.1:3000/api/sconfig/relays"
	}

	log.Printf("Using config url: %s\n", endpoint)
	log.Printf("Listening on port :%s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), websocketproxy.NewProxy(endpoint))
	if err != nil {
		log.Fatalln(err)
	}
}