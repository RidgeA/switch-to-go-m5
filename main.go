package main

import (
	"flag"
	"fmt"
	"github.com/RidgeA/switch-to-go-m5/alphabet"
	"github.com/RidgeA/switch-to-go-m5/caesar"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"io"
	"os"
)

type (
	options struct {
		shift                         int
		useStdin, useStdout           bool
		operation                     string
		inputFilePath, outputFilePath string
	}
)

func parseOptions(args []string) options {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	o := options{}

	fs.StringVar(&o.operation, "operation", "", "one of: encode|decode (Required)")
	fs.IntVar(&o.shift, "shift", 0, "shift value to be used for Caesar cipher (Required)")
	fs.StringVar(&o.inputFilePath, "input", "", "path to input file, if not specified stdin will be used")
	fs.StringVar(&o.outputFilePath, "output", "", "path to output file, if not specified stdout will be used")

	_ = fs.Parse(args)

	if o.operation == "" {
		panic("operation should be defined")
	}

	if o.shift < 1 {
		panic("shift should be greater or equal 1")
	}

	if o.inputFilePath == "" {
		o.useStdin = true
	}

	if o.outputFilePath == "" {
		o.useStdout = true
	}

	return o
}

func createReader(opt options) (io.ReadCloser, error) {
	if opt.useStdin {
		return os.Stdin, nil
	}

	file, err := os.Open(opt.inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open input file: %w", err)
	}

	return file, nil
}

func createWriter(opt options) (io.WriteCloser, error) {
	if opt.useStdout {
		return os.Stdout, nil
	}

	file, err := os.Create(opt.outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot create output file: %w", err)
	}
	return file, nil
}

func main() {
	opt := parseOptions(os.Args[1:])

	reader, err := createReader(opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	defer reader.Close()

	writer, err := createWriter(opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	defer writer.Close()

	var transformer runes.Transformer
	if opt.operation == "encode" {
		transformer = caesar.Encode(opt.shift, alphabet.EnLower)
	} else {
		transformer = caesar.Decode(opt.shift, alphabet.EnLower)
	}

	_, err = io.Copy(writer, transform.NewReader(reader, transformer))
	if err != nil {
		fmt.Printf("Cannot write: %s", err.Error())
		os.Exit(1)
	}

}
