package main

import (
	"context"
	"fmt"
	"github.com/artarts36/oas-combiner/internal"
	cli "github.com/artarts36/singlecli"
	"os"
)

func main() {
	app := &cli.App{
		BuildInfo: &cli.BuildInfo{},
		Args: []*cli.ArgDefinition{
			{
				Name:        "spec",
				Description: "Basic specification",
				Required:    true,
			},
			{
				Name:        "output",
				Description: "Path to output file",
				Required:    true,
			},
		},
		UsageExamples: []*cli.UsageExample{
			{
				Command: "oas-combiner ./base.yaml ./openapi.yaml",
			},
		},
		Action: run,
	}

	app.RunWithGlobalArgs(context.Background())
}

func run(ctx *cli.Context) error {
	spec, err := internal.LoadSpec(ctx.Args["spec"])
	if err != nil {
		return fmt.Errorf("load spec: %w", err)
	}

	newSpec, err := internal.Combine(*spec)
	if err != nil {
		return fmt.Errorf("combine spec: %w", err)
	}

	outputContent, err := internal.MarshalSpec(&newSpec)
	if err != nil {
		return fmt.Errorf("marshal output: %w", err)
	}

	err = os.WriteFile(ctx.Args["output"], outputContent, 0755)
	if err != nil {
		return fmt.Errorf("write output: %w", err)
	}

	return nil
}
