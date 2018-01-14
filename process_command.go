package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type Process struct{}

func (*Process) Name() string             { return "process" }
func (*Process) SetFlags(f *flag.FlagSet) {}
func (*Process) Synopsis() string         { return "Process jobs" }
func (*Process) Usage() string            { return "process\n" }
func (*Process) Execute(
	_ context.Context, _ *flag.FlagSet, _ ...interface{},
) subcommands.ExitStatus {
	startProcess()
	return subcommands.ExitSuccess
}
