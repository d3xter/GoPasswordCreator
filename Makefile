include $(GOROOT)/src/Make.$(GOARCH)

TARG=passwordcreator

all : main.go passwordcreator.go
	$(GC) passwordcreator.go
	$(GC) main.go
	$(LD) -o $(TARG) main.$O


clean:
	rm *.[568vq]
