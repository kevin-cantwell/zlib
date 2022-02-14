# zlib

I couldn't find a command-line tool for decompressing and compressing zlib data, so I wrote my own. 
That's about the long and short of it.

_Update_: Someone noted that [`pigz`](https://formulae.brew.sh/formula/pigz) exists and can do zlib compression with `pigz -z`. Nice!

## Installation

Download a [release](https://github.com/kevin-cantwell/zlib/releases)  or install the latest directly using Go:

`go install github.com/kevin-cantwell/zlib/cmd/zlib@latest`

For older versions of Go, try:
`go get -u github.com/kevin-cantwell/zlib/cmd/zlib`

## Usage

Without any arguments, zlib will compress an input stream. Use the `-d` flag for decompression.

```
NAME:
   zlib - A wrapper for the zlib compression algorithm.

USAGE:
   zlib [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -d, --decompress Decompresses the input instead of compressing the output.
   --help, -h       show help
   --version, -v    print the version
```

## Examples

#### Compressing a file

```bash
zlib < file > file.zlib
```

#### Decompressing a zlib-compressed file

```bash
zlib -d < file.zlib
```

#### Compressing and base64-encoding some json

```bash
echo -n '{"foo":"bar"}' | zlib | base64
```

Outputs: `eJyqVkrLz1eyUkpKLFKqBQQAAP//HXoENA==`

#### Decoding and decompressing some base64-encoded and zlib-compressed input

```bash
echo -n 'eJyqVkrLz1eyUkpKLFKqBQQAAP//HXoENA==' | base64 -D | zlib -d
```

Outputs: `{"foo":"bar"}`



