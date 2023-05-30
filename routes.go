package main

import (
	"context"
	"fmt"
	"os/exec"
)

func getRoutes(ctx context.Context) (error, []byte) {
	var (
		out []byte
		err error
	)

	if out, err = exec.CommandContext(ctx, "route", "get", "8.8.8.8").Output(); err != nil {
		return fmt.Errorf("exec command: %w", err), nil
	}

	return nil, out
}
