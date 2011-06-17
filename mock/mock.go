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
		fmt.Printf("import \"os\"\n")
		fmt.Printf("type UnexpectedCall struct{Interface,Method string;Args []interface{}}\n")
		fmt.Printf("func (e *UnexpectedCall) String() string { return \"Unexpected Call\" }\n")
		fmt.Printf("type WrongArgument struct{Interface,Method string;Idx int;Got,Want interface{}}\n")
		fmt.Printf("func (e *WrongArgument) String() string { return \"Wrong Argument\" }\n")
		for _, mock := range mocks {
			fmt.Printf("%s\n", mock.MockClass())
		}
	}
}
