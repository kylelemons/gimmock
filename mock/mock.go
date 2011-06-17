package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kylelemons/gimmock"
)

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		pkg, mocks, err := gimmock.Process(file)
		if err != nil {
			log.Fatalf("mock: %s", err)
		}
		fmt.Printf("package %s\n", pkg)

		fmt.Printf("import (\n")
		for _, imp := range gimmock.MockDependencies {
			fmt.Printf("\t%q\n", imp)
		}
		fmt.Printf(")\n")

		for _, mock := range mocks {
			fmt.Printf("%s\n", mock.MockClass())
		}
	}
}
