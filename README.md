Installation
============

To install, type `goinstall -u github.com/kylelemons/gimmock{,/mock}`

Usage
=====

Required Mocks
--------------

In order for gimmock to work, you have to tell it what mocks it needs to generate.
In your `_test.go` file, you indicate this by commenting a top-level element (usually a function)
with a comment of the following form:

    //MOCK(<interfaces>)

For example:

    import (
      "testing"
      "github.com/kylelemons/gimmock"
    )
    
    //MOCK(net.Conn)
    func TestSomething(t *testing.T) {
      // TODO (example here)
    }
