package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Image   string       `json:"image"`
	Created int64        `json:"created"`
	Status  string       `json:"status"`
	Ports   []types.Port `json:"ports"`
}

func listContainersPage(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/docker/containers/list", domain))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var containers []Container
	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join(currentdir, "listContainers.html")))
	tmpl.Execute(w, containers)
}

func listContainers(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	cli.NegotiateAPIVersion(ctx)

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var containerList []Container
	for _, container := range containers {
		c := Container{
			ID:      container.ID[:10],
			Name:    strings.TrimPrefix(container.Names[0], "/"),
			Image:   container.Image,
			Created: container.Created,
			Status:  container.Status,
			Ports:   container.Ports,
		}
		containerList = append(containerList, c)
	}

	json.NewEncoder(w).Encode(containerList)
}
