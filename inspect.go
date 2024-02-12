package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/client"
)

func inspectContainer(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	cli.NegotiateAPIVersion(ctx)

	containerID := r.URL.Query().Get("id")

	containerJSON, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(containerJSON)
}
