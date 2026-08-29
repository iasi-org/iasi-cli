package cli

import (
	"fmt"
	"strings"
)

type Arguments struct {
	Unnamed *string
	Named   map[string]string
}

func ParseArguments(args []string) (Arguments, error) {
	arguments := Arguments{
		Named: make(map[string]string),
	}

	for index, arg := range args {
		if !strings.Contains(arg, "=") {
			if index != 0 {
				return Arguments{}, fmt.Errorf("unnamed argument must be the first argument")
			}

			value := arg
			arguments.Unnamed = &value
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
	}

	return arguments, nil
}
