package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/imjoey/md-code-formatter/pkg/renderer"
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

	totalMdCount, changedMdCount := 0, 0

	err = filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".markdown" {

			content, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("Failed to read file: %s", err)
			}
			totalMdCount++
			r := renderer.NewMdCodeRenderer(content, renderer.WithHCLEnabled())

			file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				return fmt.Errorf("Failed to write file %s: %s", path, err)
			}

			defer file.Close()
			r.Render(file)
			if r.Stat.ChangedCount != 0 {
				changedMdCount++
				log.Printf("Formatted file <%s> : %d code blocks formatted", path, r.Stat.ChangedCount)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to walk dir %s: %s", dirpath, err)
	}

	log.Printf("Statistics: %d of %d markdown files formatted", changedMdCount, totalMdCount)
}
