package main

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func manageContainerPage(w http.ResponseWriter, r *http.Request) {
	containerID := r.URL.Query().Get("id")

	tmpl := template.Must(template.ParseFiles("manageContainer.html"))
	tmpl.Execute(w, containerID)
}

func manageContainer(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	cli.NegotiateAPIVersion(ctx)

	containerID := r.URL.Query().Get("id")
	operation := r.URL.Query().Get("operation")

	switch operation {
	case "start":
		if err := cli.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "stop":
		if err := cli.ContainerStop(ctx, containerID, container.StopOptions{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "pause":
		if err := cli.ContainerPause(ctx, containerID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "unpause":
		if err := cli.ContainerUnpause(ctx, containerID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "restart":
		if err := cli.ContainerRestart(ctx, containerID, container.StopOptions{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, fmt.Sprintf("Invalid operation: %s", operation), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
