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
	app.Name = "zlib"
	app.Usage = "A wrapper for the zlib compression algorithm."
	app.Action = func(c *cli.Context) {
		var err error
		var reader io.Reader = os.Stdin

		// If a filename was passed in as an argument, read from the file instead of stdin
		if len(c.Args()) > 0 {
			reader, err = os.Open(c.Args()[0])
			if err != nil {
				exit(err.Error(), 1)
			}
		}

		if c.Bool("decompress") {
			// If the input is base64 encoded, decode it first.
			if c.Bool("base64") {
				reader = base64.NewDecoder(base64.StdEncoding, reader)
			}
			z, err := zlib.NewReader(reader)
			if err != nil {
				exit(err.Error(), 1)
			}
			defer z.Close()
			if _, err := io.Copy(os.Stdout, z); err != nil {
				exit(err.Error(), 1)
			}
		} else {
			var writer io.Writer = os.Stdout
			// If base64 encoding is specified, encode the output to stdout
			if c.Bool("base64") {
				writer = base64.NewEncoder(base64.StdEncoding, writer)
			}
			z := zlib.NewWriter(writer)
			defer z.Close()
			if _, err := io.Copy(z, reader); err != nil {
				exit(err.Error(), 1)
			}
		}
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "d, decompress",
			Usage: "Decompresses the input instead of compressing the output.",
		},
		cli.BoolFlag{
			Name:  "b, base64",
			Usage: "Decodes the input or encodes the output using base64 encoding.",
		},
	}
	app.Run(os.Args)
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
