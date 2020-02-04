package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mholt/archiver"
)

func main() {
	flag.Parse()
	zipPath := flag.Arg(0)
	err := archiver.Walk(zipPath, func(f archiver.File) error {
		var entryPath string
		if h, ok := f.Header.(zip.FileHeader); ok {
			entryPath = h.Name
		} else {
			entryPath = f.Name()
		}
		// TODO: check insecure path
		if strings.HasPrefix(entryPath, "../") {
			return fmt.Errorf("Insecure path: %s", entryPath)
		}
		if f.IsDir() {
			entryPath += "/"
		}
		println(entryPath)
		io.Copy(ioutil.Discard, f)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
