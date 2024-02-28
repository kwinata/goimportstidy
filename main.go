package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kwinata/goimportstidy/format"
)

var local = flag.String("local", "", "local package name, used for grouping")
var current = flag.String("current", "", "current repo name, used for grouping")
var write = flag.Bool("w", false, "write changes")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: goimportstidy [flags] [path]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func errAndExit(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, f, args...)
	os.Exit(2)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	fileList := make([]string, 0)
	err := filepath.Walk(flag.Arg(0), func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	if err != nil {
		errAndExit("fail to walk path: %v", err)
	}

	for _, file := range fileList {
		if !strings.HasSuffix(file, ".go") {
			continue
		}
		if strings.Contains(file, "mocks/") {
			continue
		}
		if strings.Contains(file, "mock.go") {
			continue
		}
		if strings.Contains(file, "mocks.go") {
			continue
		}
		if strings.Contains(file, "/gen/") {
			continue
		}
		
		s, err := os.Stat(file)
		if err != nil {
			errAndExit("failed to stat file: %v", err)
		}
		f, err := ioutil.ReadFile(file)
		if err != nil {
			errAndExit("failed to read file: %v", err)
		}

		output := format.File(string(f), *local, *current)

		if !*write {
			fmt.Print(string(output))
		}

		if err := ioutil.WriteFile(file, []byte(output), s.Mode()); err != nil {
			errAndExit("failed to format file: %v", err)
		}
	}
}
