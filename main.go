package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	websocketproxy "github.com/relaytools/interceptor-proxy"
)

func main() {


	// TODO:  we need to support two different modes:

    // mode 1 'private relay') authenticate all connections accepting only pubkeys in the allowlist, and do not allow them through if they don't auth.
	// mode 2 'protected dms') authenticate all connections accepting any pubkey as a login, ALSO allow them through if they don't auth, EXCEPT FOR requests for DMs kind 4, and 1059/1060.

	port := os.Getenv("INTERCEPTOR_PORT")
    if port == "" {
        port = "9696" // default port if not specified
    }

	endpoint := os.Getenv("INTERCEPTOR_CONFIG_URL")
	if endpoint == "" {
		endpoint = "http://127.0.0.1:3000/api/sconfig/relays"
	}

	// modes are private_relay and protected_dms
	mode := os.Getenv("INTERCEPTOR_MODE")
	if mode == "" {
		mode = "private_relay"
	}

	log.Printf("Using config url: %s\n", endpoint)
	log.Printf("Listening on port :%s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), websocketproxy.NewProxy(endpoint, mode))
	if err != nil {
		log.Fatalln(err)
	}
}