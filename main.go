package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dirpath := "/Users/joey/Work/opensource/code/terraform-provider-alicloud/website/docs"

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".markdown" {
			log.Printf("Doer file: %s", path)
			Doer(path)
		}
		log.Printf("walk file: %s", path)
		return nil
	})
	if err != nil {
		log.Printf("Failed to walk dir %s: %s", dirpath, err)
	}
}
