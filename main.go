package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const (
	usage = "./docker host:port"
)

var (
	domain     string
	currentdir string
)

func main() {

	args := os.Args

	if len(args) > 1 {
		domain = args[1]
	} else {
		fmt.Println("usage: ", usage)
		return
	}

	exe, err := os.Executable()
	if err == nil {
		currentdir = filepath.Dir(exe)
	} else {
		fmt.Println("unable to find exe path: ", err)
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
