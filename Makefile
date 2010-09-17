include $(GOROOT)/src/Make.inc

TARG=serial
GOFILES=\
	$(GOOS).go\
	serial.go\

include $(GOROOT)/src/Make.pkg
