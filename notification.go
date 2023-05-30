package main

import (
	"context"
	"fmt"
	"os/exec"
)

func sendNotification(ctx context.Context, title string, message string) error {
	var (
		err error
		arg = fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
	)

	if err = exec.CommandContext(ctx, "osascript", "-e", arg).Run(); err != nil {
		return fmt.Errorf("exec command: %w", err)
	}

	return nil
}
