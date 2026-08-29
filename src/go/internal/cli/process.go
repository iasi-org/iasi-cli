package cli

import "fmt"

type Parameter struct {
	Name     string
	Required bool
	Default  string
}

type Signature[T any] struct {
	Parameters []Parameter
	Build      func(values map[string]string) (T, error)
}

func ProcessArguments[T any](arguments Arguments, signature Signature[T]) (T, error) {
	var zero T

	parameters := make(map[string]Parameter, len(signature.Parameters))
	for _, parameter := range signature.Parameters {
		if parameter.Name == "" {
			return zero, fmt.Errorf("signature contains a parameter with an empty name")
		}
		if _, exists := parameters[parameter.Name]; exists {
			return zero, fmt.Errorf("parameter %q is defined more than once in signature", parameter.Name)
		}
		parameters[parameter.Name] = parameter
	}

	values := make(map[string]string, len(signature.Parameters))
	if arguments.Unnamed != nil {
		if len(signature.Parameters) == 0 {
			return zero, fmt.Errorf("unnamed argument is not allowed")
		}
		first := signature.Parameters[0]
		if _, exists := arguments.Named[first.Name]; exists {
			return zero, fmt.Errorf("argument %q specified more than once", first.Name)
		}
		values[first.Name] = *arguments.Unnamed
	}

	for name, value := range arguments.Named {
		if _, exists := parameters[name]; !exists {
			return zero, fmt.Errorf("unknown argument %q", name)
		}
		values[name] = value
	}

	for _, parameter := range signature.Parameters {
		if _, exists := values[parameter.Name]; exists {
			continue
		}
		if parameter.Default != "" {
			values[parameter.Name] = parameter.Default
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
