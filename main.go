package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	usage = "./docker host:port"
)

var (
	domain string
)

func main() {

	args := os.Args

	if len(args) > 1 {
		domain = args[1]
	} else {
		fmt.Println("usage: ", usage)
		return
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/docker/containers/list", listContainers).Methods("GET")
	router.HandleFunc("/api/docker/container/manage", manageContainer).Methods("GET")
	router.HandleFunc("/api/docker/container/inspect", inspectContainer).Methods("GET")
	router.HandleFunc("/api/docker/container/logs", streamContainerLogs).Methods("GET")

	router.HandleFunc("/docker", listContainersPage).Methods("GET")
	router.HandleFunc("/docker/container", manageContainerPage).Methods("GET")

	router.HandleFunc("/loading.gif", loadingGIF).Methods("GET")

	log.Fatal(http.ListenAndServe(domain, router))
}
