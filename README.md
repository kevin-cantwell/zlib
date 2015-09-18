# zlibr

It turned out to be quite difficult for me to find a command-line tool for decompressing and compressing zlib-compressed data, so I wrote my own.

Without any arguments, zlibr will compress, then base64-encode an input stream.

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
   -r, --raw        Decodes or encodes the output without assuming base64 encoding.
   --help, -h       show help
   --version, -v    print the version
```

## Examples

#### Decompressing a zlib-compressed file

```bash
zlibr --decompress --raw < compressedfile
```

#### Decoding and decompressing a zlib-compressed, base64-encoded, file


```bash
zlibr --decompress < compressedfile
```

#### Decoding and decompressing some base64-encoded and zlib-compressed input

```bash
echo -n 'eJyqVkrLz1eyUkpKLFKq5QIEAAD//yG4' | zlibr --decompress
```

Outputs: `{"foo":"bar"}`


#### Compressing and base64-encoding some json

```bash
echo -n '{"foo":"bar"}' | zlibr
```

Outputs: `eJyqVkrLz1eyUkpKLFKq5QIEAAD//yG4`


#### Compressing a file

```bash
zlibr --raw < file
```

#### Compressing and base64-encoding a file

```bash
zlibr < file
```