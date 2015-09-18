package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zlib"
	app.Usage = "A command-line tool for using the zlib compression algorithm."
	app.Action = func(c *cli.Context) {
		var reader io.Reader = os.Stdin
		if c.Bool("decompress") {
			compressorReadCloser, err := zlib.NewReader(reader)
			if err != nil {
				exit(err.Error(), 1)
			}
			if _, err := io.Copy(os.Stdout, compressorReadCloser); err != nil {
				exit(err.Error(), 1)
			}
			compressorReadCloser.Close()
		} else {
			var writer io.Writer = os.Stdout
			compressorWriteCloser := zlib.NewWriter(writer)
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
	}
	app.Run(os.Args)
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
