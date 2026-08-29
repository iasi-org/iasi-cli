package cli

import (
	"fmt"
	"strings"
)

type Arguments struct {
	Unnamed *string
	Named   map[string]string
	Options map[string]bool
}

func ParseArguments(args []string) (Arguments, error) {
	arguments := Arguments{
		Named:   make(map[string]string),
		Options: make(map[string]bool),
	}

	unnamedSeen := false
	namedSeen := false

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			if arg == "--" || strings.HasPrefix(arg, "---") || strings.Contains(arg, "=") {
				return Arguments{}, fmt.Errorf("invalid command option %q", arg)
			}
			if unnamedSeen || namedSeen {
				return Arguments{}, fmt.Errorf("command option %q must appear before parameters", arg)
			}
			if arguments.Options[arg] {
				return Arguments{}, fmt.Errorf("command option %q specified more than once", arg)
			}
			arguments.Options[arg] = true
			continue
		}

		if strings.HasPrefix(arg, "-") {
			return Arguments{}, fmt.Errorf("unknown global option %q", arg)
		}

		if !strings.Contains(arg, "=") {
			if unnamedSeen {
				return Arguments{}, fmt.Errorf("only one unnamed argument is allowed")
			}
			if namedSeen {
				return Arguments{}, fmt.Errorf("unnamed argument must appear before named parameters")
			}

			value := arg
			arguments.Unnamed = &value
			unnamedSeen = true
			continue
		}

		parts := strings.SplitN(arg, "=", 2)
		name := parts[0]
		value := parts[1]

		if name == "" {
			return Arguments{}, fmt.Errorf("argument name cannot be empty")
		}
		if value == "" {
			return Arguments{}, fmt.Errorf("argument %q cannot have an empty value", name)
		}
		if _, exists := arguments.Named[name]; exists {
			return Arguments{}, fmt.Errorf("argument %q specified more than once", name)
		}

		arguments.Named[name] = value
		namedSeen = true
	}

	return arguments, nil
}
