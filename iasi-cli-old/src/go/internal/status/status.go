package status

import (
	"errors"
	"fmt"
	"strings"

	"iasi-cli/internal/resolver"
)

var ErrNotInstalled = errors.New("IASI is not installed for this location")

type Result struct {
	Layers []Layer
	Counts map[string]int
}

type Layer struct {
	Path, Version string
}

func Find(start string) (Result, error) {
	context, err := resolver.Resolve(start)
	if err != nil {
		if errors.Is(err, resolver.ErrNotInstalled) {
			return Result{}, ErrNotInstalled
		}
		return Result{}, err
	}
	layers := make([]Layer, len(context.Layers))
	for index, layer := range context.Layers {
		layers[index] = Layer{Path: layer.Path, Version: layer.Version}
	}
	return Result{Layers: layers, Counts: map[string]int{
		"instructions": len(context.Instructions),
		"commands":     len(context.Commands),
		"skills":       len(context.Skills),
		"mcp":          len(context.MCP),
		"adapters":     len(context.Adapters),
	}}, nil
}

func Format(result Result, binaryVersion string) string {
	var builder strings.Builder
	builder.WriteString("IASI\n\nBinary : ")
	builder.WriteString(binaryVersion)
	builder.WriteString("\n\nLayers (low → high precedence):\n")
	for index, layer := range result.Layers {
		fmt.Fprintf(&builder, "  %d. %s  %s\n", index+1, layer.Path, layer.Version)
	}
	fmt.Fprintf(&builder, "\nEffective:\n  Instructions : %d\n  Commands     : %d\n  Skills       : %d\n  MCP          : %d\n  Adapters     : %d\n", result.Counts["instructions"], result.Counts["commands"], result.Counts["skills"], result.Counts["mcp"], result.Counts["adapters"])
	return builder.String()
}
