package main

import (
	"flag"
	"log"
	"os"

	"github.com/kylelemons/gimmock"
)

var (
	pkg = flag.String("pkg", "", "Package for output file")
)

func main() {
	flag.Parse()

	var mocks []*gimmock.Mock
	var outpkg string

	for _, file := range flag.Args() {
		filepkg, filemocks, err := gimmock.Process(file)
		if err != nil {
			log.Fatalf("mock: %s", err)
		}

		if *pkg == "" {
			if outpkg == "" {
				outpkg = filepkg
			} else {
				if filepkg != outpkg {
					log.Fatalf("mock: file %s does not match package %s", file, filepkg)
				}
			}
		} else {
			outpkg = *pkg
		}

		mocks = append(mocks, filemocks...)
	}

	if len(mocks) == 0 {
		log.Fatalf("mock: no interfaces found")
	}

	gimmock.WriteMockFile(os.Stdout, outpkg, mocks)
}
