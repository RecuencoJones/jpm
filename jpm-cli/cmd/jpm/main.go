package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	installCommand := flag.NewFlagSet("install", flag.ExitOnError)
	publishCommand := flag.NewFlagSet("publish", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("install or publish subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "install":
		installCommand.Parse(os.Args[2:])
	case "publish":
		publishCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if installCommand.Parsed() {
		Install()
	}

	if publishCommand.Parsed() {
		Publish()
	}
}
