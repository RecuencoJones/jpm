package main

import (
	"flag"
	"os"
)

func main() {
	installCommand := flag.NewFlagSet("install", flag.ExitOnError)
	publishCommand := flag.NewFlagSet("publish", flag.ExitOnError)

	switch os.Args[1] {
	case "install":
		installCommand.Parse(os.Args[2:])
	case "publish":
		publishCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
