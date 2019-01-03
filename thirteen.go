package main

import (
	"log"
	"os"

	"github.com/mdouchement/13h13/proxy"
	"github.com/vulcand/oxy/testutils"
)

func main() {
	service := os.Getenv("THIRTEEN_SERVICE")
	if service == "" {
		panic("no THIRTEEN_SERVICE environment variabled defined for the service name that is monitored")
	}

	raw := os.Getenv("THIRTEEN_FORWARD_TO")
	if raw == "" {
		panic("no THIRTEEN_FORWARD_TO environment variabled defined")
	}
	forward := testutils.ParseURI(raw)

	paddr := os.Getenv("THIRTEEN_PROXY_ADDRESS")
	if paddr == "" {
		paddr = "localhost:8080"
	}

	maddr := os.Getenv("THIRTEEN_METRICS_ADDRESS")
	if maddr == "" {
		maddr = "localhost:8081"
	}

	log.Printf("THIRTEEN_SERVICE=%s", service)
	log.Printf("THIRTEEN_FORWARD_TO=%s", raw)
	log.Printf("THIRTEEN_PROXY_ADDRESS=%s", paddr)
	log.Printf("THIRTEEN_METRICS_ADDRESS=%s", maddr)

	////
	///
	//

	err := proxy.ListenAndServe(paddr, maddr, service, forward)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
