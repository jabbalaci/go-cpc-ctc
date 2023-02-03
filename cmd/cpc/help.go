package main

import "strings"

func print_help() {
	text := `
cpc (copy path to clipboard) v{ver}
https://github.com/jabbalaci/go-cpc-ctc

- copies the path of the current working directory to the clipboard
- if a parameter is given, it's also added to the path

Usage: cpc [option] [parameter]
where option can be:
    -h or --help            get this help
`

	text = strings.TrimSpace(text)
	text = strings.Replace(text, "{ver}", VERSION, 1)

	println(text)
}
