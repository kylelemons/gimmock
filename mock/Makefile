include $(GOROOT)/src/Make.inc

deppkg=github.com/kylelemons/gimmock
pkgdir=$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

TARG=mock
GOFILES=\
	mock.go\

PREREQ+=$(foreach dep,$(deppkg),$(pkgdir)/$(dep).a)

include $(GOROOT)/src/Make.cmd
