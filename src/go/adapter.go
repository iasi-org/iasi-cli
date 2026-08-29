package main

func runAdapter(args []string) {
	arguments := parseArguments(args)
	adapter   := processArgs(arguments)

	checkAdapterAvailable(adapter)
	prepareDestination(adapter)
	copyAdapter(adapter)
}