include $(GOROOT)/src/Make.inc

ICU_CONFIG=$(shell which icu-config)
CGO_LDFLAGS=$(shell $(ICU_CONFIG) --libs)
CGO_CFLAGS=$(shell $(ICU_CONFIG) --cflags)

TARG=icu4go

CGOFILES=\
  charsetdetect.go\

CGO_OFILES=\
  helper.o\

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%.o: %.c
	gcc $(_CGO_CFLAGS_$(GOARCH)) -g -O2 -fPIC $(CFLAGS) -o $@ -c $^
