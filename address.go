package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func getIpAddress(ctx context.Context) (error, string) {
	var (
		request  *http.Request
		response *http.Response
		err      error
		body     []byte
	)

	if request, err = http.NewRequestWithContext(ctx, "GET", "http://ifconfig.me", nil); err != nil {
		return fmt.Errorf("create request: %w", err), ""
	}

	client := &http.Client{}

	if response, err = client.Do(request); err != nil {
		return fmt.Errorf("do request: %w", err), ""
	}

	defer func() {
		_ = response.Body.Close()
		client.CloseIdleConnections()
	}()

	if body, err = io.ReadAll(response.Body); err != nil {
		return fmt.Errorf("read body: %w", err), ""
	}

	return nil, string(body)
}
