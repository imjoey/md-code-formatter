package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var dirpath string
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `md-code-formatter: A tool to format the HCL code blocks in markdown files
Usage: md-code-formatter [-d directory]

Options:
`)
		flag.PrintDefaults()
	}

	flag.StringVar(&dirpath, "d", "./", "The directory that contains the markdown files")
	flag.Parse()

	fileInfo, err := os.Stat(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("ERROR: %s not exists", dirpath)
		}
	}
	if !fileInfo.IsDir() {
		log.Fatalf("ERROR: the -d param only supports directory")
	}

	err = filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".markdown" {
			log.Printf("Formatted file: %s", path)
			Formatter(path)
		}
		log.Printf("walk file: %s", path)
		return nil
	})
	if err != nil {
		log.Printf("Failed to walk dir %s: %s", dirpath, err)
	}
}
