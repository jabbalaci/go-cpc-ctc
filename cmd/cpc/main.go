package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
)

const DEBUG = false

// const DEBUG = true

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if arg == "-h" || arg == "--help" {
		print_help()
		os.Exit(0)
	}
	// else
	path, _ := os.Getwd()
	if arg != "" {
		path = filepath.Join(path, arg)
	}
	clipboard.WriteAll(path)

	if DEBUG {
		text, _ := clipboard.ReadAll()
		fmt.Println("#", text)
	}
}
