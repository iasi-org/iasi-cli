package cli

import "fmt"

type Parameter struct {
	Name     string
	Required bool
	Default  string
	Option   string
}

type Signature[T any] struct {
	Parameters []Parameter
	Build      func(values map[string]string) (T, error)
}

func ProcessArguments[T any](arguments Arguments, signature Signature[T]) (T, error) {
	var zero T

	parameters := make(map[string]Parameter, len(signature.Parameters))
	options := make(map[string]Parameter)

	for _, parameter := range signature.Parameters {
		if parameter.Name == "" {
			return zero, fmt.Errorf("signature contains a parameter with an empty name")
		}
		if _, exists := parameters[parameter.Name]; exists {
			return zero, fmt.Errorf("parameter %q is defined more than once in signature", parameter.Name)
		}
		if parameter.Option != "" {
			if _, exists := options[parameter.Option]; exists {
				return zero, fmt.Errorf("command option %q is defined more than once in signature", parameter.Option)
			}
			options[parameter.Option] = parameter
		}
		parameters[parameter.Name] = parameter
	}

	values := make(map[string]string, len(signature.Parameters))

	if arguments.Unnamed != nil {
		if len(signature.Parameters) == 0 {
			return zero, fmt.Errorf("unnamed argument is not allowed")
		}
		first := signature.Parameters[0]
		if first.Option != "" {
			return zero, fmt.Errorf("first parameter %q cannot be a command option", first.Name)
		}
		if _, exists := arguments.Named[first.Name]; exists {
			return zero, fmt.Errorf("argument %q specified more than once", first.Name)
		}
		values[first.Name] = *arguments.Unnamed
	}

	for name, value := range arguments.Named {
		parameter, exists := parameters[name]
		if !exists {
			return zero, fmt.Errorf("unknown argument %q", name)
		}
		if parameter.Option != "" {
			return zero, fmt.Errorf("argument %q must be specified as %s", name, parameter.Option)
		}
		values[name] = value
	}

	for option := range arguments.Options {
		parameter, exists := options[option]
		if !exists {
			return zero, fmt.Errorf("unknown command option %q", option)
		}
		if _, exists := values[parameter.Name]; exists {
			return zero, fmt.Errorf("argument %q specified more than once", parameter.Name)
		}
		values[parameter.Name] = "true"
	}

	for _, parameter := range signature.Parameters {
		if _, exists := values[parameter.Name]; exists {
			continue
		}
		if parameter.Default != "" {
			values[parameter.Name] = parameter.Default
			continue
		}
		if parameter.Option != "" {
			values[parameter.Name] = "false"
			continue
		}
		if parameter.Required {
			return zero, fmt.Errorf("argument %q is required", parameter.Name)
		}
	}

	if signature.Build == nil {
		return zero, fmt.Errorf("signature does not define a builder")
	}

	return signature.Build(values)
}
