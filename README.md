# go-cpc-ctc

This repository contains two binary programs,
namely **cpc** and **ctc**.

## cpc

This is a Go implementation of the project https://github.com/jabbalaci/jabba-cpc .

The name "cpc" stands for _**c**opy **p**ath to **c**lipboard_.

`cpc` copies the path of the current working directory to the clipboard.
If a parameter is given, it's also added to the path.

Supported platforms: Windows and Linux (with X server).

### Installation

    $ go install github.com/jabbalaci/go-cpc-ctc/cmd/cpc@latest

### Help

    $ cpc -h
    cpc (copy path to clipboard) v0.1.0
    https://github.com/jabbalaci/go-cpc-ctc

    - copies the path of the current working directory to the clipboard
    - if a parameter is given, it's also added to the path

    Usage: cpc [option] [parameter]
    where option can be:
        -h or --help            get this help

### Windows

Example:

```
c:\> cd c:\download

c:\download> cpc
```

Now the current directory's path, "**c:\download**" is copied to the clipboard.

```
c:\download> cpc tree.jpg
```

Now the absolute path of the given file, "**c:\download\tree.jpg**" is copied to the clipboard.

### Linux

`cpc` relies on the external command `xsel` to manipulate the content of the clipboard.
Thus, you must install `xsel` using your package manager (under Ubuntu it's
`sudo apt install xsel`).

See the example above, it works similarly under Linux.

## ctc

This is a Go implementation of the project https://github.com/jabbalaci/jabba-ctc .

The name "ctc" stands for _**c**opy **t**ext to **c**lipboard_.

`ctc` can copy its argument (as text) to the clipboard OR it can
copy the content of a text file to the clipboard.

### Installation

    $ go install github.com/jabbalaci/go-cpc-ctc/cmd/ctc@latest

### Help

    $ ctc -h
    ctc (copy text to clipboard) v0.1.0
    https://github.com/jabbalaci/go-cpc-ctc

    Usage examples:

    * ctc                   print this help
    * ctc -h, ctc --help    print this help
    * ctc <text>            copy the given text to clipboard
    * ctc -f <file>         copy the content of the given file to clipboard

### Windows

Examples:

```
c:\> ctc hello.txt
```

Now the text "**hello.txt**" is copied to the clipboard.

```
c:\> ctc -f hello.txt
```

Now the **content** of the text file is copied to the clipboard.

### Linux

`ctc` relies on the external command `xsel` to manipulate the content of the clipboard.
Thus, you must install `xsel` using your package manager (under Ubuntu it's
`sudo apt install xsel`).

See the example above, it works similarly under Linux.

## Links

* https://github.com/jabbalaci/jabba-cpc , cpc in Rust
* https://github.com/jabbalaci/jabba-ctc , ctc in Rust
