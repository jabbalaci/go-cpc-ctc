package main

import "strings"

func print_help() {
	text := `
ctc (copy text to clipboard) v{ver}
https://github.com/jabbalaci/go-cpc-ctc

Usage examples:

* ctc                   print this help
* ctc -h, ctc --help    print this help
* ctc <text>            copy the given text to clipboard
* ctc -f <file>         copy the content of the given file to clipboard
`

	text = strings.TrimSpace(text)
	text = strings.Replace(text, "{ver}", VERSION, 1)

	println(text)
}
