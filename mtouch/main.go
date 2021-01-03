// Package main .
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// touch file, create parent directory when not exists.

var (
	version = "1.0.0"
	v       bool
)

func main() {
	flag.BoolVar(&v, "v", false, "show version")

	flag.Parse()

	if v {
		fmt.Println("Version: ", version)
		return
	}

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("usage: mtouch file ...")
		return
	}

	wd, _ := os.Getwd()

	for _, file := range files {
		file = strings.ReplaceAll(file, "\\", "/")

		if !strings.HasPrefix(file, "/") {
			file = filepath.Join(wd, file)
		}

		// file exists
		if _, err := os.Stat(file); err == nil {
			continue
		}

		fp := filepath.Dir(file)

		// fp not exists
		// prepare parent path
		if _, err := os.Stat(fp); err != nil {
			// create fp
			if err := os.MkdirAll(fp, 0755); err != nil {
				fmt.Println("create %s path %s error: %v", file, fp, err)
				continue
			}
		}

		// write blank file
		if err := ioutil.WriteFile(file, []byte(""), 0755); err != nil {
			fmt.Println("create %s error: %v", file, err)
			continue
		}

	}
}
