# zlibr

It turned out to be quite difficult for me to find a command-line tool for decompressing zlib-compressed data, so I wrote my own.

## Usage

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