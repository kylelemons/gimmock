include $(GOROOT)/src/Make.inc

.PHONY : mock all

extra : mock all

TARG=example
GOFILES=\
	iface.go\

mock :
	mock iface.go | gofmt > _gomock_test.go

include $(GOROOT)/src/Make.pkg
