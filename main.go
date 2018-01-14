package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&Enqueuer{}, "")
	subcommands.Register(&Process{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
