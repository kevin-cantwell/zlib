# zlibr

It turned out to be quite difficult for me to find a command-line tool for decompressing zlib-compressed data, so I wrote my own.

## Installation

`go get -u github.com/kevin-cantwell/zlibr/cmd/zlibr`

## Usage

```
NAME:
   zlibr - A wrapper for the zlib compression algorithm.

USAGE:
   zlibr [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -d, --decompress Decompresses the input instead of compressing the output.
   -b, --base64   Decodes the input or encodes the output using base64 encoding.
   --help, -h   show help
   --version, -v  print the version
```

## Examples

#### Decompressing a zlib-compressed file

```bash
zlibr -d < compressedfile
```

```bash
zlibr -d compressedfile
```

#### Decoding and decompressing a zlib-compressed and base64-encoded file


```bash
zlibr -d -b < compressedfile
```

```bash
zlibr -d -b compressedfile
```

#### Decoding and decompressing some base64-encoded and zlib-compressed input

```bash
echo 'eJyqVkrLz1eyUkpKLFKq5QIEAAD//yG4' | zlibr -d -b
```

Outputs: `{"foo":"bar"}`


#### Compressing and base64-encoding some json

```bash
echo '{"foo":"bar"}' | zlibr -b
```

Outputs: `eJyqVkrLz1eyUkpKLFKq5QIEAAD//yG4`


#### Compressing a file

```bash
zlibr < file
```

#### Compressing and base64-encoding a file

```bash
zlibr -b < file
```