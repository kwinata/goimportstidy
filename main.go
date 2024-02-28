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
var ignore = flag.String("ignore", "", "ignore glob patterns")
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

	patterns := strings.Split(*ignore, ",")

	for _, file := range fileList {
		if !strings.HasSuffix(file, ".go") {
			continue
		}
		matchPattern := false
		for _, pattern := range patterns {
			isMatch, err := filepath.Match(pattern, file)
			if isMatch {
				matchPattern = true
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "err on ignore pattern: %v", err)
			}
		}
		if matchPattern {
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
