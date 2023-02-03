package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

const DEBUG = false

// const DEBUG = true

func main() {
	text := ""
	must_set := false

	if len(os.Args) == 1 {
		print_help()
		os.Exit(0)
	}
	if len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments")
		os.Exit(1)
	}
	// else, if we have 1 or 2 arguments
	first := os.Args[1]
	if len(os.Args) == 2 {
		if first == "-h" || first == "--help" {
			print_help()
			os.Exit(0)
		}
		// else, we have a text
		text = first
		must_set = true
	}

	if len(os.Args) == 3 {
		second := os.Args[2]
		if first == "-f" {
			content, err := Read(second)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: cannot read file")
				os.Exit(1)
			}
			// else
			text = content
			must_set = true
		} else {
			fmt.Fprintln(os.Stderr, "Argument error")
			os.Exit(1)
		}
	}

	if must_set {
		clipboard.WriteAll(text)
	}

	if DEBUG {
		text, _ := clipboard.ReadAll()
		fmt.Println("#", text)
	}
}
