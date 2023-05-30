package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getCountry(ctx context.Context, ip string) (error, string) {
	var (
		url      = fmt.Sprintf("http://ip-api.com/json/%s", ip)
		request  *http.Request
		response *http.Response
		err      error
		body     []byte
	)

	if request, err = http.NewRequestWithContext(ctx, "GET", url, nil); err != nil {
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

	result := struct {
		Country string `json:"country"`
	}{}

	if err = json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("unmarshal json: %w", err), ""
	}

	return nil, result.Country
}
