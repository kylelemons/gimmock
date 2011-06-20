package gimmock

import (
	"fmt"
	"io"
	"os"
)

func WriteMockFile(w io.Writer, pkg string, mocks []*Mock) os.Error {
	fmt.Fprintf(w, "package %s\n", pkg)

	fmt.Fprintf(w, "import (\n")
	for _, imp := range MockDependencies {
		fmt.Fprintf(w, "\t%q\n", imp)
	}
	fmt.Fprintf(w, ")\n")

	for _, mock := range mocks {
		fmt.Fprintf(w, "%s\n", mock.MockClass())
	}

	// TODO(kevlar): Error messages from printf
	return nil
}
