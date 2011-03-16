include $(GOROOT)/src/Make.inc

TARG=passwordcreator

GOFILES=\
	passwordcreator.go\
	main.go\
	

include $(GOROOT)/src/Make.cmd
