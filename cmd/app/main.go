package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// ...
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Getenv, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}