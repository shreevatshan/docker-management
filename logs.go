package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func streamContainerLogs(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	cli.NegotiateAPIVersion(ctx)

	// Get containerID from query parameters
	containerID := r.URL.Query().Get("id")

	options := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Since:      fmt.Sprintf("%v", time.Now().Unix())}
	out, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	conn, err := upgrader.Upgrade(w, r, nil) // upgrade http connection to websocket
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not open websocket connection: %v", err), http.StatusBadRequest)
		return
	}

	buf := make([]byte, 1024)
	for {
		numBytes, err := out.Read(buf)
		if numBytes > 0 {
			if err := conn.WriteMessage(websocket.TextMessage, buf[8:numBytes]); err != nil {
				return
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
