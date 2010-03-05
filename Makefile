include $(GOROOT)/src/Make.$(GOARCH)

TARG=password_creator

all : main.go password_creator.go
	$(GC) password_creator.go
	$(GC) main.go
	$(LD) -o $(TARG) main.$O


clean:
	rm *.6
