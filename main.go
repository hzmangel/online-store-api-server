package main

import (
	"log"
	"net/http"

	"github.com/hzmangel/online-store-api-server/config"
)

func main() {
	log.Fatal(http.ListenAndServe(":9090", config.Routes()))
}
