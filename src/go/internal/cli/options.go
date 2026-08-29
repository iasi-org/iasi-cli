package cli

type OutputOptions struct {
	Silent  bool
	Verbose bool
	Debug   bool
}

func ParseOutputOptions(args []string) (OutputOptions, []string) {
	options := OutputOptions{}
	remaining := make([]string, 0, len(args))

	for _, arg := range args {
		switch arg {
		case "-s":
			options.Silent = true
		case "-v":
			options.Verbose = true
		case "-d":
			options.Debug = true
		default:
			remaining = append(remaining, arg)
		}
	}

	return options, remaining
}
