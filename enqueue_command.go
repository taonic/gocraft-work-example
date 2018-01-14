package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type Enqueuer struct{}

func (*Enqueuer) Name() string             { return "enqueue" }
func (*Enqueuer) SetFlags(f *flag.FlagSet) {}
func (*Enqueuer) Synopsis() string         { return "Enqueue jobs" }
func (*Enqueuer) Usage() string            { return "enqueue <number of jobs>\n" }
func (*Enqueuer) Execute(
	_ context.Context, _ *flag.FlagSet, _ ...interface{},
) subcommands.ExitStatus {
	enqueueEmail()
	enqueueS3()
	return subcommands.ExitSuccess
}
