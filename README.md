Installation
============

To install, type `goinstall -u github.com/kylelemons/gimmock{,/mock}`

Usage
=====

When testing source code with complicated object relationships, it can be
difficult to test them independently.  Mocking frameworks allow programmers to
inject an extra layer of code in between the code under test and the components
on which it depends.  Gimmock generates mock structures compatible with specific
interfaces and allows the programmer to examine calls and specify returns and
apply stub methods.

Source Code Markup
------------------

In order for gimmock to work, you have to tell it what mocks it needs to
generate.  In your `_test.go` file, you indicate this by commenting a top-level
element (usually a function) with a comment of the following form:

    //MOCK(<interfaces>)

For example:

    import (
      "os"
      "testing"
      "github.com/kylelemons/gimmock"
    )
    
    //MOCK(os.Error)
    func TestSomething(t *testing.T) {
      defer gimmock.Check(t)
      var err os.Error
      gimmock.Mock(&err).
        ExpectString().AndReturn("test").
        ExpectString().AndReturn("another")
    }

If the interface to be mocked is not in the current package, it can be specified
in the above `os.Error` format.  The package name, however, must match either the
last file segment in its import declaration or the assigned local package name
from its import statement.  For example:

    import somelib "somelib.googlecode.com/hg"
    
    //MOCK(somelib.SomeInterface)
    func TestSomething(t *testing.T) {
      defer gimmock.Check(t)

      gimmock.Mock(&err).
        ExpectString().AndReturn("test").
        ExpectString().AndReturn("another")
    }

The above is necessary even if `somelib` is the package name.  Additionally, the
package must be installed to the GOROOT package directory, not a GOPATH directory.
This restriction may be lifted in the future.
