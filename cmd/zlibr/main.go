package main

import (
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zlibr"
	app.Usage = "A wrapper for the zlib compression algorithm."
	app.Action = func(c *cli.Context) {
		// All input comes from stdin
		var reader io.Reader = os.Stdin

		if c.Bool("decompress") {
			if !c.Bool("raw") { // Then assume base64 encoding
				reader = pipeToBase64Decoder(reader)
			}
			compressorReadCloser := pipeToZlibDecompressor(reader)
			if _, err := io.Copy(os.Stdout, compressorReadCloser); err != nil {
				exit(err.Error(), 1)
			}
			compressorReadCloser.Close()
		} else {
			var writer io.Writer = os.Stdout
			if !c.Bool("raw") { // Then assume base64 encoding
				encoderWriteCloser := pipeToBase64Encoder(writer)
				defer encoderWriteCloser.Close()
				writer = encoderWriteCloser
			}
			compressorWriteCloser := pipeToZlibCompressor(writer)
			if _, err := io.Copy(compressorWriteCloser, reader); err != nil {
				exit(err.Error(), 1)
			}
			compressorWriteCloser.Close()
		}
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "d, decompress",
			Usage: "Decompresses the input instead of compressing the output.",
		},
		cli.BoolFlag{
			Name:  "r, raw",
			Usage: "Decodes or encodes the output without assuming base64 encoding.",
		},
	}
	app.Run(os.Args)
}

func pipeToBase64Decoder(reader io.Reader) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, reader)
}

func pipeToZlibDecompressor(reader io.Reader) io.ReadCloser {
	readerCloser, err := zlib.NewReader(reader)
	if err != nil {
		exit(err.Error(), 1)
	}
	// defer readerCloser.Close()
	return readerCloser
}

func pipeToBase64Encoder(writer io.Writer) io.WriteCloser {
	writerCloser := base64.NewEncoder(base64.StdEncoding, writer)
	// defer writerCloser.Close()
	return writerCloser
}

func pipeToZlibCompressor(writer io.Writer) io.WriteCloser {
	writerCloser := zlib.NewWriter(writer)
	// defer writerCloser.Close()
	return writerCloser
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
